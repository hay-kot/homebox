package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/hay-kot/homebox/backend/app/api/handlers/debughandlers"
	v1 "github.com/hay-kot/homebox/backend/app/api/handlers/v1"
	_ "github.com/hay-kot/homebox/backend/app/api/static/docs"
	"github.com/hay-kot/homebox/backend/internal/repo"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

const prefix = "/api"

var (
	ErrDir = errors.New("path is dir")

	//go:embed all:static/public/*
	public embed.FS
)

func (a *app) debugRouter() *http.ServeMux {
	dbg := http.NewServeMux()
	debughandlers.New(dbg)

	return dbg
}

// registerRoutes registers all the routes for the API
func (a *app) newRouter(repos *repo.AllRepos) *chi.Mux {
	registerMimes()

	r := chi.NewRouter()
	a.setGlobalMiddleware(r)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s://%s/swagger/doc.json", a.conf.Swagger.Scheme, a.conf.Swagger.Host)),
	))

	// =========================================================================
	// API Version 1

	v1Base := v1.BaseUrlFunc(prefix)
	v1Ctrl := v1.NewControllerV1(a.services,
		v1.WithMaxUploadSize(a.conf.Web.MaxUploadSize),
		v1.WithRegistration(a.conf.AllowRegistration),
		v1.WithDemoStatus(a.conf.Demo), // Disable Password Change in Demo Mode
	)
	r.Get(v1Base("/status"), v1Ctrl.HandleBase(func() bool { return true }, v1.Build{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	}))

	r.Post(v1Base("/users/register"), v1Ctrl.HandleUserRegistration())
	r.Post(v1Base("/users/login"), v1Ctrl.HandleAuthLogin())

	// Attachment download URl needs a `token` query param to be passed in the request.
	// and also needs to be outside of the `auth` middleware.
	r.Get(v1Base("/items/{id}/attachments/download"), v1Ctrl.HandleItemAttachmentDownload())

	r.Group(func(r chi.Router) {
		r.Use(a.mwAuthToken)
		r.Get(v1Base("/users/self"), v1Ctrl.HandleUserSelf())
		r.Put(v1Base("/users/self"), v1Ctrl.HandleUserSelfUpdate())
		r.Delete(v1Base("/users/self"), v1Ctrl.HandleUserSelfDelete())
		r.Post(v1Base("/users/logout"), v1Ctrl.HandleAuthLogout())
		r.Get(v1Base("/users/refresh"), v1Ctrl.HandleAuthRefresh())
		r.Put(v1Base("/users/self/change-password"), v1Ctrl.HandleUserSelfChangePassword())

		r.Post(v1Base("/groups/invitations"), v1Ctrl.HandleGroupInvitationsCreate())

		// TODO: I don't like /groups being the URL for users
		r.Get(v1Base("/groups"), v1Ctrl.HandleGroupGet())
		r.Put(v1Base("/groups"), v1Ctrl.HandleGroupUpdate())

		r.Get(v1Base("/locations"), v1Ctrl.HandleLocationGetAll())
		r.Post(v1Base("/locations"), v1Ctrl.HandleLocationCreate())
		r.Get(v1Base("/locations/{id}"), v1Ctrl.HandleLocationGet())
		r.Put(v1Base("/locations/{id}"), v1Ctrl.HandleLocationUpdate())
		r.Delete(v1Base("/locations/{id}"), v1Ctrl.HandleLocationDelete())

		r.Get(v1Base("/labels"), v1Ctrl.HandleLabelsGetAll())
		r.Post(v1Base("/labels"), v1Ctrl.HandleLabelsCreate())
		r.Get(v1Base("/labels/{id}"), v1Ctrl.HandleLabelGet())
		r.Put(v1Base("/labels/{id}"), v1Ctrl.HandleLabelUpdate())
		r.Delete(v1Base("/labels/{id}"), v1Ctrl.HandleLabelDelete())

		r.Get(v1Base("/items"), v1Ctrl.HandleItemsGetAll())
		r.Post(v1Base("/items/import"), v1Ctrl.HandleItemsImport())
		r.Post(v1Base("/items"), v1Ctrl.HandleItemsCreate())
		r.Get(v1Base("/items/{id}"), v1Ctrl.HandleItemGet())
		r.Put(v1Base("/items/{id}"), v1Ctrl.HandleItemUpdate())
		r.Delete(v1Base("/items/{id}"), v1Ctrl.HandleItemDelete())

		r.Post(v1Base("/items/{id}/attachments"), v1Ctrl.HandleItemAttachmentCreate())
		r.Get(v1Base("/items/{id}/attachments/{attachment_id}"), v1Ctrl.HandleItemAttachmentToken())
		r.Put(v1Base("/items/{id}/attachments/{attachment_id}"), v1Ctrl.HandleItemAttachmentUpdate())
		r.Delete(v1Base("/items/{id}/attachments/{attachment_id}"), v1Ctrl.HandleItemAttachmentDelete())
	})

	r.NotFound(notFoundHandler())
	return r
}

// logRoutes logs the routes of the server that are registered within Server.registerRoutes(). This is useful for debugging.
// See https://github.com/go-chi/chi/issues/332 for details and inspiration.
func (a *app) logRoutes(r *chi.Mux) {
	desiredSpaces := 10

	walkFunc := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		text := "[" + method + "]"

		for len(text) < desiredSpaces {
			text = text + " "
		}

		fmt.Printf("Registered Route: %s%s\n", text, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

func registerMimes() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		panic(err)
	}

	err = mime.AddExtensionType(".mjs", "application/javascript")
	if err != nil {
		panic(err)
	}
}

// notFoundHandler perform the main logic around handling the internal SPA embed and ensuring that
// the client side routing is handled correctly.
func notFoundHandler() http.HandlerFunc {
	tryRead := func(fs embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
		f, err := fs.Open(path.Join(prefix, requestedPath))
		if err != nil {
			return err
		}
		defer f.Close()

		stat, _ := f.Stat()
		if stat.IsDir() {
			return ErrDir
		}

		contentType := mime.TypeByExtension(filepath.Ext(requestedPath))
		w.Header().Set("Content-Type", contentType)
		_, err = io.Copy(w, f)
		return err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := tryRead(public, "static/public", r.URL.Path, w)
		if err == nil {
			return
		}
		err = tryRead(public, "static/public", "index.html", w)
		if err != nil {
			panic(err)
		}
	}
}

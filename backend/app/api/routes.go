package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	_ "github.com/hay-kot/content/backend/app/api/docs"
	v1 "github.com/hay-kot/content/backend/app/api/v1"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

//go:embed all:public/*
var public embed.FS

const prefix = "/api"

// registerRoutes registers all the routes for the API
func (a *app) newRouter(repos *repo.AllRepos) *chi.Mux {
	r := chi.NewRouter()
	a.setGlobalMiddleware(r)

	// =========================================================================
	// Base Routes

	DumpEmbedContents()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s://%s/swagger/doc.json", a.conf.Swagger.Scheme, a.conf.Swagger.Host)),
	))

	// =========================================================================
	// API Version 1

	v1Base := v1.BaseUrlFunc(prefix)
	v1Ctrl := v1.NewControllerV1(a.services)
	{
		r.Get(v1Base("/status"), v1Ctrl.HandleBase(func() bool { return true }, "v1"))

		r.Post(v1Base("/users/register"), v1Ctrl.HandleUserRegistration())
		r.Post(v1Base("/users/login"), v1Ctrl.HandleAuthLogin())

		r.Group(func(r chi.Router) {
			r.Use(a.mwAuthToken)
			r.Get(v1Base("/users/self"), v1Ctrl.HandleUserSelf())
			r.Put(v1Base("/users/self"), v1Ctrl.HandleUserSelfUpdate())
			r.Delete(v1Base("/users/self"), v1Ctrl.HandleUserSelfDelete())
			r.Put(v1Base("/users/self/password"), v1Ctrl.HandleUserUpdatePassword())
			r.Post(v1Base("/users/logout"), v1Ctrl.HandleAuthLogout())
			r.Get(v1Base("/users/refresh"), v1Ctrl.HandleAuthRefresh())

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
			r.Post(v1Base("/items"), v1Ctrl.HandleItemsCreate())
			r.Get(v1Base("/items/{id}"), v1Ctrl.HandleItemGet())
			r.Put(v1Base("/items/{id}"), v1Ctrl.HandleItemUpdate())
			r.Delete(v1Base("/items/{id}"), v1Ctrl.HandleItemDelete())
		})
	}

	r.NotFound(NotFoundHandler)

	return r
}

// LogRoutes logs the routes of the server that are registered within Server.registerRoutes(). This is useful for debugging.
// See https://github.com/go-chi/chi/issues/332 for details and inspiration.
func (a *app) LogRoutes(r *chi.Mux) {
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

var ErrDir = errors.New("path is dir")

func init() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		panic(err)
	}

	err = mime.AddExtensionType(".mjs", "application/javascript")
	if err != nil {
		panic(err)
	}
}

func tryRead(fs embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := tryRead(public, "public", r.URL.Path, w)
	if err == nil {
		return
	}
	log.Debug().
		Str("path", r.URL.Path).
		Msg("served from embed not found - serving index.html")
	err = tryRead(public, "public", "index.html", w)
	if err != nil {
		panic(err)
	}
}

func DumpEmbedContents() {
	// recursively prints all contents in the embed.FS
	err := fs.WalkDir(public, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

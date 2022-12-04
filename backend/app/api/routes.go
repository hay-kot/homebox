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

	"github.com/hay-kot/homebox/backend/app/api/handlers/debughandlers"
	v1 "github.com/hay-kot/homebox/backend/app/api/handlers/v1"
	_ "github.com/hay-kot/homebox/backend/app/api/static/docs"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/pkgs/server"
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
func (a *app) mountRoutes(repos *repo.AllRepos) {
	registerMimes()

	a.server.Get("/swagger/*", server.ToHandler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s://%s/swagger/doc.json", a.conf.Swagger.Scheme, a.conf.Swagger.Host)),
	)))

	// =========================================================================
	// API Version 1

	v1Base := v1.BaseUrlFunc(prefix)

	v1Ctrl := v1.NewControllerV1(
		a.services,
		a.repos,
		v1.WithMaxUploadSize(a.conf.Web.MaxUploadSize),
		v1.WithRegistration(a.conf.Options.AllowRegistration),
		v1.WithDemoStatus(a.conf.Demo), // Disable Password Change in Demo Mode
	)

	a.server.Get(v1Base("/status"), v1Ctrl.HandleBase(func() bool { return true }, v1.Build{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	}))

	a.server.Post(v1Base("/users/register"), v1Ctrl.HandleUserRegistration())
	a.server.Post(v1Base("/users/login"), v1Ctrl.HandleAuthLogin())

	userMW := []server.Middleware{
		a.mwAuthToken,
		a.mwRoles(RoleModeOr, authroles.RoleUser.String()),
	}

	a.server.Get(v1Base("/users/self"), v1Ctrl.HandleUserSelf(), userMW...)
	a.server.Put(v1Base("/users/self"), v1Ctrl.HandleUserSelfUpdate(), userMW...)
	a.server.Delete(v1Base("/users/self"), v1Ctrl.HandleUserSelfDelete(), userMW...)
	a.server.Post(v1Base("/users/logout"), v1Ctrl.HandleAuthLogout(), userMW...)
	a.server.Get(v1Base("/users/refresh"), v1Ctrl.HandleAuthRefresh(), userMW...)
	a.server.Put(v1Base("/users/self/change-password"), v1Ctrl.HandleUserSelfChangePassword(), userMW...)

	a.server.Post(v1Base("/groups/invitations"), v1Ctrl.HandleGroupInvitationsCreate(), userMW...)
	a.server.Get(v1Base("/groups/statistics"), v1Ctrl.HandleGroupStatistics(), userMW...)
	a.server.Get(v1Base("/groups/statistics/purchase-price"), v1Ctrl.HandleGroupStatisticsPriceOverTime(), userMW...)

	// TODO: I don't like /groups being the URL for users
	a.server.Get(v1Base("/groups"), v1Ctrl.HandleGroupGet(), userMW...)
	a.server.Put(v1Base("/groups"), v1Ctrl.HandleGroupUpdate(), userMW...)

	a.server.Post(v1Base("/actions/ensure-asset-ids"), v1Ctrl.HandleEnsureAssetID(), userMW...)

	a.server.Get(v1Base("/locations"), v1Ctrl.HandleLocationGetAll(), userMW...)
	a.server.Post(v1Base("/locations"), v1Ctrl.HandleLocationCreate(), userMW...)
	a.server.Get(v1Base("/locations/{id}"), v1Ctrl.HandleLocationGet(), userMW...)
	a.server.Put(v1Base("/locations/{id}"), v1Ctrl.HandleLocationUpdate(), userMW...)
	a.server.Delete(v1Base("/locations/{id}"), v1Ctrl.HandleLocationDelete(), userMW...)

	a.server.Get(v1Base("/labels"), v1Ctrl.HandleLabelsGetAll(), userMW...)
	a.server.Post(v1Base("/labels"), v1Ctrl.HandleLabelsCreate(), userMW...)
	a.server.Get(v1Base("/labels/{id}"), v1Ctrl.HandleLabelGet(), userMW...)
	a.server.Put(v1Base("/labels/{id}"), v1Ctrl.HandleLabelUpdate(), userMW...)
	a.server.Delete(v1Base("/labels/{id}"), v1Ctrl.HandleLabelDelete(), userMW...)

	a.server.Get(v1Base("/items"), v1Ctrl.HandleItemsGetAll(), userMW...)
	a.server.Post(v1Base("/items/import"), v1Ctrl.HandleItemsImport(), userMW...)
	a.server.Post(v1Base("/items"), v1Ctrl.HandleItemsCreate(), userMW...)
	a.server.Get(v1Base("/items/{id}"), v1Ctrl.HandleItemGet(), userMW...)
	a.server.Put(v1Base("/items/{id}"), v1Ctrl.HandleItemUpdate(), userMW...)
	a.server.Delete(v1Base("/items/{id}"), v1Ctrl.HandleItemDelete(), userMW...)

	a.server.Post(v1Base("/items/{id}/attachments"), v1Ctrl.HandleItemAttachmentCreate(), userMW...)
	a.server.Put(v1Base("/items/{id}/attachments/{attachment_id}"), v1Ctrl.HandleItemAttachmentUpdate(), userMW...)
	a.server.Delete(v1Base("/items/{id}/attachments/{attachment_id}"), v1Ctrl.HandleItemAttachmentDelete(), userMW...)

	a.server.Get(
		v1Base("/items/{id}/attachments/{attachment_id}"),
		v1Ctrl.HandleItemAttachmentGet(),
		a.mwAuthToken, a.mwRoles(RoleModeOr, authroles.RoleUser.String(), authroles.RoleAttachments.String()),
	)

	a.server.NotFound(notFoundHandler())
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
func notFoundHandler() server.HandlerFunc {
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

	return func(w http.ResponseWriter, r *http.Request) error {
		err := tryRead(public, "static/public", r.URL.Path, w)
		if err != nil {
			// Fallback to the index.html file.
			// should succeed in all cases.
			err = tryRead(public, "static/public", "index.html", w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

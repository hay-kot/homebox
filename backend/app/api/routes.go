package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hay-kot/content/backend/app/api/base"
	_ "github.com/hay-kot/content/backend/app/api/docs"
	v1 "github.com/hay-kot/content/backend/app/api/v1"
	"github.com/hay-kot/content/backend/internal/repo"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

const prefix = "/api"

// registerRoutes registers all the routes for the API
func (a *app) newRouter(repos *repo.AllRepos) *chi.Mux {
	r := chi.NewRouter()
	a.setGlobalMiddleware(r)

	// =========================================================================
	// Base Routes

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s://%s/swagger/doc.json", a.conf.Swagger.Scheme, a.conf.Swagger.Host)),
	))

	// Server Favicon
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	baseHandler := base.NewBaseController(a.server)
	r.Get(prefix+"/status", baseHandler.HandleBase(func() bool { return true }, "v1"))

	// =========================================================================
	// API Version 1

	v1Base := v1.BaseUrlFunc(prefix)
	{
		v1Handlers := v1.NewControllerV1(a.services)
		r.Post(v1Base("/users/register"), v1Handlers.HandleUserRegistration())
		r.Post(v1Base("/users/login"), v1Handlers.HandleAuthLogin())
		r.Group(func(r chi.Router) {
			r.Use(a.mwAuthToken)
			r.Get(v1Base("/users/self"), v1Handlers.HandleUserSelf())
			r.Put(v1Base("/users/self"), v1Handlers.HandleUserUpdate())
			r.Put(v1Base("/users/self/password"), v1Handlers.HandleUserUpdatePassword())
			r.Post(v1Base("/users/logout"), v1Handlers.HandleAuthLogout())
			r.Get(v1Base("/users/refresh"), v1Handlers.HandleAuthRefresh())

			r.Get(v1Base("/locations"), v1Handlers.HandleLocationGetAll())
			r.Post(v1Base("/locations"), v1Handlers.HandleLocationCreate())
			r.Get(v1Base("/locations/{id}"), v1Handlers.HandleLocationGet())
			r.Put(v1Base("/locations/{id}"), v1Handlers.HandleLocationUpdate())
			r.Delete(v1Base("/locations/{id}"), v1Handlers.HandleLocationDelete())

			r.Get(v1Base("/labels"), v1Handlers.HandleLabelsGetAll())
			r.Post(v1Base("/labels"), v1Handlers.HandleLabelsCreate())
			r.Get(v1Base("/labels/{id}"), v1Handlers.HandleLabelGet())
			r.Put(v1Base("/labels/{id}"), v1Handlers.HandleLabelUpdate())
			r.Delete(v1Base("/labels/{id}"), v1Handlers.HandleLabelDelete())

			r.Get(v1Base("/items"), v1Handlers.HandleItemsGetAll())
			r.Post(v1Base("/items"), v1Handlers.HandleItemsCreate())
			r.Get(v1Base("/items/{id}"), v1Handlers.HandleItemGet())
			r.Put(v1Base("/items/{id}"), v1Handlers.HandleItemUpdate())
			r.Delete(v1Base("/items/{id}"), v1Handlers.HandleItemDelete())
		})
	}

	return r
}

// LogRoutes logs the routes of the server that are registered within Server.registerRoutes(). This is useful for debugging.
// See https://github.com/go-chi/chi/issues/332 for details and inspiration.
func (a *app) LogRoutes(r *chi.Mux) {
	desiredSpaces := 10

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
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

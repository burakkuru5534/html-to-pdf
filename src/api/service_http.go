package api

import (
	"github.com/burakkuru5534/src/api/register"
	"github.com/burakkuru5534/src/api/sys"
	"github.com/burakkuru5534/src/helper"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func HttpService() http.Handler {
	mux := chi.NewRouter()

	acors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	mux.Use(acors.Handler)

	mux.Route("/api", func(mr chi.Router) {
		mr.Group(func(r chi.Router) {
			//r.Get("/users/{id}", UserGet)
			r.Post("/login", sys.Login)
			r.Post("/register", register.NewRegister)

		})
		//protected end points
		mr.Group(func(r chi.Router) {
			//Token Middleware
			r.Use(jwtauth.Verifier(helper.Conf.Auth.JWTAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/document", DocumentCreate)
			r.Get("/document", DocumentGet)
			r.Delete("/document", DocumentDelete)
			r.Get("/documents", DocumentList)
			r.Get("/document/download", DocumentDownload)
		})
	})

	return mux
}

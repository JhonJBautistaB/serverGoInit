package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Routes struct{}

func (r *Routes) Routes() http.Handler {
	mux := chi.NewRouter()

	// especify CORS
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// especify Middleware
	mux.Use(middleware.Heartbeat("/ping"))

	//Version API
	mux.Route("/v1", func(mux chi.Router) {
		// routes API
		mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to the API V1"))
		})
	})

	mux.Route("/v2", func(mux chi.Router) {
		// routes API
		mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to the API V2"))
		})
	})

	return mux
}

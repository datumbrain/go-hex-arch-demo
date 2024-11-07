package http

import (
	"fmt"
	"hexashop/internal/adapter/http/api"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	*chi.Mux
}

func NewRouter(userApi *api.User, productApi *api.Product) *Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin", "Cache-Control"},
		ExposedHeaders:   []string{"JWT-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userApi.CreateUser)
			r.Get("/{id}", userApi.GetUser)
			r.Patch("/{id}", userApi.UpdateUser)
			r.Delete("/{id}", userApi.DeleteUser)
		})

		r.Route("/products", func(r chi.Router) {
			r.Post("/", productApi.CreateProduct)
			r.Get("/{id}", productApi.GetProduct)
			r.Patch("/{id}", productApi.UpdateProduct)
			r.Delete("/{id}", productApi.DeleteProduct)
		})
	})

	return &Router{r}
}

func (r *Router) Servce(port string) error {
	listenAddr := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(listenAddr, r)
}

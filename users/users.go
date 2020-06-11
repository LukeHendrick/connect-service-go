package users

import (
	"github.com/go-chi/chi"
)

// Routes for the ContactFlow Operations
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Put("/{instanceID}", CreateUser)
	return router
}

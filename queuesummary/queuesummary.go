package queuesummary

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{instanceID}", ListQueues)
	return router
}

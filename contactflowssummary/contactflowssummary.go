package contactflowssummary

import (
	"github.com/go-chi/chi"
)

// Routes for the ContactFlow Operations
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{instanceId}", ListContactFlows)
	return router
}

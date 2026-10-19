// Package api is the HTTP layer of the service.
package api

import (
	"net/http"
	"time"

	"github.com/AndraN20/personal-finance-tracker/internal/store"
)

type Server struct {
	store   *store.Store
	started time.Time
	mux     *http.ServeMux
}

func NewServer(st *store.Store) *Server {
	s := &Server{store: st, started: time.Now(), mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) routes() {
	// Common to every project.
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("GET /version", s.handleVersion)
	s.mux.HandleFunc("GET /{$}", s.handleIndex)
	s.mux.HandleFunc("POST /reset", s.handleReset)

	// Specific to this project.
	s.mux.HandleFunc("POST /transactions", s.handleCreateTransaction)
	s.mux.HandleFunc("POST /transactions/import", s.handleImportTransactions)
	s.mux.HandleFunc("GET /transactions", s.handleListTransactions)
	s.mux.HandleFunc("PUT /rules", s.handlePutRules)
	s.mux.HandleFunc("PUT /budgets", s.handlePutBudget)
	s.mux.HandleFunc("GET /reports/monthly", s.handleMonthlyReport)

	// Anything else gets a JSON 404 instead of the default plain-text one.
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusNotFound, codeNotFound, "no route for "+r.Method+" "+r.URL.Path)
	})
}

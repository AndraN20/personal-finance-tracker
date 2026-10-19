package api

import "net/http"

// The project-specific routes are registered but not implemented yet.
// Each answers 501 with the standard error body until it is written.

func notImplemented(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, codeNotImplemented, r.Method+" "+r.URL.Path+" is not implemented yet")
}

func (s *Server) handleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Server) handleImportTransactions(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Server) handleListTransactions(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Server) handlePutRules(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Server) handlePutBudget(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Server) handleMonthlyReport(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

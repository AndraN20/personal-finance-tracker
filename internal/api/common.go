package api

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/AndraN20/personal-finance-tracker/internal/buildinfo"
)

// Shown on GET /.
const (
	authorName  = "Andra Nichifor"
	groupName   = "DMWD"
	projectName = "Personal Finance Tracker"
)

var indexPage = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.App}}</title>
</head>
<body>
<h1>{{.App}}</h1>
<p>{{.Project}}</p>
<p>Author: {{.Author}}</p>
<p>Group: {{.Group}}</p>
<p>Version {{.Version}} ({{.Commit}})</p>
</body>
</html>
`))

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":         "ok",
		"uptime_seconds": int64(time.Since(s.started).Seconds()),
	})
}

func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"app":      buildinfo.App,
		"version":  buildinfo.Version,
		"commit":   buildinfo.Commit,
		"built_at": buildinfo.BuiltAt,
	})
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := indexPage.Execute(w, map[string]string{
		"App":     buildinfo.App,
		"Project": projectName,
		"Author":  authorName,
		"Group":   groupName,
		"Version": buildinfo.Version,
		"Commit":  buildinfo.Commit,
	})
	if err != nil {
		log.Printf("render index: %v", err)
	}
}

func (s *Server) handleReset(w http.ResponseWriter, r *http.Request) {
	s.store.Reset()
	w.WriteHeader(http.StatusNoContent)
}

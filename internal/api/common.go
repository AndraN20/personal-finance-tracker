package api

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"atad-project/internal/version"
)

// Shown on GET /, in the format the ATAD template uses. Must match atad.json.
const (
	authorName    = "Nichifor Andra"
	groupName     = "DMWD"
	projectNumber = "2"
	projectName   = "Personal Finance Tracker"
)

var indexPage = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.App}}</title>
</head>
<body>
<h1>{{.App}}</h1>
<p>{{.Author}}, group {{.Group}}</p>
<p>Project {{.Number}}: {{.Project}}</p>
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
	writeJSON(w, http.StatusOK, version.Current())
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	v := version.Current()
	err := indexPage.Execute(w, map[string]string{
		"App":     v.App,
		"Number":  projectNumber,
		"Project": projectName,
		"Author":  authorName,
		"Group":   groupName,
		"Version": v.Version,
		"Commit":  v.Commit,
	})
	if err != nil {
		log.Printf("render index: %v", err)
	}
}

func (s *Server) handleReset(w http.ResponseWriter, r *http.Request) {
	s.store.Reset()
	w.WriteHeader(http.StatusNoContent)
}

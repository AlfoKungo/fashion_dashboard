package web

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"fashion_dashboard/internal/processing"
	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/scheduler"
)

type Server struct {
	mux       *http.ServeMux
	store     repository.Store
	dashboard *processing.DashboardService
	workflow  scheduler.Runner
	templates *template.Template
}

func NewServer(store repository.Store) (*Server, error) {
	return NewServerWithWorkflow(store, scheduler.NewWorkflow(store))
}

func NewServerWithWorkflow(store repository.Store, workflow scheduler.Runner) (*Server, error) {
	templateGlob, err := projectPath("internal/web/templates/*.html")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.ParseGlob(templateGlob)
	if err != nil {
		return nil, err
	}
	s := &Server{
		mux:       http.NewServeMux(),
		store:     store,
		dashboard: processing.NewDashboardService(store),
		workflow:  workflow,
		templates: tmpl,
	}
	s.routes()
	return s, nil
}

func (s *Server) Handler() http.Handler {
	return s.mux
}

func (s *Server) routes() {
	staticDir, _ := projectPath("internal/web/static")
	s.mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	s.mux.HandleFunc("GET /", s.TodayPage)
	s.mux.HandleFunc("GET /looks", s.LooksPage)
	s.mux.HandleFunc("GET /items", s.ItemsPage)
	s.mux.HandleFunc("GET /articles", s.ArticlesPage)
	s.mux.HandleFunc("GET /api/articles", s.ArticlesAPI)
	s.mux.HandleFunc("GET /api/looks", s.LooksAPI)
	s.mux.HandleFunc("GET /api/items", s.ItemsAPI)
	s.mux.HandleFunc("POST /api/update", s.UpdateAPI)
	s.mux.HandleFunc("GET /images/articles/{id}", s.ArticleImage)
	s.mux.HandleFunc("GET /images/looks/{id}", s.LookImage)
	s.mux.HandleFunc("GET /images/items/{id}", s.ItemImage)
}

func (s *Server) render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectPath(rel string) (string, error) {
	if _, err := os.Stat(rel); err == nil || hasGlobMatch(rel) {
		return rel, nil
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return rel, nil
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	return filepath.Join(root, rel), nil
}

func hasGlobMatch(pattern string) bool {
	matches, err := filepath.Glob(pattern)
	return err == nil && len(matches) > 0
}

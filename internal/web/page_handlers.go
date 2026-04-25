package web

import (
	"net/http"
)

type ListPageData struct {
	ActivePage string
	DateLabel  string
	Title      string
	Articles   any
	Looks      any
	Items      any
}

func (s *Server) TodayPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data, err := s.dashboard.Today(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.render(w, "today.html", data)
}

func (s *Server) LooksPage(w http.ResponseWriter, r *http.Request) {
	looks, err := s.dashboard.Looks(r.Context(), 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.render(w, "looks.html", ListPageData{ActivePage: "looks", Title: "Looks", Looks: looks})
}

func (s *Server) ItemsPage(w http.ResponseWriter, r *http.Request) {
	items, err := s.dashboard.Items(r.Context(), 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.render(w, "items.html", ListPageData{ActivePage: "items", Title: "Items", Items: items})
}

func (s *Server) ArticlesPage(w http.ResponseWriter, r *http.Request) {
	articles, err := s.dashboard.Articles(r.Context(), 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.render(w, "articles.html", ListPageData{ActivePage: "articles", Title: "Articles", Articles: articles})
}

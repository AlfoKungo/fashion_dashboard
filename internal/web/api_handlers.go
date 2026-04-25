package web

import (
	"net/http"
	"time"

	"fashion_dashboard/internal/processing"
)

func (s *Server) ArticlesAPI(w http.ResponseWriter, r *http.Request) {
	amount, err := ParseAmount(r, 4)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	articles, err := s.dashboard.Articles(r.Context(), amount)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"articles": articles})
}

func (s *Server) LooksAPI(w http.ResponseWriter, r *http.Request) {
	amount, err := ParseAmount(r, 4)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	looks, err := s.dashboard.Looks(r.Context(), amount)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"looks": looks})
}

func (s *Server) ItemsAPI(w http.ResponseWriter, r *http.Request) {
	amount, err := ParseAmount(r, 6)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	items, err := s.dashboard.Items(r.Context(), amount)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	category := processing.DailyCategory("")
	if len(items) > 0 {
		category = items[0].Category
	}
	writeJSON(w, http.StatusOK, map[string]any{"category": category, "items": items})
}

func (s *Server) UpdateAPI(w http.ResponseWriter, r *http.Request) {
	if err := s.workflow.Run(r.Context()); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	date := time.Now().Format("2006-01-02")
	writeJSON(w, http.StatusOK, map[string]any{
		"status":   "updated",
		"date":     date,
		"category": processing.DailyCategory(date),
	})
}

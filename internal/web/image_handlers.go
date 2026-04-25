package web

import (
	"net/http"

	"fashion_dashboard/internal/models"
)

func (s *Server) ArticleImage(w http.ResponseWriter, r *http.Request) {
	image, ok, err := s.store.GetArticleImage(r.Context(), r.PathValue("id"))
	serveImage(w, r, image, ok, err)
}

func (s *Server) LookImage(w http.ResponseWriter, r *http.Request) {
	image, ok, err := s.store.GetLookImage(r.Context(), r.PathValue("id"))
	serveImage(w, r, image, ok, err)
}

func (s *Server) ItemImage(w http.ResponseWriter, r *http.Request) {
	image, ok, err := s.store.GetItemImage(r.Context(), r.PathValue("id"))
	serveImage(w, r, image, ok, err)
}

func serveImage(w http.ResponseWriter, r *http.Request, image models.Image, ok bool, err error) {
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !ok {
		writeError(w, http.StatusNotFound, "image not found")
		return
	}
	if len(image.Bytes) > 0 {
		contentType := image.ContentType
		if contentType == "" {
			contentType = http.DetectContentType(image.Bytes)
		}
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(image.Bytes)
		return
	}
	if image.URL != "" {
		http.Redirect(w, r, image.URL, http.StatusFound)
		return
	}
	writeError(w, http.StatusNotFound, "image not found")
}

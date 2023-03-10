package articles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"

	"github.com/xavimg/articles/internal/services/articles"
)

type Server struct {
	Service articles.Service
}

func NewServer(router *chi.Mux) {
	s := &Server{
		Service: *articles.NewService(),
	}
	s.registerEndpoints(router)
}

func (s *Server) registerEndpoints(router *chi.Mux) {
	router.Get("/teams/{team}/news", s.List)
	router.Get("/teams/{team}/news/{id}", s.Get)
}

func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	team := chi.URLParam(r, "team")
	res, err := s.Service.List(context.Background(), team)
	if err != nil {
		logrus.Error(err)
		errorResponse(w, err)
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	team := chi.URLParam(r, "team")
	id := chi.URLParam(r, "id")
	res, err := s.Service.Get(context.Background(), team, id)
	if err != nil {
		logrus.Error(err)
		errorResponse(w, err)
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type apiError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func errorResponse(w http.ResponseWriter, err error) {
	errorApi := &apiError{
		Status:  "error",
		Message: err.Error(),
	}

	jujuErr, ok := err.(*errors.Err)
	if !ok {
		// jujuErr library not includes http.StatusInternalServerError, so we can't check it out in switch.
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorApi)
		return
	}

	status := http.StatusInternalServerError
	switch {
	case errors.Is(jujuErr, errors.NotFound):
		status = http.StatusNotFound
	case errors.Is(jujuErr, errors.NotValid), errors.Is(jujuErr, errors.BadRequest):
		status = http.StatusBadRequest
	case errors.Is(jujuErr, errors.Unauthorized):
		status = http.StatusUnauthorized
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorApi)
}

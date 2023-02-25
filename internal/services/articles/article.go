package articles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
)

type Server struct {
	articles *ArticlesJSON
	article  *ArticleJSON
}

func NewServer(router *chi.Mux) {
	s := &Server{}
	s.registerEndpoints(router)
}

func (s *Server) registerEndpoints(router *chi.Mux) {
	router.Get("/teams/t94/news", makeHTTPHandleFunc(s.All))
	router.Get("/teams/t94/news/{id}", makeHTTPHandleFunc(s.ByID))
}

func (ah *Server) All(w http.ResponseWriter, r *http.Request) error {
	articles, err := models.Repo.GetAll(context.Background())
	if err != nil {
		logrus.Errorf("err %s\n", err)
		return err
	}

	return writeJSON(w, http.StatusOK, ah.articles.Serialize(articles))
}

func (ah *Server) ByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	article, err := models.Repo.GetByID(context.Background(), id)
	if err != nil {
		logrus.Errorf("err %s\n", err)
		return err
	}

	return writeJSON(w, http.StatusOK, ah.article.Serialize(article))
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			logrus.Errorf(" %s\n", err)
			writeJSON(w, http.StatusBadRequest, err)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

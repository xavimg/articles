package handlers

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/incrowd-test/internal/models"
	"github.com/xavimg/incrowd-test/internal/store"
)

type APIServer struct {
	listenAddr string
	store      store.Storer
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Get("/all-items", makeHTTPHandleFunc(s.AllArticles))

	http.ListenAndServe(s.listenAddr, router)
}

func NewArticleHandler(addr string, aStore store.Storer) *APIServer {
	return &APIServer{
		listenAddr: addr,
		store:      aStore,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println("xd")
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (ah *APIServer) AllArticles(w http.ResponseWriter, r *http.Request) error {

	// Show which come from feed provider.
	ctx := context.Background()
	articles, err := ah.store.GetAll(ctx)
	if err != nil {
		return err
	}

	// ah.store.Insert()

	return WriteJSON(w, http.StatusOK, articles)
}

func callAllItems() *models.DataJSONAll {
	resp, err := http.Get("https://www.wearehullcity.co.uk/api/incrowd/getnewlistinformation?count=50")
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLALL
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil
	}

	var body models.BodyAll
	var bodys []models.BodyAll
	for _, item := range dataClean.NewsletterNewsItems.NewsletterNewsItem {
		body = models.BodyAll{
			ArticleURL:        item.ArticleURL,
			NewsArticleID:     item.NewsArticleID,
			PublishDate:       item.PublishDate,
			Taxonomies:        item.Taxonomies,
			TeaserText:        item.TeaserText,
			ThumbnailImageURL: item.ThumbnailImageURL,
			Title:             item.Title,
			OptaMatchId:       item.OptaMatchId,
			LastUpdateDate:    item.LastUpdateDate,
			IsPublished:       item.IsPublished,
		}

		bodys = append(bodys, body)
	}
	responsejson := &models.DataJSONAll{
		Status: "succes",
	}
	responsejson.Data = bodys

	return responsejson
}

package articles

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/xavimg/articles/internal/models"
)

const (
	host         = "www.wearehullcity.co.uk"
	endpointList = "/api/incrowd/getnewlistinformation"
	endpointOne  = "/api/incrowd/getnewsarticleinformation"
)

type Server struct {
}

func NewServer(router *chi.Mux) {

	s := &Server{}
	s.registerEndpoints(router)

}

func (s *Server) registerEndpoints(router *chi.Mux) {
	router.Get("/teams/t94/news/{count}", makeHTTPHandleFunc(s.All))
	router.Get("/teams/t94/new/{id}", makeHTTPHandleFunc(s.ByID))
}

func (ah *Server) All(w http.ResponseWriter, r *http.Request) error {
	// count := chi.URLParam(r, "count")

	// Fetch coming from provider-feed.
	// var savedArticles []models.Article
	// if count == "5" {
	// 	providerArticles, err := providerArticles(count)
	// 	if err != nil {
	// 		// Alternative fetch if our provider-feed fails.
	// 		// savedArticles, err = ah.repo.GetAll(context.Background())
	// 		// if err != nil {
	// 		// 	logrus.Printf("err %s\n", err)
	// 		// 	return err
	// 		// }
	// 		// logrus.Println(err)
	// 		// return writeJSON(w, http.StatusOK, savedArticles)
	// 	}

	// 	insertArticles := []*models.Article{}
	// 	for _, article := range providerArticles.Data {
	// 		insertArticles = append(insertArticles, &models.Article{
	// 			ID:                primitive.NewObjectID(),
	// 			ArticleURL:        article.ArticleURL,
	// 			NewsArticleID:     article.NewsArticleID,
	// 			PublishDate:       article.PublishDate,
	// 			Taxonomies:        article.Taxonomies,
	// 			TeaserText:        article.TeaserText,
	// 			ThumbnailImageURL: article.ThumbnailImageURL,
	// 			Title:             article.Title,
	// 			OptaMatchId:       article.OptaMatchId,
	// 			LastUpdateDate:    article.LastUpdateDate,
	// 			IsPublished:       article.IsPublished,
	// 		})
	// 	}

	// 	go func() error {
	// 		if err := models.Repo.InsertMany(context.Background(), insertArticles); err != nil {
	// 			logrus.Println(err)
	// 			return err
	// 		}
	// 		return nil
	// 	}()
	// 	return writeJSON(w, http.StatusOK, providerArticles)

	// } else {
	savedArticles, err := models.Repo.GetAll(context.Background())
	if err != nil {
		logrus.Printf("err %s\n", err)
		return err
	}

	var resp ArticlesJSON
	for _, article := range savedArticles {
		resp = ArticlesJSON{Status: "succes"}
		resp.Data = append(resp.Data, Article{
			ArticleURL:        article.ArticleURL,
			NewsArticleID:     article.NewsArticleID,
			PublishDate:       article.PublishDate,
			Taxonomies:        article.Taxonomies,
			TeaserText:        article.TeaserText,
			ThumbnailImageURL: article.ThumbnailImageURL,
			Title:             article.Title,
			OptaMatchId:       article.OptaMatchId,
			LastUpdateDate:    article.LastUpdateDate,
			IsPublished:       article.IsPublished,
		})
	}

	return writeJSON(w, http.StatusOK, resp)
}

// }

func (ah *Server) ByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	// Fetch coming from provider-feed.
	// var savedArticle models.Article
	if id == "5" {
		providerArticle, err := providerArticle(id)
		if err != nil {
			logrus.Error(err)
			return err
		}

		fmt.Println("response provider:", providerArticle)

		insertArticles := &models.Article{
			ArticleURL:        providerArticle.Data.ArticleURL,
			NewsArticleID:     providerArticle.Data.NewsArticleID,
			PublishDate:       providerArticle.Data.PublishDate,
			Taxonomies:        providerArticle.Data.Taxonomies,
			TeaserText:        providerArticle.Data.TeaserText,
			ThumbnailImageURL: providerArticle.Data.ThumbnailImageURL,
			Title:             providerArticle.Data.Title,
			OptaMatchId:       providerArticle.Data.OptaMatchId,
			LastUpdateDate:    providerArticle.Data.LastUpdateDate,
			IsPublished:       providerArticle.Data.IsPublished,
		}

		go func() error {
			if err := models.Repo.InsertOne(context.Background(), insertArticles); err != nil {
				logrus.Println(err)
				return err
			}
			return nil
		}()
		return writeJSON(w, http.StatusOK, providerArticles)

	} else {
		savedArticle, err := models.Repo.GetByID(context.Background(), id)
		if err != nil {
			logrus.Printf("err %s\n", err)
			return err
		}

		resp := &ArticleJSON{
			Status: "Succes",
			Data: Article{
				ArticleURL:        savedArticle.ArticleURL,
				PublishDate:       savedArticle.PublishDate,
				Taxonomies:        savedArticle.Taxonomies,
				TeaserText:        savedArticle.TeaserText,
				ThumbnailImageURL: savedArticle.ThumbnailImageURL,
				Title:             savedArticle.Title,
				OptaMatchId:       savedArticle.OptaMatchId,
				LastUpdateDate:    savedArticle.LastUpdateDate,
				IsPublished:       savedArticle.IsPublished,
			},
		}

		return writeJSON(w, http.StatusOK, resp)
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			logrus.Println("xd")
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func providerArticles(count string) (*ArticlesJSON, error) {
	params := url.Values{}
	params.Add("count", count)
	url := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     endpointList,
		RawQuery: params.Encode(),
	}
	resp, err := http.Get(url.String())
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLALL
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	var article Article
	var articles []Article
	for _, item := range dataClean.NewsletterNewsItems.NewsletterNewsItem {
		article = Article{
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

		articles = append(articles, article)
	}

	responsejson := &ArticlesJSON{
		Status: "succes",
		Data:   articles,
	}

	return responsejson, nil
}

func providerArticle(id string) (*ArticleJSON, error) {
	params := url.Values{}
	params.Add("id", id)
	url := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     endpointList,
		RawQuery: params.Encode(),
	}
	resp, err := http.Get(url.String())
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	var dataClean models.DataXMLOne
	if err := xml.Unmarshal(respByte, &dataClean); err != nil {
		logrus.Printf("Error %s", err)
		return nil, err
	}

	body := Article{
		ArticleURL:        dataClean.NewsletterNewsItem.ArticleURL,
		NewsArticleID:     dataClean.NewsletterNewsItem.NewsArticleID,
		PublishDate:       dataClean.NewsletterNewsItem.PublishDate,
		Taxonomies:        dataClean.NewsletterNewsItem.Taxonomies,
		TeaserText:        dataClean.NewsletterNewsItem.TeaserText,
		ThumbnailImageURL: dataClean.NewsletterNewsItem.ThumbnailImageURL,
		Title:             dataClean.NewsletterNewsItem.Title,
		OptaMatchId:       dataClean.NewsletterNewsItem.OptaMatchId,
		LastUpdateDate:    dataClean.NewsletterNewsItem.LastUpdateDate,
		IsPublished:       dataClean.NewsletterNewsItem.IsPublished,
	}
	responsejson := &ArticleJSON{
		Status: "succes",
		Data:   body,
	}

	return responsejson, nil
}

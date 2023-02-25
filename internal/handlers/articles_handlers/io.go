package articles_handlers

import (
	"fmt"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ArticlesJSON struct {
	Status string    `json:"status"`
	Data   []Article `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type ArticleJSON struct {
	Status string  `json:"status"`
	Data   Article `json:"data"`
	// Metadata Metadata `json:"metadata"`
}

type Article struct {
	ArticleURL        string `json:"articleURL"`
	NewsArticleID     string `json:"newsArticleID"`
	PublishDate       string `json:"publishDate"`
	Taxonomies        string `json:"taxonomies"`
	TeaserText        string `json:"teaserText"`
	ThumbnailImageURL string `json:"thumbnailImageURL"`
	Title             string `json:"title"`
	OptaMatchId       string `json:"optaMatchID"`
	LastUpdateDate    string `json:"lastUpdateDate"`
	IsPublished       string `json:"published"`
}

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println("xd")
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

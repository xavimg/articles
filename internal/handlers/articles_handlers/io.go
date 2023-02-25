package articles_handlers

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

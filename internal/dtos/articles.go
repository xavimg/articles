package dtos

import "time"

type ListReply struct {
	Status   string     `json:"status"`
	Data     []*Article `json:"data"`
	Metadata *Metadata  `json:"metadata"`
}

type GetReply struct {
	Status   string    `json:"status"`
	Data     *Article  `json:"data"`
	Metadata *Metadata `json:"metadata"`
}

type Metadata struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems,omitempty"`
	Sort       string    `json:"sort,omitempty"`
}

type Article struct {
	ID             string    `json:"id"`
	TeamID         string    `json:"teamId"`
	OptaMatchId    string    `json:"optaMatchID"`
	Title          string    `json:"title"`
	Type           []string  `json:"type"`
	TeaserText     string    `json:"teaserText"`
	Content        string    `json:"content"`
	ArticleURL     string    `json:"articleURL"`
	ImageUrl       string    `json:"imageUrl"`
	GalleryURLs    []string  `json:"galletyUrls"`
	VideoUrl       string    `json:"videoUrl"`
	NewsArticleID  string    `json:"newsArticleID"`
	PublishDate    time.Time `json:"publishDate"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
	IsPublished    bool      `json:"published"`
}

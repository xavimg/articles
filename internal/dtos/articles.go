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
	ID                string    `json:"id"`
	TeamID            string    `json:"teamId"`
	ArticleURL        string    `json:"articleURL"`
	NewsArticleID     string    `json:"newsArticleID"`
	PublishDate       time.Time `json:"publishDate"`
	Type              []string  `json:"type"`
	TeaserText        string    `json:"teaserText"`
	ThumbnailImageURL string    `json:"thumbnailImageURL"`
	Title             string    `json:"title"`
	OptaMatchId       string    `json:"optaMatchID"`
	LastUpdateDate    time.Time `json:"lastUpdateDate"`
	IsPublished       bool      `json:"published"`
}

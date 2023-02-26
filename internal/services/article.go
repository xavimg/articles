package services

import (
	"context"
	"time"

	"github.com/juju/errors"
	"github.com/xavimg/articles/internal/dtos"
	"github.com/xavimg/articles/internal/models"
)

type Service struct{}

func NewService() *Service {
	return new(Service)
}

func serializeArticle(article *models.Article) *dtos.Article {
	return &dtos.Article{
		ID:                article.ID.Hex(),
		TeamID:            article.TeamID,
		ArticleURL:        article.ArticleURL,
		NewsArticleID:     article.NewsArticleID,
		PublishDate:       article.PublishDate,
		Type:              article.Type,
		TeaserText:        article.TeaserText,
		ThumbnailImageURL: article.ThumbnailImageURL,
		Title:             article.Title,
		OptaMatchId:       article.OptaMatchId,
		LastUpdateDate:    article.LastUpdateDate,
		IsPublished:       article.IsPublished,
	}
}

func (ah *Service) List(ctx context.Context, team string) (*dtos.ListReply, error) {
	articles, err := models.Repo.List(ctx, team)
	if err != nil {
		return nil, errors.Trace(err)
	}

	reply := &dtos.ListReply{
		Status: "succes",
		Metadata: &dtos.Metadata{
			CreatedAt:  time.Now().Format("2006-01-02"),
			TotalItems: len(articles),
			Sort:       "-published",
		},
	}
	for _, article := range articles {
		reply.Data = append(reply.Data, serializeArticle(article))
	}

	return reply, nil
}

func (ah *Service) Get(ctx context.Context, id, team string) (*dtos.Article, error) {
	article, err := models.Repo.Get(ctx, id, team)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return serializeArticle(article), nil
}

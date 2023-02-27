package articles

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
		ID:          article.ID.Hex(),
		TeamID:      article.TeamID,
		OptaMatchId: article.OptaMatchId,
		Title:       article.Title,
		Type:        article.Type,
		TeaserText:  article.TeaserText,
		Content:     article.Content,
		ArticleURL:  article.ArticleURL,
		ImageUrl:    article.ArticleURL,
		GalleryURLs: article.GalleryURLs,
		VideoUrl:    article.VideoURL,
		PublishDate: article.PublishDate,
	}
}

func (ah *Service) List(ctx context.Context, teamId string) (*dtos.ListReply, error) {
	articles, err := models.Repo.ListArticles(ctx, teamId)
	if err != nil {
		return nil, errors.Trace(err)
	}

	reply := &dtos.ListReply{
		Status: "succes",
		Metadata: &dtos.Metadata{
			CreatedAt:  time.Now(),
			TotalItems: len(articles),
			Sort:       "-published", // It means we order articles in DESC way in repo package using --> options.Find().SetSort(bson.M{"lastUpdateDate": -1})
		},
	}
	for _, article := range articles {
		reply.Data = append(reply.Data, serializeArticle(article))
	}

	return reply, nil
}

func (ah *Service) Get(ctx context.Context, teamId, id string) (*dtos.GetReply, error) {
	article, err := models.Repo.GetArticle(ctx, teamId, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	reply := &dtos.GetReply{
		Status: "succes",
		Data:   serializeArticle(article),
		Metadata: &dtos.Metadata{
			CreatedAt: time.Now(),
		},
	}

	return reply, nil
}

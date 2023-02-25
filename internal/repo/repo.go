package repo

import (
	"context"

	"github.com/xavimg/articles/internal/models"
)

type Repo interface {
	InsertMany(context.Context, []models.Article) error
	GetAll(context.Context) ([]*models.Article, error)
	GetByID(context.Context, string) (*models.Article, error)
}

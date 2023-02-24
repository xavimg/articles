package store

import (
	"context"

	"github.com/xavimg/incrowd-test/internal/models"
)

type Storer interface {
	Insert(context.Context, *models.Article) error
	GetAll(context.Context) ([]*models.Article, error)
	GetByID(context.Context, string) (*models.Article, error)
}

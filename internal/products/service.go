package products

import (
	"context"

	repo "github.com/sorinqu-org/go-backend-api/internal/adapters/sqlc"
)

type service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProduct(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProduct(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

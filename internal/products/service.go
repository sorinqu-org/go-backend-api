package products

import (
	"context"

	repo "github.com/sorinqu-org/go-backend-api/internal/adapters/sqlc"
)

type service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
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

package products

import (
	"context"
)

type service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	// TODO: add repository
}

func (s *svc) ListProducts(ctx context.Context) error  {
	return nil
}

func NewService() service {
	return &svc{}
}

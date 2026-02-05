package orders

import (
	"context"

	repo "github.com/sorinqu-org/go-backend-api/internal/adapters/sqlc"
)

type service interface {
	PlaceOrder(
		ctx context.Context,
		customerID int64,
		items []OrderItemParams,
	) (int64, error)

	AddOrderItem(
		ctx context.Context,
		order_id int64,
		item OrderItemParams,
	) (int64, error)
	GetOrderByID(ctx context.Context, id int64) (repo.Order, error)
	GetItemByID(ctx context.Context, id int64) (repo.OrderItem, error)
	DeleteOrderByID(ctx context.Context, id int64) error
	ListOrders(ctx context.Context) ([]repo.Order, error)
	// TODO: add get item by id, list orders and items for order
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) service {
	return &svc{repo: repo}
}

func (s *svc) PlaceOrder(
	ctx context.Context,
	customerID int64,
	items []OrderItemParams,
) (int64, error) {
	order_id, err := s.repo.PlaceOrder(ctx, customerID)
	if err != nil {
		return 0, err
	}

	for _, item := range items {

		if _, err := s.AddOrderItem(ctx, order_id, OrderItemParams{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}); err != nil {
			s.DeleteOrderByID(ctx, order_id)
			return 0, err
		}

		if _, err := s.repo.ChangeProductQuantity(ctx, repo.ChangeProductQuantityParams{
			ID:       item.ProductID,
			Quantity: item.Quantity,
		}); err != nil {
			s.DeleteOrderByID(ctx, order_id)
			return 0, err
		}
	}

	return order_id, nil
}

func (s *svc) AddOrderItem(
	ctx context.Context,
	order_id int64,
	item OrderItemParams,
) (int64, error) {
	product, err := s.repo.FindProductByID(ctx, item.ProductID)
	if err != nil {
		return 0, err
	}

	return s.repo.AddOrderItem(ctx, repo.AddOrderItemParams{
		OrderID:    order_id,
		ProductID:  item.ProductID,
		Quantity:   item.Quantity,
		PriceInUsd: product.PriceInUsd * item.Quantity,
	})
}

func (s *svc) GetItemByID(
	ctx context.Context,
	id int64,
) (repo.OrderItem, error) {
	return s.repo.GetItemByID(ctx, id)
}

func (s *svc) GetOrderByID(ctx context.Context, id int64) (repo.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *svc) DeleteOrderByID(ctx context.Context, id int64) error {
	if err := s.repo.DeleteOrderItemsByOrderID(ctx, id); err != nil {
		return err
	}
	return s.repo.DeleteOrderByID(ctx, id)
}

func (s *svc) ListOrders(ctx context.Context) ([]repo.Order, error) {
	return s.repo.ListOrders(ctx)
}

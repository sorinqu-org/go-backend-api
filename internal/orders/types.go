package orders

type PlaceOrderBodyParams struct {
	CustomerID int64             `json:"customer_id"`
	Items      []OrderItemParams `json:"items"`
}

type OrderItemParams struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

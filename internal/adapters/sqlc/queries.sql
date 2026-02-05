-- name: ListProducts :many
SELECT *
FROM products;

-- name: FindProductByID :one
SELECT *
FROM products
WHERE id = $1;

-- name: PlaceOrder :one
INSERT INTO orders (customer_id)
VALUES ($1)
RETURNING id;

-- name: AddOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity, price_in_usd)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetOrderByID :one
SELECT *
FROM orders
WHERE id = $1;

-- name: ListOrders :many
SELECT *
FROM orders;

-- name: GetItemByID :one
SELECT *
from order_items
WHERE id = $1;

-- name: DeleteOrderItemsByOrderID :exec
DELETE FROM order_items
WHERE order_id = $1;

-- name: DeleteOrderByID :exec
DELETE FROM orders
WHERE id = $1;

-- name: ChangeProductQuantity :one
UPDATE products
SET quantity = quantity - $2
WHERE id = $1
  AND quantity >= $2
RETURNING quantity;

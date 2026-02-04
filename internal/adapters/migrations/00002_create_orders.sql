-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT now ()
  );

CREATE TABLE
  IF NOT EXISTS order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGSERIAL NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL,
    price_in_usd INTEGER NOT NULL,
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id)
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;

DROP TABLE IF EXISTS orders;

-- +goose StatementEnd

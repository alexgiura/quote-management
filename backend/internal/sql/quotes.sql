-- name: CreateQuote :one
INSERT INTO core.quotes (quote_code, customer_id, status, total_amount, currency)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, quote_code, customer_id, status, total_amount, currency, created_at, updated_at;

-- name: AddProductToQuote :exec
INSERT INTO core.quote_products (quote_id, product_id, quantity, price_per_unit, tax_rate, total_price)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetQuoteByID :one
SELECT
    id AS quote_id,
    quote_code,
    customer_id,
    status,
    total_amount,
    currency,
    created_at,
    updated_at
FROM core.quotes
WHERE id = $1;


-- name: GetProductsByQuoteID :many
SELECT
    qp.product_id,
    p.name AS product_name,
    qp.quantity,
    qp.price_per_unit,
    qp.tax_rate,
    qp.total_price
FROM core.quote_products qp
         JOIN core.products p ON p.id = qp.product_id
WHERE qp.quote_id = $1;







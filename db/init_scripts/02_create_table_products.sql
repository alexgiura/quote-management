CREATE TABLE core.products (
                               id SERIAL PRIMARY KEY,
                               product_code VARCHAR(255) NOT NULL,
                               name VARCHAR(255) NOT NULL,
                               created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                               updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

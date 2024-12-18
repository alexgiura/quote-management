CREATE TABLE core.quotes (
                             id SERIAL PRIMARY KEY,
                             quote_code VARCHAR(50) UNIQUE,
                             customer_id INT NOT NULL REFERENCES core.customers(id),
                             status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
                             total_amount  double precision NOT NULL DEFAULT 0.0,
                             currency VARCHAR(10) NOT NULL DEFAULT 'RON',
                             created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                             updated_at TIMESTAMP NOT NULL DEFAULT NOW()

);

CREATE TABLE core.quote_products (
                                     id SERIAL PRIMARY KEY,
                                     quote_id INT NOT NULL REFERENCES core.quotes (id),
                                     product_id INT NOT NULL REFERENCES core.products (id),
                                     quantity INT NOT NULL,
                                     price_per_unit double precision NOT NULL DEFAULT 0.0,
                                     tax_rate double precision NOT NULL DEFAULT 0.0,
                                     total_price double precision NOT NULL DEFAULT 0.0,
                                     created_at TIMESTAMP DEFAULT NOW()


);

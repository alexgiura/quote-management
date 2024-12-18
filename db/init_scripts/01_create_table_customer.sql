CREATE SCHEMA IF NOT EXISTS core;
CREATE TABLE core.customers (
                                id SERIAL PRIMARY KEY,
                                customer_code VARCHAR(50) NOT NULL UNIQUE,
                                name VARCHAR(255) NOT NULL,
                                email VARCHAR(255),
                                phone VARCHAR(20),
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);
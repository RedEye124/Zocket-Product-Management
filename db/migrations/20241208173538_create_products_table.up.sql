CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_images JSONB,
    compressed_product_images JSONB,
    product_price DECIMAL(10, 2)
);
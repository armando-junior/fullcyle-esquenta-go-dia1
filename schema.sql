CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL
);

-- Insert some sample data
INSERT INTO products (name, description, price) VALUES
    ('Product 1', 'Description for product 1', 99.99),
    ('Product 2', 'Description for product 2', 149.99),
    ('Product 3', 'Description for product 3', 199.99),
    ('Product 4', 'Description for product 4', 249.99),
    ('Product 5', 'Description for product 5', 299.99),
    ('Product 6', 'Description for product 6', 349.99),
    ('Product 7', 'Description for product 7', 399.99),
    ('Product 8', 'Description for product 8', 449.99),
    ('Product 9', 'Description for product 9', 499.99),
    ('Product 10', 'Description for product 10', 549.99),
    ('Product 11', 'Description for product 11', 599.99),
    ('Product 12', 'Description for product 12', 649.99),
    ('Product 13', 'Description for product 13', 699.99),
    ('Product 14', 'Description for product 14', 749.99),
    ('Product 15', 'Description for product 15', 799.99),
    

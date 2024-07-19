CREATE TABLE product (
                         product_id VARCHAR(50) PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT,
                         price NUMERIC(10, 2) NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

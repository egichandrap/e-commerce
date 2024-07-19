CREATE TABLE orders (
                       order_id VARCHAR(50) PRIMARY KEY,
                       customer_id VARCHAR(50) NOT NULL,
                       order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       status VARCHAR(50) NOT NULL,
                       total_amount NUMERIC(10, 2) NOT NULL,
                       CONSTRAINT fk_customer
                           FOREIGN KEY(customer_id)
                               REFERENCES customer(customer_id),
                       CHECK (total_amount >= 0)
);

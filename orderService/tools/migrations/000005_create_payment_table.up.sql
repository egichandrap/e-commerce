CREATE TABLE payment (
                         payment_id VARCHAR(50) PRIMARY KEY,
                         order_id VARCHAR(50) NOT NULL,
                         payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         amount NUMERIC(10, 2) NOT NULL,
                         payment_method VARCHAR(50) NOT NULL,
                         status VARCHAR(50) NOT NULL,
                         CONSTRAINT fk_order
                             FOREIGN KEY(order_id)
                                 REFERENCES orders(order_id),
                         CHECK (amount >= 0)
);

CREATE TABLE order_item (
                            order_item_id VARCHAR(50) PRIMARY KEY,
                            order_id VARCHAR(50) NOT NULL,
                            product_id VARCHAR(50) NOT NULL,
                            quantity INT NOT NULL,
                            price NUMERIC(10, 2) NOT NULL,
                            CONSTRAINT fk_order
                                FOREIGN KEY(order_id)
                                    REFERENCES orders(order_id),
                            CONSTRAINT fk_product
                                FOREIGN KEY(product_id)
                                    REFERENCES product(product_id),
                            CHECK (quantity > 0),
                            CHECK (price >= 0)
);

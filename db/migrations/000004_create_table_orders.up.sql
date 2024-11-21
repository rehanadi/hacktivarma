create table orders (
  id CHAR(8) DEFAULT SUBSTRING(REPLACE(uuid_generate_v4()::text, '-', ''), 1, 8),
  user_id CHAR(8) NOT NULL,
  drug_id CHAR(8) NOT NULL,
  quantity INT NOT NULL DEFAULT 0,
  price DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
  total_price DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
  payment_method VARCHAR(100),
  payment_status VARCHAR(100) NOT NULL DEFAULT 'unpaid',
  payment_at TIMESTAMP,
  delivery_status VARCHAR(100) NOT NULL DEFAULT 'pending',
  delivered_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_orders PRIMARY KEY(id),
  CONSTRAINT fk_order_user FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_order_drug FOREIGN KEY(drug_id) REFERENCES drugs(id)
);

INSERT INTO orders(user_id, drug_id, quantity, price, total_price, payment_method, payment_status, delivery_status)
VALUES
(
  (SELECT id FROM users WHERE email = 'user1@mail.com'),
  (SELECT id FROM drugs WHERE name = 'Paracetamol'),
  2,
  (SELECT price FROM drugs WHERE name = 'Paracetamol'),
  (SELECT price * 2 FROM drugs WHERE name = 'Paracetamol'),
  'Gopay',
  'unpaid',
  'pending'
);
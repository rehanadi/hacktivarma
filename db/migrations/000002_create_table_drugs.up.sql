create table drugs (
  id CHAR(8) DEFAULT SUBSTRING(REPLACE(uuid_generate_v4()::text, '-', ''), 1, 8),
  name VARCHAR(255) UNIQUE NOT NULL,
  form VARCHAR(255) NOT NULL,
  dose DECIMAL(10, 2) NOT NULL,
  stock INT DEFAULT 0,
  price DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
  expired_date DATE NOT NULL,
  category INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_drugs PRIMARY KEY(id),
  CONSTRAINT fk_drugs_categories FOREIGN KEY(category) REFERENCES categories(id)
);

INSERT INTO drugs (name, form, dose, stock, price, expired_date, category) VALUES 
('Paracetamol', 'Tablet', 500, 12, 5.0, '2025-06-01', 1),
('Amoxicillin', 'Kapsul', 500, 10, 20.0, '2025-07-15', 3),
('Ibuprofen', 'Tablet', 400, 8, 7.0, '2025-08-10', 2),
('Morfin', 'Tablet', 10, 22, 50.0, '2025-01-10', 7),
('Jahe', 'Kapsul', 1000, 10, 15.0, '2025-12-15', 5);
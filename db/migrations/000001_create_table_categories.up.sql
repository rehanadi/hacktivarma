CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

INSERT INTO categories (name) VALUES
('Obat Bebas'),
('Obat Bebas Terbatas'),
('Obat Keras'),
('Jamu'),
('Obat Herbal Terstandar'),
('Fitofarmaka'),
('Narkotika');
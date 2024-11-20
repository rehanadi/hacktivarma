CREATE TABLE categories (
   id INT auto_increment,
   name VARCHAR(100) NOT NULL,
   CONSTRAINT pk_categories PRIMARY KEY(id)
);

INSERT INTO categories (name) VALUES
('Obat Bebas'),
('Obat Bebas Terbatas'),
('Obat Keras'),
('Jamu'),
('Obat Herbal Terstandar'),
('Fitofarmaka'),
('Narkotika');
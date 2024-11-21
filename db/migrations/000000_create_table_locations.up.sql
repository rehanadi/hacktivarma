CREATE TABLE locations (
  id SERIAL,
  name VARCHAR(100) NOT NULL,
  CONSTRAINT pk_locations PRIMARY KEY(id)
);

INSERT INTO locations (name) VALUES
('jakarta'),
('bogor'),
('depok'),
('tangerang'),
('bekasi');
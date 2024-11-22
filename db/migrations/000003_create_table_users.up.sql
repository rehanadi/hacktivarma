CREATE TABLE users (
    id char(8) default (SUBSTRING( REPLACE(uuid_generate_v4()::text, '-', ''), 1, 8)),
    name VARCHAR(255) not null,
    role VARCHAR(255) default 'customer',
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    location INT NOT NULL,
    created_at TIMESTAMP default current_timestamp,
    CONSTRAINT pk_users PRIMARY KEY(id),
    CONSTRAINT fk_users_locations FOREIGN KEY(location) REFERENCES locations(id) ON DELETE CASCADE
);

INSERT INTO users (name, role, email, password, location) VALUES
('user1', 'customer', 'user1@mail.com', '$2a$10$3jnVvDAaPq/o6l6JzqBYnu6//c1/fYtDulbGyIF14SgM0gQDiL9iO', 1),
('user2', 'employee', 'user2@mail.com', '$2a$10$jYfN0ZJ5u3nWq5mcyGgKZu.NbwhnmIW/.gxBsjc4Nz0Dl6Jmlkflu', 1);
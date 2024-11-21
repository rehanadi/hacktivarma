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
('user1', 'employee', 'user1@mail.com', 'user1pass', 1),
('user2', 'employee', 'user2@mail.com', 'user2pass', 1);
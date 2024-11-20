CREATE TABLE users (
    id char(8) default (SUBSTRING( REPLACE(UUID(), '-', ''), 1, 8)),
    name VARCHAR(255) not null,
    role VARCHAR(255) default 'customer',
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    created_at TIMESTAMP default current_timestamp,
    CONSTRAINT pk_users primary key(id)
);

INSERT INTO users (name, role, email, password) VALUES
('user1', 'employee', 'user1@mail.com', 'user1pass'),
('user2', 'employee', 'user2@mail.com', 'user2pass'),
('Jono', 'customer', 'jono@mail.com', 'jonopass'),
('Miranda', 'customer', 'miranda@mail.com', 'mirandapass'),
('Lestari', 'customer', 'lestari@mail.com', 'lestaripass');
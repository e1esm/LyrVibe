CREATE TABLE users(
    id uuid primary key,
    username varchar(255) unique,
    password text,
    role varchar(10),
    profile_picture bytea
);


CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_role ON users(role);
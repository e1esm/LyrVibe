CREATE TABLE users(
    id uuid primary key,
    username varchar(255) unique,
    password text,
    role varchar(10),
    country varchar(100),
    first_name varchar(15),
    second_name varchar(30),
    profile_picture bytea
);


CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_country ON users(country);
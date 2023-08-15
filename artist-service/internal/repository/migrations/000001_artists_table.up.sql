CREATE TABLE artists(
    id UUID PRIMARY KEY,
    username varchar(10) unique,
    country varchar(40),
    first_name varchar(10),
    second_name varchar(20),
    overall_views int
);
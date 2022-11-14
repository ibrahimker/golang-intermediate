BEGIN;

CREATE TABLE users (
    id serial not null,
    username varchar(50) unique not null,
    first_name varchar(200) not null,
    last_name varchar(200) not null,
    password varchar(120) not null
);

COMMIT;
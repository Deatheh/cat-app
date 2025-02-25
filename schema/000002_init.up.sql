CREATE TABLE users
(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE cat_lists
(
    id serial PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id serial PRIMARY KEY,
    user_id int references users (id) on delete cascade      not null,
    list_id int references cat_lists (id) on delete cascade not null
);

CREATE TABLE cats
(
    id serial PRIMARY KEY,
    age int,
    name varchar(255) not null,
    filename varchar(255) not null,
    description varchar(255)
);

CREATE TABLE cats_lists
(
    id serial PRIMARY KEY,
    cat_id int references cats (id) on delete cascade      not null,
    list_id int references cat_lists (id) on delete cascade not null
);

CREATE TABLE Users
(
  id serial not null unique,
  name varchar(255) not null,
  nick_name varchar(255) not null unique,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);
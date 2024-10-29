CREATE DATABASE IF NOT EXISTS stardescription;
USE stardescription;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id varchar(36) default(uuid()) primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    criado_em timestamp default current_timestamp
) ENGINE=INNODB
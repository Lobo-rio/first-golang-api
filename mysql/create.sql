CREATE DATABASE IF NOT EXISTS first_db;
USE first_db;

DROP TABLE IF EXISTS notes;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick_name varchar(30) not null unique,
    email varchar(100) not null unique,
    phone varchar(15) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE notes(
    id int auto_increment primary key,
    title varchar(50) not null,
    description varchar(300) not null

    author_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

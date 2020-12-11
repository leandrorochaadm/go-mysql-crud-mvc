CREATE DATABASE IF NOT EXISTS usuarios;
USE usuarios;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    email varchar(30) not null unique,
    senha varchar(64) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;
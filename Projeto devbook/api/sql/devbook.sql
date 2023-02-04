create database if not exists devbook;
use devbook;
drop table if exists usuarios;
create table usuarios (
    id       int auto_increment primary key,
	nome     varchar(100) not null   ,
	nick     varchar(100) not null  unique ,
	email    varchar(100) not null  unique ,
	Senha    varchar(120)  not null ,
	CriadoEm DateTime default current_timestamp
) ENGINE=INNODB;

create user 'golang_devbook'@'localhost' identified by 'devbook_golang';
grant  all privileges on devbook.* to 'golang_devbook'@'localhost';

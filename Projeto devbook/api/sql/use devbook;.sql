use devbook;
drop table if exists usuarios;
create table usuarios (
    id       int auto_increment primary key,
	nome     varchar(100) not null   ,
	nick     varchar(100) not null  unique ,
	email    varchar(100) not null  unique ,
	Senha    varchar(20)  not null  unique,
	CriadoEm DateTime default current_timestamp
) ENGINE=INNODB;
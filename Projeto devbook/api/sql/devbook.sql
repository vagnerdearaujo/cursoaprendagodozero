create database if not exists devbook;
use devbook;
drop table if exists publicacoes;
drop table if exists seguidores;
drop table if exists usuarios;
create table usuarios (
    id       int auto_increment primary key,
	nome     varchar(100) not null   ,
	nick     varchar(100) not null  unique ,
	email    varchar(100) not null  unique ,
	Senha    varchar(120)  not null ,
	CriadoEm DateTime default current_timestamp
) ENGINE=INNODB;

--A sintaxe adotada para a tabela seguidores, segue o padrão ANSI para tornar mais fácil o uso com outros bancos
create table seguidores (
    usuario_id int not null,
    seguidor_id int not null
) ENGINE=INNODB;

alter table seguidores add constraint pk_seguidores primary key (usuario_id, seguidor_id);
alter table seguidores add constraint fk_seguidores_usuario foreign key (usuario_id) references usuarios (id) ON delete cascade;
alter table seguidores add constraint fk_seguidores_seguidor foreign key (seguidor_id)  references usuarios (id) ON delete cascade;

create table publicacoes (
	id int auto_increment primary key,
	titulo varchar(100),
	conteudo varchar(500),
	autorID int not null,
	curtidas  int default 0,
	criadaEm  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;
alter table publicacoes add constraint fk_publicacoes_usuario foreign key (autorID) references usuarios(id) on delete cascade;

create user 'golang_devbook'@'localhost' identified by 'devbook_golang';
grant  all privileges on devbook.* to 'golang_devbook'@'localhost';

--Senha usuarios: usuario1


use devbook;
insert into usuarios (nome, nick, email, senha)
values ("usuario-1","usuario1","usuario1@gmail.com","$2a$10$sgbXY9XCjyyKIaZYjabn4.l2pgjiQSNj1MzjAJwQz.z68MnJDDT4y"),
        ("usuario-2","usuario2","usuario2@gmail.com","$2a$10$sgbXY9XCjyyKIaZYjabn4.l2pgjiQSNj1MzjAJwQz.z68MnJDDT4y"),
        ("usuario-3","usuario3","usuario3@gmail.com","$2a$10$sgbXY9XCjyyKIaZYjabn4.l2pgjiQSNj1MzjAJwQz.z68MnJDDT4y"),
        ("usuario-4","usuario4","usuario4@gmail.com","$2a$10$sgbXY9XCjyyKIaZYjabn4.l2pgjiQSNj1MzjAJwQz.z68MnJDDT4y");

insert into seguidores(usuario_id, seguidor_id) 
values (1,2),
        (1,4),
        (2,1),
        (2,4),
        (3,1),
        (4,1),
        (4,2),
        (4,3);

INSERT INTO publicacoes(titulo,conteudo,`autorID`,curtidas)
VALUES('Publicação 01 - Usuário 01','Publicação do usuário 01 criada via script de inicialização do banco',1,10),
		('Publicação 02 - Usuário 01','Publicação do usuário 01 criada via script de inicialização do banco',1,10),
		('Publicação 03 - Usuário 01','Publicação do usuário 01 criada via script de inicialização do banco',1,10),
		('Publicação 04 - Usuário 01','Publicação do usuário 01 criada via script de inicialização do banco',1,10),
		('Publicação 01 - Usuário 02','Publicação do usuário 02 criada via script de inicialização do banco',2,10),
		('Publicação 02 - Usuário 02','Publicação do usuário 02 criada via script de inicialização do banco',2,10),
		('Publicação 03 - Usuário 02','Publicação do usuário 02 criada via script de inicialização do banco',2,10),
		('Publicação 04 - Usuário 02','Publicação do usuário 02 criada via script de inicialização do banco',2,10),
		('Publicação 01 - Usuário 03','Publicação do usuário 03 criada via script de inicialização do banco',3,10),
		('Publicação 02 - Usuário 03','Publicação do usuário 03 criada via script de inicialização do banco',3,10),
		('Publicação 03 - Usuário 03','Publicação do usuário 03 criada via script de inicialização do banco',3,10),
		('Publicação 04 - Usuário 03','Publicação do usuário 03 criada via script de inicialização do banco',3,10);


select * from usuarios;
select * from seguidores;

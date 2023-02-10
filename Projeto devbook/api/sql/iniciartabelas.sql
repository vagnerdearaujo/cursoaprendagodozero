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


select * from usuarios;
select * from seguidores;
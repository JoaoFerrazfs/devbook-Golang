insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1","usuario1@gmail.com","$2a$10$YL73ZWR0CvAtoHxqJPneLu7hnS3m0KEEGLLtX6iB77PEBpEujavp2"),
("Usuario 2", "usuario_2","usuario2@gmail.com","$2a$10$YL73ZWR0CvAtoHxqJPneLu7hnS3m0KEEGLLtX6iB77PEBpEujavp2"),
("Usuario 3", "usuario_3","usuario3@gmail.com","$2a$10$YL73ZWR0CvAtoHxqJPneLu7hnS3m0KEEGLLtX6iB77PEBpEujavp2"),
("Usuario 4", "usuario_4","usuario4@gmail.com","$2a$10$YL73ZWR0CvAtoHxqJPneLu7hnS3m0KEEGLLtX6iB77PEBpEujavp2");

insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(3,1),
(4,2),
(2,1);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
values
("Publicao do Usuário 1" , "Conteudo do 1", 1),
("Publicao do Usuário 2" , "Conteudo do 2", 2),
("Publicao do Usuário 3" , "Conteudo do 3", 3)
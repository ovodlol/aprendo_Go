DROP TABLE IF EXISTS empregrados;
CREATE TABLE empregrados (
  id int auto_increment not null,
  nome VARCHAR(30) not null,
  cargo varchar(100) not null,
  ganha int not null,
  primary key (`id`)
);

insert into empregrados
  (nome, cargo, ganha)
values
  ('Marcos', 'Programador', 5000),
  ('Leandro', 'Programardor junior', 3000),
  ('Julia', 'Gerente', 9000);

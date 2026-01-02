SELECT nome, preco FROM PRODUTOS;

SELECT nome FROM FUNCIONARIOS WHERE salario > 3000;

SELECT * FROM CLIENTES WHERE cidade = 'Salvador';

SELECT nickname, pontuacao FROM JOGADORES order by pontuacao DESC;


-- --------------------- INNER JOIN -----------------------------

SELECT FUNCIONARIOS.nome, EMPRESAS.nome
FROM FUNCIONARIOS
INNER JOIN EMPRESAS ON FUNCIONARIOS.empresa_id = EMPRESAS.id;

SELECT f.nome, e.nome
FROM funcionarios AS f
INNER JOIN empresas AS e
ON f.empresa_id = e.id;



-- --------------------- LEFT JOIN -----------------------------
SELECT c.nome, p.valor
FROM clientes c
LEFT JOIN pedidos p ON c.id = p.cliente_id;


------------------------- FASE 4 (RELACIONAMENTOS) -------------
CREATE TABLE PEDIDOS (
    id int primary key,
    valor decimal not null,
    cliente_id int not null,

    FOREIGN KEY(cliente_id) REFERENCES clientes(id)
);


SELECT f.nome, d.nome
FROM funcionarios AS f
INNER JOIN departamentos AS d
ON f.departamento_id = d.id;


-- --------------------- LEFT JOIN -----------------------------
SELECT d.nome
FROM departamentos AS d
LEFT JOIN funcionarios AS f ON d.id = f.departamento_id
WHERE f.id IS NULL;


SELECT a.nome
FROM AUTORES AS a
LEFT JOIN LIVROS AS l ON a.id = l.autor_id
WHERE l.id IS NULL;






























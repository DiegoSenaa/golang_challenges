-- Criação do banco de dados
CREATE DATABASE loja;

-- Conexão ao banco de dados recém-criado
-- (Você pode usar o cliente SQL apropriado para se conectar ao banco de dados)

-- Define o caminho de pesquisa para o esquema 'public'
SET search_path = public;

-- Criação da tabela de produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao VARCHAR(255),
    preco DECIMAL(10,2) NOT NULL,
    qtd INT NOT NULL
);

-- Inserção de dados na tabela (10 produtos)
INSERT INTO produtos (nome, descricao, preco, qtd) VALUES 
    ('Camisa', 'Preta', 19.99, 10),
    ('Shorts', 'Azul', 14.50, 5),
    ('Tênis', 'Branco', 59.99, 8),
    ('Boné', 'Vermelho', 12.99, 20),
    ('Mochila', 'Preto', 34.99, 15),
    ('Luvas', 'Fitness', 9.99, 30),
    ('Garrafa', 'Esportiva', 7.50, 25),
    ('Corda', 'Pular', 5.99, 40),
    ('Squeeze', 'Azul', 3.99, 50),
    ('Meias', 'Esportivas', 6.50, 12);

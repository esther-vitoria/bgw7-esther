## Consultas SQL Avançadas ##

### Bancos de Dados Relacionais ###
Pratica na Aula - Bancos de dados relacionais

### Resolva as seguintes perguntas ###
Você está solicitando o banco de dados movies_db.sql:

1. Adicionar um filme à tabela de movies.
````
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
VALUES ('Como eu era antes de voce', 7.4, 4, '2018-06-08', 110, NULL);
````

2. Adicione um gênero à tabela de genres.
```
INSERT INTO genres (name, ranking, active)
VALUES ('Romance', 13, 1);
```

3. Associe o gênero criado no ponto 2 ao filme no ponto 1. gênero.
```
UPDATE movies
SET genre_id = 13
WHERE title = 'Como eu era antes de voce';
```
4. Modifique a tabela de atores para que pelo menos um ator tenha como favorito o filme adicionado no ponto 1.
```
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 32;
```
5. Crie uma cópia temporária da tabela de movies.
```
CREATE TEMPORARY TABLE temp_movies AS SELECT * FROM movies;
```
6. Remova dessa tabela temporária todos os filmes que ganharam menos de 5 awards.
````
DELETE FROM temp_movies WHERE awards < 5;
````
7. Obtenha a lista de todos os gêneros que têm pelo menos um movies.
````
SELECT DISTINCT g.*
FROM genres g
JOIN movies m ON m.genre_id = g.id;
````
8. Obtenha a lista de atores cujo filme favorito ganhou mais de 3 awards.
````
SELECT a.*
FROM actors a
JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;
````
9. Crie um índice sobre o nome na tabela de movies.
````
CREATE INDEX idx_movies_title ON movies(title);
````
10. Verifique se o índice foi criado corretamente.
`````
SHOW INDEX FROM movies;
`````
11. No banco de dados de movies, você notaria uma melhora significativa com a criação de índices? Analise e justifique sua resposta.        
**R:** Indices melhoram muito a performance das consultas em tabelas grandes, acredito que o ideal é ter indices nos campos **title**, **genre_id**, **favorite_movie_id** e campos de relacionamento aceleram bastante o acesso.

12. Em qual outra tabela você criaria um índice e por quê? Justifique sua resposta.     
**R:** Criaria um índice na tabela **episodes** no campo **season_id**, porque as consultas frequentemente buscam episódios de uma temporada específica ou realizam JOINs entre episodes e seasons. Isso melhora muito a performance da consulta.
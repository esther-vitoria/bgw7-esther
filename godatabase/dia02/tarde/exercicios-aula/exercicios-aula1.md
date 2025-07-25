### Criação de tabelas de tempo e índices ###

#### // Prática na Aula ####

#### Exercício 1 ####

1. Com a base de dados "movies", propõe-se criar uma tabela temporária chamada "TWD" e armazenar nela os episódios de todas as temporadas de "The Walking Dead".
````
CREATE TEMPORARY TABLE TWD AS
SELECT e.*
FROM episodes e
JOIN seasons s ON e.season_id = s.id
JOIN series sr ON s.serie_id = sr.id
WHERE sr.title = 'The Walking Dead';

````
2. Executa uma consulta à tabela de tempo para ver os episódios da primeira temporada.
````
SELECT * 
FROM TWD
WHERE season_id = (
  SELECT id FROM seasons WHERE serie_id = (SELECT id FROM series WHERE title = 'The Walking Dead') AND number = 1
);
````

#### Exercício 2 ####

1. No banco de dados "movies", selecione uma tabela para criar um índice e, em seguida, verifique a criação do índice.
````
CREATE INDEX idx_episodes_season_id ON episodes(season_id);
SHOW INDEX FROM episodes;
````


2. Analise por que você criaria um índice na tabela indicada e com quais critérios você escolheria o(s) campo(s).   
**R:** Acredito que ficaria melhor e mais rápida a pesquisa se houver um filtro de episódios por temporada, para isso criei um index na tabela **episodes**, no campo **seasons_id**, para os critérios foi levado em consideração as relações (JOINs) mais frequentes e o valume de dados.
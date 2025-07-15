## Bancos de dados relacionais ##
**Consultas SQL Avançadas**

// Prática Individual

Primeira Parte  
Responda às seguintes perguntas:  
O que é um **JOIN** em um banco de dados e para que ele é usado?

1. É usado para obter dados de várias tabelas relacionadas. Consiste na combinação de dados de uma tabela com dados de outra tabela, com base em uma ou mais condições em comum.  
     
Explicar dois tipos de **JOIN**.  
1. Inner Join é usado para obter dados relacionados de duas ou mais tabelas.  
2. Left Join é usado para obter dados da tabela da esquerda com dados relacionados da tabela da direita.  
     
Para que serve o **GROUP BY**?  
1. Agrupa os resultados de acordo com as colunas especificadas.  
2. Gera um único registro para cada grupo de linhas que compartilham as colunas especificadas.  
3. Reduz o número de linhas na consulta.  
4. Geralmente usado em conjunto com funções de agregação, para obter dados resumidos e agrupados pelas colunas necessárias.  
     
Para que é usado o **HAVING**?  
1. A cláusula **HAVING** é usada para incluir condições em algumas funções **SQL**.  
2. Isso afeta os resultados obtidos pelo Group By.  
3. Escreva uma consulta genérica para cada um dos diagramas a seguir:  
![Diagrama](/godatabase/dia02/manha/exercicios-pratica/image/diagrama.png)

SELECT movies.\*, actors.first\_name, actors.last\_name  
FROM movies INNER JOIN actors  
ON movies.id \= actors.favorite\_movie\_id;  
SELECT \* FROM movies mo LEFT JOIN actors ac ON mo.id \= ac.favorite\_movie\_id;

Segunda Parte  
Propõe-se realizar as seguintes consultas ao banco de dados **movies\_db.sql** trabalhado na primeira aula. Importe o arquivo **movies\_db.sql** do ***PHPMyAdmin*** ou do ***MySQL Workbench*** e resolva as seguintes consultas:

1. Exibir o título e o nome do gênero de todas as séries. 

![Exercício 1](/godatabase/dia02/manha/exercicios-pratica/image/exc1.png)

2. Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.  

![Exercício 2](/godatabase/dia02/manha/exercicios-pratica/image/exc2.png)

3. Mostre o título de todas as séries e o número total de temporadas de cada série.  

![Exercício 3](/godatabase/dia02/manha/exercicios-pratica/image/exc3.png)

4. Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3\.  

![Exercício 4](/godatabase/dia02/manha/exercicios-pratica/image/exc4.png)

5. Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.

![Exercício 5](/godatabase/dia02/manha/exercicios-pratica/image/exc5.png)

Como guia para as consultas, lembre-se do DER do cenário:  
![DER](/godatabase/dia02/manha/exercicios-pratica/image/der.png)

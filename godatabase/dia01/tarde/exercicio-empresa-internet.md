  
Bancos de dados relacional  

**DER e SQL**  

**//** Prática de grupo integrativa  

**Objetivo**  
O objetivo deste guia prático é poder integrar os conteúdos de bancos de dados relacionais vistos até agora. Para isso, propõe-se a seguinte prática.  
**Boa sorte\!**  

**Cenário**  
Uma empresa provedora de Internet precisa de um banco de dados para armazenar cada um de seus clientes, juntamente com o plan/pack contratado.  
Por meio de uma análise prévia, sabe-se que as seguintes informações devem ser armazenadas:

* Para clientes: número de identificação, nome, sobrenome, data de nascimento, província, cidade.  
* Para planos de Internet: identificação do plano, velocidade oferecida em megabytes, preço, desconto.

 **Exercício 1**  
Após a definição dos requisitos da empresa, é solicitado que eles sejam modelados por meio de um DER (Diagrama Entidade-Relacionamento). 

Resposta:
![DER exercício 1](/godatabase/dia01/tarde/image/DER-empresa-internet.png)


**Exercício 2**   
Depois que o banco de dados tiver sido modelado e apresentado, responda às seguintes perguntas:  

**a.** Qual é a **primary key** da tabela de clientes? Justifique sua resposta  
**resposta:** A PK da tabela clientes é o **customer_id** ela representa e identifica os clientes.

**b.** Qual é a **primary key** da tabela de planos de Internet? Justifique a resposta.  
**resposta:** A PK da tabela planos é o **service_id** ela representa e identifica cada serviço oferecido pela empresa.

**c.** Como seriam os relacionamentos entre as tabelas? Em qual tabela deve haver uma **foreign key**? A qual campo de qual tabela essa **foreign key** se refere? Justifique a resposta.  
**resposta:** Para esse caso foi colocado uma FK na tabela **customers** identificada como **service_id** que se refere a tabela **services**. Tipo de relação Um-para-muitos (1 plano → N clientes).

**Exercício 3**  
Depois que o diagrama tiver sido criado e essas perguntas tiverem sido respondidas, use o **PHPMyAdmin** ou o **MySQL Workbench** para executar o seguinte:

* É solicitado que você crie um novo banco de dados chamado **"empresa\_internet"**.  

![exercício 3.1](/godatabase/dia01/tarde/image/create-database.png)

* Incorpore 10 registros na tabela de clientes e 5 na tabela de planos de Internet.  

Tabela Clientes.      
![exercício 3.2](/godatabase/dia01/tarde/image/insert-customers.png)

Tabela Serviços/Planos.     
![exercício 3.2](/godatabase/dia01/tarde/image/insert-services.png)

* Faça as associações/relacionamentos correspondentes entre esses registros.

![exercício 3.2](/godatabase/dia01/tarde/image/create-tables-fk.png)

**Exercício 4**  
Indique 10 consultas SQL que poderiam ser feitas no banco de dados. Expresse as instruções.  

1. Listar todos os clientes
    ```
    SELECT * FROM customers;
    ````
2. Listar todos os planos de internet
    ````
    SELECT * FROM services;
    ````
3. Mostrar o nome dos clientes e a velocidade do plano contratado
    `````
    SELECT c.first_name, c.last_name, s.speed_mb
    FROM customers c
    JOIN services s ON c.service_id = s.service_id;
    `````
4. Clientes que moram em Manaus
    ````
    SELECT * FROM customers WHERE city = 'Manaus';
    ````
5. Nome, cidade e desconto dos clientes que têm planos com desconto maior que 20
    ````
    SELECT c.first_name, c.city, s.discount
    FROM customers c
    JOIN services s ON c.service_id = s.service_id
    WHERE s.discount > 20;
    ````
6. Planos com velocidade maior que 300Mb
    ````
    SELECT * FROM services WHERE speed_mb > 300;
    ````
7. Contar o número de clientes por plano
    ````
    SELECT s.service_id, s.speed_mb, COUNT(*) AS num_customers
    FROM services s
    JOIN customers c ON s.service_id = c.service_id
    GROUP BY s.service_id, s.speed_mb;
    ````
8. Clientes nascidos após o ano 1990
    ````
    SELECT first_name, last_name, birth_date FROM customers WHERE birth_date > '1990-01-01';
    ````

9. Mostrar todos os clientes ordenados pela data de nascimento
    ````
    SELECT * FROM customers ORDER BY birth_date;
    ````
10. Listar clientes e o valor final do plano considerando o desconto
    ````
    SELECT c.first_name, c.last_name, 
       (s.price - s.discount) AS final_price 
    FROM customers c
    JOIN services s ON c.service_id = s.service_id;
    ````

## Prática 1 ##

### Objetivo ###

O objetivo deste guia prático é que possamos consolidar e aprofundar os conceitos sobre a implementação de bancos de dados. Para isso, apresentaremos uma série de exercícios que nos permitirão revisar os tópicos que estamos estudando. Vamos começar com o trabalho realizado no módulo Go Web.

Nesta prática, vamos **implementar um CRUD no GO**.

### Problema ###

Depois de entregar a solução ao nosso cliente, um supermercado de produtos frescos, surgiram novos requisitos. Agora precisamos implementar uma persistência mais robusta do registro de dados. Para isso, vamos implementar um armazenamento em um **banco de dados SQL** sem precisar excluir o armazenamento em um arquivo json.

## Exercício 1: configuração de um banco de dados ##

Para implementar a alteração, precisamos criar um banco de dados com os produtos registrados. Para isso, temos um **script** que compartilhamos abaixo:


[build_database.sql](https://drive.google.com/file/d/1qCBE27m4Vme2Akk4oBvZo3tLPZPmAvpP/view)

Isso cria um novo usuário, gera um esquema com o nome **my_db** e uma série de produtos pré-carregados.

Para carregar esse script, podemos usar o seguinte comando em um terminal onde o script está localizado:

````
sudo mysql -u root -p < build_database.sql
````

## Exercício 2: CRUD do produto ##

Com base no que fizemos no Go Web, vamos implementar as operações **Create, Read, Update e Delete**. Da mesma forma que resolvemos com um arquivo JSON.

*Com base na interface do repositório, crie uma implementação que interaja com um banco de dados.

Então, devemos implementar:

- **CREATE**: Criar um novo produto.
- **LEIA**: Trazendo um novo produto.
- **UPDATE**: Atualizar um produto.
- **DELETE**: Excluir um produto.
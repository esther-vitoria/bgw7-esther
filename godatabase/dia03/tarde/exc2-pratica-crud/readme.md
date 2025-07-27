## Prática 2 ##

### Objetivo ###

O objetivo deste guia prático é que possamos consolidar e aprofundar os conceitos sobre a implementação de bancos de dados. Para isso, apresentaremos uma série de exercícios que nos permitirão revisar os tópicos que estudamos. Vamos nos basear no trabalho realizado na aula anterior.

Nesta prática, vamos implementar um método **GetAll** e realizar um **Join** entre tabelas.

### Problema ###

Depois de entregar a solução ao nosso cliente, um supermercado de produtos frescos, surgiram novos requisitos. Decidiu-se abrir novas filiais, na forma de depósitos em diferentes locais onde os novos produtos seriam armazenados.

## Exercício 1: Joins ##

Vamos adicionar um novo domínio ao nosso aplicativo, **warehouses**. Essa é uma representação dos depósitos onde os produtos são armazenados. *Cada produto é armazenado em um depósito e um depósito pode ter muitos produtos*. O domínio terá os seguintes campos:

````
{
   "id": 1,
   "name": "SuperMarket",
   "address": "123 Main Street",
   "telephone": "555-555-5555",
   "capacity": 500
}
````

Para implementar esse domínio, precisamos atualizar nosso aplicativo. Para isso, adicionaremos o campo **werehouse_id** à estrutura *do product*, que será a *chave estrangeira* que indica a qual depósito o produto pertence. Quando um novo produto for criado, o ID do depósito ao qual ele pertence deverá ser enviado; esse campo será **obrigatório**.


Compartilhamos com você um script para criar um novo domínio no banco de dados, que é atribuído a cada produto já criado no depósito com o ID 1:

[adicionar warehouses.sql](https://drive.google.com/file/d/16HbiGxCsLsnXsYWXe2cxZRla3TVoYhfw/view)

Depois de definir o campo, vamos criar novos pontos de extremidade com o domínio, que são

- **GET warehouses/{id} ->** *Read One* | Buscar um depósito por id
- **POST warehouses -> Create** | Criar um depósito

Além disso, dentro do produto, adicionaremos uma nova consulta:

- **GET warehouse/reportProducts?id=[int] ->** *Join* | Obtenha um relatório do número de produtos nos depósitos (inclusive se eles tiverem 0). Caso o ID do depósito chegue por consulta, aplique o filtro correspondente.

````
[
    {

        "name": "SuperMarket",
        "product_count": 100
    }
]
````

## Exercício 2: Read All ##

O objetivo deste exercício é implementar um método **Read All** dos **products** e **warehouses** no banco de dados.

- **GET warehouses ->** *Read All* | Fetch list of **warehouses** (Obter lista de **warehouses**)
- **GET products ->** *Read All* | Obter uma lista de **products**



## Exercício 3: Integration Test ## 
Para validar se os novos requisitos são funcionais para as próximas iterações do aplicativo, vamos testar as novas funcionalidades. Vamos fazer um teste de unidade do repositório. Para isso, usaremos o pacote [go-txdb](https://pkg.go.dev/github.com/DATA-DOG/go-txdb).

| Método                        | Resposta esperada                  | Descrição                                                          |
|-------------------------------|------------------------------------|--------------------------------------------------------------------|
| **Read All \| warehouses**    | *Lista de todos os warehouses*     | Espera-se que todos os *warehouses* registrados sejam obtidos.     |
| **Read All \| Products**      | *Lista de todos os products*       | Espera-se que todos os *products* registrados sejam obtidos.       |
| **Read One \| warehouses**    | *Warehouse esperado por id*        | Obter o warehouse com o ID solicitado.                             |
| **Create \| warehouses**      | *Warehouse criado*                 | Um warehouse é adicionado e devolvido.                             |


## Exercício 3.1: Expansão de tests [ Opcional ]## 
Para melhorar a capacidade de manutenção do aplicativo, convidamos você a implementar testes de unidade dos outros métodos de repositório **de products (Store, Delete, Update)**.
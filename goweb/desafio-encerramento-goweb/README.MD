## Desafio Encerramento Go Web
**Objetivo** 

Este projeto tem como objetivo consolidar conhecimentos no desenvolvimento de APIs RESTful do modulo GoWeb ministrado no Bootcamp Go, seguindo boas práticas de arquitetura e Domain Oriented Design (DOD). O sistema simula operações de uma companhia aérea, permitindo o cálculo e extração de dados estatísticos sobre voos e passageiros, a partir de uma API.

## Estrutura do Projeto

````
desafio-encerramento-goweb/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── docs/ 
│   └── tickets.csv    #CSV com as informações dos tickets                                    
├── internal/
│   ├── application/
│   │   └── application.go # Configuração das rotas e servidor
│   ├── handler/
│   │   └── ticket_handler.go      # Controllers/handlers da API
│   ├── loader/
│   │   └── ticket_loadercsv.go    # Leitura de tickets do CSV
│   ├── repository/
│   │   ├── ticket_map.go
│   │   ├── ticket_mock.go
│   └── service/
│       ├── ticket_service_test.go      # Testes via camada de serviço  
│       └── ticket_service.go      # Regras de negócio
├── ticket_json.go   # Persistência/acesso aos dados
├── ticket_repository.go   # Persistência/acesso aos dados
├── ticket_service.go   # Persistência/acesso aos dados
├── pkg/apperrors/
│   └── errors.go                  # Erros customizados da aplicação
├── utils/
│   └── json_respond.go            # Função para resposta padronizada JSON
├── go.mod
├── go.sum
├── Makefile
└── README.md
````

## Endpoints Disponíveis

| Método | Endpoint | Descrição |
| -------- | -------- | -------- |
| GET  | ``/tickets/get``  | Lista todos os tickets cadastrados  |
| GET  | ``/tickets/getByCountry/{dest}``  | Busca todos os tickets para um país de destino específico  |
| GET  | ``/tickets/getPercent/{dest}``  | Calcula a porcentagem de tickets para um país específico  |

**Exemplos de Uso** 

Listar todos os tickets:
````
curl http://localhost:8080/tickets/get
````
**Resposta** 
````
{
    "message": "Success!",
    "data": {
        "1": {
            "name": "Tait Mc Caughan",
            "email": "tmc0@scribd.com",
            "country": "Finland",
            "hour": "17:11",
            "price": 785
        },
        "10": {
            "name": "Fax Cordaroy",
            "email": "fcordaroy9@i2i.jp",
            "country": "Russia",
            "hour": "1:44",
            "price": 1163
        }
    }
}        
````

Buscar tickets por país:
````
curl http://localhost:8080/tickets/getByCountry/Finland
````
**Resposta** 
````
{
    "message": "Success!",
    "data": {
        "1": {
            "name": "Tait Mc Caughan",
            "email": "tmc0@scribd.com",
            "country": "Finland",
            "hour": "17:11",
            "price": 785
        }
    }
}        
````

Buscar percentual de tickets para um país:
````
curl http://localhost:8080/tickets/getPercent/Finland
````
**Resposta** 
````
{
    "message": "The tickets percentage for this country is:",
    "data": 0.008,
    "Error": false
}        
````
## Como Executar o Projeto

**Pré-requisitos** 

- Go instalado

- air (opcional, para hot reload em desenvolvimento)

**Instale as dependências** 
````
go mod tidy
````

**Execute (modo padrão/produção)** 
````
make start
````

**Para desenvolvimento (hot reload)** 

Instale o air e rode:
````
make dev
````


## Testes

Testes de unidade foram implementados com Testify.

Para executar os testes:
````
go test ./...
````
Os testes cobrem, por exemplo, o cálculo correto do total de tickets via camada de serviço, utilizando mocks para simular os dados da camada de repositório.

## Tecnologias

- Go 1.23.5: linguagem de programação principal.  
- go-chi/chi: roteamento HTTP.  
- stretchr/testify: biblioteca de testes unitários.  
- air (opcional): hot reload no desenvolvimento.  


## Observações

O sistema foi estruturado em camadas (handlers, services, repositories) e utiliza um loader para importar os tickets de um arquivo CSV.      
Boas práticas DOD e arquitetura REST separam claramente responsabilidades e facilitam futuras manutenções.    
Issues, sugestões ou pull requests são bem-vindos!
# Desafio Go Bases

Este projeto foi desenvolvido durante o Bootcamp Go, tem como objetivo manipular, analisar e exibir informações referentes a tickets de viagens, utilizando dados presentes em um arquivo CSV, colocando em prática o que foi aprendido no módulo de Go Bases.

## Funcionalidades

- **Consulta de Tickets por Destino:**  
  Mostra o total de tickets para um destino informado pelo usuário.

- **Porcentagem de Tickets por Destino:**  
  Informa o percentual de tickets referentes a um destino em relação ao total.

- **Contagem de Tickets por Período do Dia:**  
  Exibe quantos tickets correspondem a um período específico (madrugada, manhã, tarde, noite).

## Estrutura de Diretórios

```
.
├── internal
│   ├── db
│   │   └── tickets.csv      # Base de dados de tickets
│   └── tickets
│       ├── tickets.go       # Implementação das regras de negócio
│       └── tickets_test.go  # Testes automatizados
├── main.go                  # Ponto de entrada do projeto
├── go.mod
└── README.md
```

## Como Executar

1. **Pré-requisitos:**
   - Go instalado (versão mínima recomendada: 1.17)
   - Arquivo de dados `tickets.csv` presente em `internal/db/`.

2. **Rodando o programa:**

   ```bash
   go run cmd/main.go

**O programa irá perguntar:**  
O destino desejado (ex: Brazil)  
O período do dia (madrugada, manhã, tarde, noite)

**Exemplo de interação:** 
```
Insira um destino: Brazil
Total de tickets para o destino Brazil: 45
Porcentagem para o destino Brazil: 4.50%
Insira um periodo (madrugada, manhã, tarde, noite): tarde
Quantidade no periodo manhã: 289  
```

## Principais Funções
**LoadTickets() ([]Ticket, error):**  
Carrega todos os tickets do arquivo CSV.

**GetTotalTickets(destination string) (int, error):**   
Retorna o total de tickets para um destino.

**GetCountByPeriod(period string) (int, error):**  
Conta os tickets no período (madrugada, manhã, tarde, noite).

**PercentageDestination(destination string, total int) (float64, error):**  
Calcula a porcentagem de tickets para um destino em relação ao total.

## Testes

Para garantir o correto funcionamento, execute os testes:


1. **Rodando o programa:**

   ```bash
   go run cmd/main.go

**Os testes cobrem:**  
- Retorno correto do total de tickets por destino
- Contagem por período
- Porcentagem de destinos



**Autor:**
Esther Vitória


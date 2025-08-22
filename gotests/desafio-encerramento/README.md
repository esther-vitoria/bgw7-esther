# API de Produtos - Desafio de Encerramento

Uma API REST em Go para gerenciamento de produtos, desenvolvida com arquitetura limpa e cobertura de testes abrangente.

## 🚀 Funcionalidades

- **Consulta de Produtos**: Endpoint para listar todos os produtos ou buscar por ID específico
- **Arquitetura Limpa**: Separação clara entre camadas (handler, service, repository)
- **Cobertura de Testes**: Testes unitários para todas as camadas da aplicação, coverage >80% para todo o projeto
- **Middleware**: Logging e recuperação de erros integrados

## 🏗️ Arquitetura

O projeto segue os princípios da arquitetura limpa com as seguintes camadas:

```
├── cmd/                    # Ponto de entrada da aplicação
├── internal/               # Lógica de negócio interna
│   ├── application/        # Configuração e inicialização da aplicação
│   ├── handler/           # Handlers HTTP (controllers)
│   ├── service/           # Lógica de negócio
│   └── repository/        # Acesso a dados
├── database/              # Dados dos produtos (JSON)
├── loader/                # Utilitários para carregar dados
├── plataform/             # Utilitários da plataforma web
│   └── web/
│       ├── request/       # Processamento de requisições
│       └── response/      # Formatação de respostas
└── coverage.out           # Relatório de cobertura de testes
```

## 🛠️ Tecnologias Utilizadas

- **Go 1.21**: Linguagem de programação
- **Chi Router**: Roteador HTTP leve e performático
- **Testify**: Framework de testes
- **JSON**: Armazenamento de dados em arquivo

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Make (opcional, para usar os comandos do Makefile)

## 🚀 Como Executar

### Instalação

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd desafio-encerramento

# Baixe as dependências
go mod download
```

### Executar a Aplicação

```bash
# Usando Make
make start

# Ou diretamente com Go
go run ./cmd/main.go
```

A aplicação será iniciada em `http://127.0.0.1:8080`

## 📚 API Endpoints

### GET /product

Retorna a lista de produtos ou um produto específico por ID.

**Parâmetros de Query:**
- `id` (opcional): ID do produto a ser buscado

**Exemplos de Uso:**

```bash
# Listar todos os produtos
curl http://127.0.0.1:8080/product

# Buscar produto específico por ID
curl http://127.0.0.1:8080/product?id=1
```

**Resposta de Sucesso:**
```json
{
  "message": "success",
  "data": {
    "1": {
      "id": 1,
      "description": "French Pastry - Mini Chocolate",
      "price": 97.01,
      "seller_id": 1
    }
  }
}
```

**Resposta de Erro:**
```json
{
  "status": "Bad Request",
  "message": "invalid id"
}
```

## 🧪 Executar Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com cobertura
make cover

# Gerar relatório HTML de cobertura
make cover-html

# Ver cobertura filtrada
make cover-filtered

# Ver total de cobertura filtrada
make cover-filtered-total
```

## 📊 Estrutura de Dados

### Produto

```go
type Product struct {
    Id          int     `json:"id"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    SellerId    int     `json:"seller_id"`
}
```

## 📝 Comandos Make Disponíveis

- `make start`: Inicia a aplicação
- `make cover`: Executa testes com cobertura
- `make cover-html`: Gera relatório HTML de cobertura
- `make cover-filtered`: Executa cobertura filtrada
- `make cover-filtered-total`: Mostra total de cobertura filtrada


## 👥 Autores

- Desenvolvido como parte do desafio de encerramento do curso

---

**Nota**: Este projeto é uma demonstração de boas práticas em Go, incluindo arquitetura limpa, testes unitários e documentação completa.

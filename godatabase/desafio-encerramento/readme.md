# 🚀 Fantasy Products API - Desafio de Fechamento

## 📖 Sobre o Projeto

Sistema de vendas **Fantasy Products** desenvolvido em Go com MySQL. O projeto implementa recuperação de dados perdidos através de arquivos JSON e uma API REST completa para gerenciamento de clientes, produtos, faturas e vendas.

### 🎯 Contexto do Desafio
Os dados do banco foram perdidos, mas arquivos JSON das tabelas foram recuperados. O sistema implementa:
- **Migração automática** de dados JSON para MySQL
- **Recálculo de totais** de faturas baseado em vendas
- **Relatórios analíticos** de clientes e produtos
- **API REST completa** com endpoints CRUD

![DER do Sistema](der.png)

## 🛠️ Tecnologias Utilizadas

### Backend
- **Go 1.23.5** - Linguagem principal
- **Chi Router v5.2.2** - Framework web HTTP
- **MySQL 8.0** - Banco de dados relacional
- **Docker & Docker Compose** - Containerização

### Dependências Principais
- `github.com/go-sql-driver/mysql` - Driver MySQL
- `github.com/joho/godotenv` - Variáveis de ambiente
- `github.com/stretchr/testify` - Framework de testes
- `github.com/DATA-DOG/go-txdb` - Mock de transações

### Ferramentas
- **Air** - Hot reload
- **Portainer** - Gerenciamento Docker

## ⚙️ Pré-requisitos

- **Go 1.23.5+**
- **Docker & Docker Compose**
- **Make** (recomendado)

## 🚀 Como Executar

### 1. Configuração Inicial
```bash
# Clone e navegue para o projeto
git clone <url-do-repositorio>
cd bgw7-esther/godatabase/desafio-encerramento

# Instalar dependências
go mod tidy
```

### 2. Criar arquivo .env
```env
# Banco de Dados
DB_USER=root
DB_PASSWORD=root
DB_NET=tcp
DB_ADDR=localhost:3306
DB_NAME=fantasy_products

# Servidor
HOST=localhost
PORT=8080

# Arquivos JSON
CUSTOMER_JSON=./database/json/customers.json
PRODUCT_JSON=./database/json/products.json
INVOICE_JSON=./database/json/invoices.json
SALE_JSON=./database/json/sales.json
```

### 3. Executar com Docker (Recomendado)
```bash
# Subir MySQL + Portainer
make up-build

# Executar aplicação (com hot reload)
make dev
```

### 4. Execução Manual
```bash
# Subir apenas MySQL
docker-compose up -d db

# Executar aplicação
go run ./cmd/main.go
```

## 🌐 Endpoints da API

### 👥 Customers
- `GET /customers` - Lista todos os clientes
- `POST /customers` - Cria novo cliente
- `GET /customers/top-active` - Top 5 clientes ativos por valor gasto
- `GET /customers/invoices-by-condition` - Faturas por condição (ativo/inativo)

### 🛍️ Products
- `GET /products` - Lista todos os produtos
- `POST /products` - Cria novo produto
- `GET /products/top-sold` - Top 5 produtos mais vendidos

### 🧾 Invoices
- `GET /invoices` - Lista todas as faturas
- `POST /invoices` - Cria nova fatura
- `PUT /invoices/total` - **Recalcula totais** de todas as faturas

### 💰 Sales
- `GET /sales` - Lista todas as vendas
- `POST /sales` - Registra nova venda

## 📊 Funcionalidades Principais

### 🔄 Migração de Dados
- **Loaders**: Carregam dados de arquivos JSON
- **Migrators**: Migram dados para o banco MySQL
- **Automático**: Executa na inicialização da aplicação

### 📈 Relatórios Analíticos

#### 1. Valores por Condição de Cliente
```bash
GET /customers/invoices-by-condition
```
```json
{
    "data": [
        {
            "condition": "Inactivo ( 0 )",
            "total": 570326.75
        },
        {
            "condition": "Activo ( 1 )",
            "total": 752394.69
        }
    ],
    "message": "customers found"
}
```

#### 2. Top 5 Produtos Mais Vendidos
```bash
GET /products/top-sold
```
```json
{
    "data": [
        {
            "description": "Vinegar - Raspberry",
            "total": 660
        },
        {
            "description": "Flour - Corn, Fine",
            "total": 521
        },
        {
            "description": "Cookie - Oatmeal",
            "total": 467
        },
        {
            "description": "Pepper - Red Chili",
            "total": 439
        },
        {
            "description": "Chocolate - Milk Coating",
            "total": 436
        }
    ],
    "message": "products found"
}
```

#### 3. Top 5 Clientes Ativos
```bash
GET /customers/top-active
```
```json
{
    "data": [
        {
            "first_name": "Lannie",
            "last_name": "Tortis",
            "total": 58513.55
        },
        {
            "first_name": "Jasen",
            "last_name": "Crowcum",
            "total": 48291.03
        },
        {
            "first_name": "Lazaro",
            "last_name": "Anstis",
            "total": 40792.06
        },
        {
            "first_name": "Tomasina",
            "last_name": "Kieran",
            "total": 39162.4
        },
        {
            "first_name": "Cassondra",
            "last_name": "Penbarthy",
            "total": 33749.85
        }
    ],
    "message": "customers found"
}
```

## 🧪 Testes

### Executar Testes
```bash
# Todos os testes
go test ./...

# Com cobertura
make cover
make cover-html

# Por módulo
go test ./internal/customer/...
go test ./internal/product/...
```

### 🎯 Cobertura de Testes
- ✅ **CRUD básico** - Create, Read operations
- ✅ **Consultas complexas** - JOINs e agregações
- ✅ **Funcionalidades específicas** - Recálculo de totais
- ✅ **Repositories** - 12 testes cobrindo todos os módulos
- ✅ **Mock de transações** - Usando go-txdb

## 🔧 Comandos Úteis

### Make Commands
```bash
make up          # Sobe containers
make down        # Para containers
make logs        # Visualiza logs
make cover       # Executa testes com cobertura
```

### Docker
```bash
# Logs do MySQL
docker-compose logs -f db

# Acessar MySQL
docker exec -it database mysql -uroot -proot fantasy_products

# Status dos containers
docker-compose ps
```

## 📁 Estrutura do Projeto

```
.
├── cmd/main.go                 # Ponto de entrada
├── internal/
│   ├── application/            # Configuração da aplicação
│   ├── customer/              # Módulo de clientes
│   ├── product/               # Módulo de produtos
│   ├── invoice/               # Módulo de faturas
│   ├── sale/                  # Módulo de vendas
│   ├── domain/                # Entidades de domínio
│   ├── loader/                # Carregadores JSON
│   └── migrator/              # Migradores de dados
├── handler/                   # Handlers HTTP
├── database/
│   ├── json/                  # Arquivos JSON (dados recuperados)
│   └── mysql/                 # Scripts SQL
└── pkg/                       # Pacotes compartilhados
```

## 🔍 Troubleshooting

### Problemas Comuns
```bash
# MySQL não conecta
docker-compose restart db
docker-compose logs db

# Porta ocupada
lsof -i :8080
# Altere PORT no .env

# Dependências Go
go clean -modcache
go mod tidy

# Air não funciona
go install github.com/air-verse/air@latest
```

## 🎉 Funcionalidades Implementadas

✅ **Migração completa** de dados JSON para MySQL  
✅ **API REST** com 12 endpoints  
✅ **Recálculo automático** de totais de faturas  
✅ **Relatórios analíticos** com consultas complexas  
✅ **Testes unitários** com cobertura completa  
✅ **Containerização** com Docker Compose  
✅ **Hot reload** para desenvolvimento  

---

## 🚀 Quick Start

```bash
# 1. Clone e navegue
git clone <repo> && cd bgw7-esther/godatabase/desafio-encerramento

# 2. Configure ambiente
cp .env.example .env  # Ajuste as variáveis

# 3. Execute
make up-build && make dev

# 4. Teste a API
curl http://localhost:8080/customers
```
⚙️ **Postaman Colletion:** `postman_collection.json`        
🌐 **API:** http://localhost:8080  
🐳 **Portainer:** http://localhost:9000  
📊 **Banco:** MySQL em localhost:3306
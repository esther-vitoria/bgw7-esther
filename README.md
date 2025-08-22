# 🚀 BGW7 - Bootcamp Go Wave 7 - Esther

Este repositório contém todos os projetos e exercícios desenvolvidos durante o **Bootcamp Go Wave 7 (BGW7)**, cobrindo desde fundamentos da linguagem Go até desenvolvimento de APIs REST completas, bancos de dados e testes automatizados.

## 📚 Estrutura do Repositório

### 🔰 **gobases/** - Fundamentos Go
Exercícios e conceitos básicos da linguagem Go, incluindo:
- **Dia 01**: Variáveis, tipos, estruturas condicionais, loops, arrays, slices, maps
- **Dia 02**: Funções, múltiplos retornos, tratamento de erros
- **Dia 03**: Structs, métodos, interfaces, ponteiros, embedding
- **Dia 04**: Tratamento de erros avançado, panic, recover, defer
- **Dia 05**: Manipulação de arquivos, I/O, formatação
- **🎯 Desafio**: Sistema de análise de tickets de viagem com CSV

### 🗄️ **godatabase/** - Banco de Dados
Integração com bancos de dados MySQL, incluindo:
- **Dia 01**: Modelagem de dados, consultas SQL básicas
- **Dia 02**: Consultas avançadas, JOINs, agregações
- **Dia 03**: CRUD com Go, conexões de banco
- **🎯 Desafio**: **Fantasy Products API** - Sistema completo de vendas com migração de dados JSON para MySQL

### 🧪 **gotests/** - Testes Automatizados
Estratégias e práticas de testes em Go:
- **Dia 01**: Testes unitários básicos, table-driven tests
- **Dia 02**: Mocks, stubs, testes de integração
- **Dia 03**: Cobertura de código, benchmarks, testes avançados
- **🎯 Desafio**: API de produtos com cobertura de testes >80%

### 🌐 **goweb/** - APIs REST
Desenvolvimento de APIs web com Go:
- **Dia 01**: Servidores HTTP, JSON, roteamento
- **Dia 02**: Middlewares, validação, tratamento de erros
- **🎯 Desafio**: API de companhia aérea com análise estatística de voos

## 🏆 Principais Projetos (Desafios)

### 1. 🎫 **Sistema de Tickets de Viagem** (`gobases/desafio/`)
- **Tecnologias**: Go, CSV
- **Funcionalidades**: Análise de dados de viagem, estatísticas por destino e período
- **Conceitos**: Manipulação de arquivos, structs, funções

### 2. 🛍️ **Fantasy Products API** (`godatabase/desafio-encerramento/`)
- **Tecnologias**: Go, MySQL, Docker, Chi Router
- **Funcionalidades**: Sistema completo de vendas, migração de dados, relatórios analíticos
- **Conceitos**: Arquitetura limpa, CRUD, migrações, containerização

### 3. 📦 **API de Produtos com Testes** (`gotests/desafio-encerramento/`)
- **Tecnologias**: Go, Chi Router, Testify
- **Funcionalidades**: CRUD de produtos com cobertura completa de testes
- **Conceitos**: Testes unitários, mocks, arquitetura em camadas

### 4. ✈️ **API de Companhia Aérea** (`goweb/desafio-encerramento-goweb/`)
- **Tecnologias**: Go, Chi Router, CSV
- **Funcionalidades**: Análise estatística de voos e passageiros
- **Conceitos**: APIs REST, middlewares, Domain-Oriented Design

## 🛠️ Tecnologias Utilizadas

### **Backend**
- **Go 1.23.5** - Linguagem principal
- **Chi Router** - Framework web HTTP
- **MySQL 8.0** - Banco de dados relacional

### **Ferramentas**
- **Docker & Docker Compose** - Containerização
- **Air** - Hot reload para desenvolvimento
- **Testify** - Framework de testes
- **Make** - Automação de tarefas

### **Conceitos Aplicados**
- ✅ **Arquitetura Limpa** - Separação de responsabilidades
- ✅ **Domain-Oriented Design** - Modelagem orientada ao domínio
- ✅ **Testes Automatizados** - Cobertura >80%
- ✅ **APIs RESTful** - Endpoints padronizados
- ✅ **Containerização** - Deploy com Docker

## 🚀 Como Executar

### Pré-requisitos
- **Go 1.21+**
- **Docker & Docker Compose** (para projetos com banco)
- **Make** (recomendado)

### Execução Rápida
```bash
# Navegue para qualquer desafio
cd gobases/desafio/tickets/
cd godatabase/desafio-encerramento/
cd gotests/desafio-encerramento/
cd goweb/desafio-encerramento-goweb/

# Execute o projeto
make start
# ou
go run ./cmd/main.go
```

### Para projetos com Docker
```bash
# Suba os containers
make up-build

# Execute com hot reload
make dev
```

## 📈 Evolução do Aprendizado

| Módulo | Foco | Principais Conceitos |
|--------|------|---------------------|
| **Go Bases** | Fundamentos | Sintaxe, tipos, estruturas de controle |
| **Go Database** | Persistência | SQL, migrações, arquitetura de dados |
| **Go Tests** | Qualidade | Testes unitários, mocks, cobertura |
| **Go Web** | APIs | HTTP, REST, middlewares, JSON |

## 🎯 Destaques dos Projetos

### 🌟 **Fantasy Products API** - Projeto mais completo
- **12 endpoints REST** implementados
- **Migração automática** de dados JSON → MySQL
- **Relatórios analíticos** com consultas complexas
- **Containerização completa** com Docker Compose
- **Cobertura de testes** abrangente

### 🚀 **Funcionalidades Implementadas**
- ✅ CRUD completo para todas as entidades
- ✅ Consultas complexas com JOINs e agregações
- ✅ Recálculo automático de totais
- ✅ Top 5 clientes e produtos mais vendidos
- ✅ Análise por condição de cliente (ativo/inativo)
- ✅ Hot reload para desenvolvimento
- ✅ Testes unitários com mocks

## 👨‍💻 Sobre

**Desenvolvido por**: Esther Vitória  
**Bootcamp**: Go Web (BGW7)  
**Período**: 2025

Este repositório demonstra a evolução completa no aprendizado de Go, desde conceitos básicos até aplicações web robustas e escaláveis, seguindo as melhores práticas da linguagem.

---

## 📝 Navegação Rápida

- [🎫 Desafio Go Bases](./gobases/desafio/tickets/README.md)
- [🛍️ Fantasy Products API](./godatabase/desafio-encerramento/readme.md)
- [📦 API Produtos com Testes](./gotests/desafio-encerramento/README.md)
- [✈️ API Companhia Aérea](./goweb/desafio-encerramento-goweb/README.MD)

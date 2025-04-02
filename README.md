# Imersão FullCycle 2025

[![Go Report Card](https://goreportcard.com/badge/github.com/armandonjunior/imersao-fullcycle-2025)](https://goreportcard.com/report/github.com/armandonjunior/imersao-fullcycle-2025)
[![GoDoc](https://godoc.org/github.com/armandonjunior/imersao-fullcycle-2025?status.svg)](https://godoc.org/github.com/armandonjunior/imersao-fullcycle-2025)
[![Docker Pulls](https://img.shields.io/docker/pulls/armandonjunior/imersao-fullcycle-2025)](https://hub.docker.com/r/armandonjunior/imersao-fullcycle-2025)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://github.com/armandonjunior/imersao-fullcycle-2025/actions/workflows/go.yml/badge.svg)](https://github.com/armandonjunior/imersao-fullcycle-2025/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/armandonjunior/imersao-fullcycle-2025/branch/main/graph/badge.svg)](https://codecov.io/gh/armandonjunior/imersao-fullcycle-2025)
[![Maintainability](https://api.codeclimate.com/v1/badges/your-repo-id/maintainability)](https://codeclimate.com/github/armandonjunior/imersao-fullcycle-2025/maintainability)
[![Coverage Status](https://coveralls.io/repos/github/armandonjunior/imersao-fullcycle-2025/badge.svg?branch=main)](https://coveralls.io/github/armandonjunior/imersao-fullcycle-2025?branch=main)

<div align="center">
  <img src="https://img.shields.io/github/stars/armandonjunior/imersao-fullcycle-2025?style=social" alt="GitHub Stars">
  <img src="https://img.shields.io/github/forks/armandonjunior/imersao-fullcycle-2025?style=social" alt="GitHub Forks">
  <img src="https://img.shields.io/github/watchers/armandonjunior/imersao-fullcycle-2025?style=social" alt="GitHub Watchers">
</div>

Este é um projeto Go que implementa uma API REST simples para gerenciamento de produtos, utilizando PostgreSQL como banco de dados e Redis para cache.

## 🚀 Tecnologias Utilizadas

- [Go 1.21](https://golang.org/)
- [PostgreSQL 16](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## 📁 Estrutura do Projeto

```
.
├── main.go           # Aplicação principal em Go
├── schema.sql        # Schema do banco de dados
├── docker-compose.yaml # Configuração dos containers
├── Dockerfile        # Configuração do container da aplicação
└── go.mod           # Gerenciamento de dependências Go
```

## ⚙️ Pré-requisitos

- Docker
- Docker Compose
- Go 1.21 (para desenvolvimento local)

## 🛠️ Configuração do Ambiente

1. Clone o repositório:
```bash
git clone https://github.com/armandonjunior/imersao-fullcycle-2025.git
cd imersao-fullcycle-2025
```

2. Inicie os containers:
```bash
docker-compose up --build -d
```

## 🔌 Endpoints da API

### Listar Produtos
```bash
GET http://localhost:8000/products
```

Resposta:
```json
[
  {
    "id": 1,
    "name": "Product 1",
    "description": "Description for product 1",
    "price": 99.99
  },
  {
    "id": 2,
    "name": "Product 2",
    "description": "Description for product 2",
    "price": 149.99
  },
  {
    "id": 3,
    "name": "Product 3",
    "description": "Description for product 3",
    "price": 199.99
  }
]
```

## 💾 Estrutura do Banco de Dados

### Tabela: products
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(255))
- description (TEXT)
- price (DECIMAL(10,2))

## 🔐 Variáveis de Ambiente

A aplicação utiliza as seguintes variáveis de ambiente:

- POSTGRES_HOST: Host do PostgreSQL (default: db)
- POSTGRES_PORT: Porta do PostgreSQL (default: 5432)
- POSTGRES_USER: Usuário do PostgreSQL (default: postgres)
- POSTGRES_PASSWORD: Senha do PostgreSQL (default: postgres)
- POSTGRES_DB: Nome do banco de dados (default: postgres)

## 💻 Desenvolvimento Local

Para executar a aplicação localmente sem Docker:

1. Configure as variáveis de ambiente
2. Execute:
```bash
go mod tidy
go run main.go
```

## 🐳 Docker

O projeto utiliza três containers:

1. **app**: Container da aplicação Go
   - Porta: 8000
   - Depende do PostgreSQL

2. **db**: Container do PostgreSQL
   - Porta: 5432
   - Inclui schema inicial com dados de exemplo

3. **redis**: Container do Redis
   - Porta: 6379
   - Preparado para implementação de cache

## 💾 Volumes

- postgres_data: Armazena os dados do PostgreSQL
- redis_data: Armazena os dados do Redis

## 🏥 Health Checks

O PostgreSQL inclui health checks para garantir que o serviço esteja pronto antes de iniciar a aplicação.

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👥 Autores

- **Armando Junior** - *Trabalho Inicial* - [armandonjunior](https://github.com/armandonjunior)

## 🙏 Agradecimentos

- FullCycle
- Comunidade Go
- Todos os contribuidores

## 📈 Status do Projeto

![GitHub last commit](https://img.shields.io/github/last-commit/armandonjunior/imersao-fullcycle-2025)
![GitHub issues](https://img.shields.io/github/issues/armandonjunior/imersao-fullcycle-2025)
![GitHub pull requests](https://img.shields.io/github/issues-pr/armandonjunior/imersao-fullcycle-2025)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/armandonjunior/imersao-fullcycle-2025)

## 📞 Suporte

Se você tiver alguma dúvida ou precisar de ajuda, abra uma issue no GitHub ou entre em contato através do email: seu-email@exemplo.com

---
<div align="center">
  <sub>Built with ❤️ by <a href="https://github.com/armandonjunior">Armando Junior</a></sub>
</div>


# Imersão FullCycle 2025

Este é um projeto Go que implementa uma API REST simples para gerenciamento de produtos, utilizando PostgreSQL como banco de dados e Redis para cache.

## Tecnologias Utilizadas

- Go 1.21
- PostgreSQL 16
- Redis
- Docker
- Docker Compose

## Estrutura do Projeto

```
.
├── main.go           # Aplicação principal em Go
├── schema.sql        # Schema do banco de dados
├── docker-compose.yaml # Configuração dos containers
├── Dockerfile        # Configuração do container da aplicação
└── go.mod           # Gerenciamento de dependências Go
```

## Pré-requisitos

- Docker
- Docker Compose
- Go 1.21 (para desenvolvimento local)

## Configuração do Ambiente

1. Clone o repositório:
```bash
git clone <seu-repositorio>
cd imersao-fullcycle-2025
```

2. Inicie os containers:
```bash
docker-compose up --build -d
```

## Endpoints da API

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

## Estrutura do Banco de Dados

### Tabela: products
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(255))
- description (TEXT)
- price (DECIMAL(10,2))

## Variáveis de Ambiente

A aplicação utiliza as seguintes variáveis de ambiente:

- POSTGRES_HOST: Host do PostgreSQL (default: db)
- POSTGRES_PORT: Porta do PostgreSQL (default: 5432)
- POSTGRES_USER: Usuário do PostgreSQL (default: postgres)
- POSTGRES_PASSWORD: Senha do PostgreSQL (default: postgres)
- POSTGRES_DB: Nome do banco de dados (default: postgres)

## Desenvolvimento Local

Para executar a aplicação localmente sem Docker:

1. Configure as variáveis de ambiente
2. Execute:
```bash
go mod tidy
go run main.go
```

## Docker

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

## Volumes

- postgres_data: Armazena os dados do PostgreSQL
- redis_data: Armazena os dados do Redis

## Health Checks

O PostgreSQL inclui health checks para garantir que o serviço esteja pronto antes de iniciar a aplicação.

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.


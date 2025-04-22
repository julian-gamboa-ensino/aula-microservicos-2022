# Sistema de Otimização de Embalagem com Microsserviços e Docker

Este projeto implementa um sistema de otimização de embalagens utilizando arquitetura de microsserviços. A solução usa Docker para containerização e Docker Compose para orquestração dos serviços.

## Arquitetura

O sistema é composto por três microsserviços principais:

- **Login:** Responsável pela autenticação dos usuários e autorização de acesso aos demais serviços. Gera tokens JWT para autorização das requisições.
- **Otimizar:** Calcula a melhor forma de embalar produtos em caixas, minimizando o número de caixas usadas. Recebe a lista de produtos com suas dimensões e utiliza um algoritmo de otimização para definir a melhor alocação nas caixas disponíveis.
- **Notificação:** Envia notificações aos usuários sobre o status do pedido e informações de embalagem. Recebe o resultado da otimização e notifica os clientes via e-mail ou outro método configurado.

Os microsserviços se comunicam via uma rede interna do Docker.

## Tecnologias

- **.NET 8:** Framework principal para o desenvolvimento dos microsserviços.
- **ASP.NET Core 8:** Utilizado para criação das APIs RESTful.
- **Docker:** Plataforma de containerização.
- **Docker Compose:** Orquestrador para múltiplos contêineres Docker.
- **Swagger:** Documentação automática das APIs RESTful.
- **JWT (JSON Web Token):** Utilizado para autenticação e autorização.

## Como executar o projeto

1. **Clone o repositório:**
   ```bash
   git clone [URL-do-repositorio]
   ```

2. **Navegue até o diretório do projeto:**
   ```bash
   cd sistema-otimizacao-embalagem
   ```

3. **Construa e inicie os contêineres:**
   ```bash
   docker-compose up -d
   ```

4. **Acesse os serviços:**
   - **Login:** `http://localhost:5001/swagger`
   - **Otimizar:** `http://localhost:5002/swagger`
   - **Notificação:** `http://localhost:5003/swagger`

## Estrutura do Projeto

```
├── login
│   ├── Controllers
│   ├── Program.cs
│   └── Dockerfile
├── otimizar
│   ├── Controllers
│   ├── Business
│   ├── Services
│   ├── Models
│   ├── Program.cs
│   └── Dockerfile
├── notificacao
│   ├── Program.cs
│   └── Dockerfile
└── docker-compose.yml
```

### Detalhes dos microsserviços

- **Login:** Responsável pela autenticação dos usuários.
- **Otimizar:** Processa os pedidos e define a melhor alocação de produtos em caixas.
- **Notificação:** Envia informações sobre o status da otimização e embalagem ao usuário.

## Docker Compose

O arquivo `docker-compose.yml` configura os três microsserviços:

```yaml
services:
  login:
    build: ./login
    ports:
      - "5001:80"
    networks:
      - minha-rede

  otimizar:
    build: ./otimizar
    ports:
      - "5002:80"
    networks:
      - minha-rede

  notificacao:
    build: ./notificacao
    ports:
      - "5003:80"
    networks:
      - minha-rede

networks:
  minha-rede:
    driver: bridge
```

## Considerações

- A comunicação entre os microsserviços é feita através da rede Docker `minha-rede`.
- Cada serviço pode ser escalado de forma independente.
- O projeto usa o Swagger para documentação e teste das APIs.

## Próximos Passos

- Adicionar persistência de dados com um banco de dados (por exemplo, PostgreSQL).
- Criar testes unitários e de integração para cada serviço.
- Adicionar monitoramento e logging para acompanhar o estado dos serviços.

## Contribuições

Contribuições são bem-vindas! Para contribuir, abra uma issue ou pull request no repositório.


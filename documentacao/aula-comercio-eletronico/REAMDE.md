---

# 📘 Documentação do Projeto: Sistema de Comércio Eletrônico de Veículos

---

## 📌 Visão Geral

Este projeto simula um sistema simples de **comércio eletrônico**, com o objetivo didático de mostrar como funciona o processo de compra e venda de produtos (neste caso, veículos), incluindo o uso de carrinho de compras, controle de estoque, perfis de usuários e segurança da aplicação com Spring Security.  

---

## 🎯 Objetivos

- Aplicar conceitos de **modelagem de dados**, **microserviços**, **segurança (autenticação/autorização)**.
- Implementar um **sistema funcional** de e-commerce baseado em **Spring Boot** e **Docker**.
- Demonstrar o uso de **TDD** (Test-Driven Development).
- Ensinar boas práticas com **comentários didáticos** no código.
- Usar **JPA** com banco de dados MySQL (motor InnoDB, garantindo ACID).
- Utilizar testes automatizados e ferramentas de documentação como **OpenAPI/Swagger**.

---

## 🧱 Arquitetura do Sistema

### 🔹 Microserviços

| Microserviço               | Função Principal                                                       |
|---------------------------|------------------------------------------------------------------------|
| `login-service`           | Gerenciar login/autenticação (Spring Security)                         |
| `frontend-service`        | Servir as páginas HTML e aplicar regras de interface                   |
| `db-service`              | Banco de dados MySQL com volume persistente                            |
| `estoque-service`         | Gerenciar disponibilidade dos veículos (estoque e carrinho)            |

Cada microserviço contém seu próprio `Dockerfile` e `application-config.properties`.

---

## 🧩 Modelagem de Dados

### 🔸 Entidade: Usuário
```java
class Usuario {
    private String cpf;         // Criptografado
    private String nome;        // Criptografado
    private String login;       // Criptografado
    private String senha;       // Criptografado
    private PerfilAcesso perfil; // Enum: CLIENTE, VENDEDOR
}
```

### 🔸 Entidade: Veículo
```java
class Veiculo {
    private Long id;
    private String modelo;
    private String cor;      // branco, prata ou preta
    private int ano;
    private BigDecimal preco;
    private boolean disponivel;
}
```

### 🔸 Entidade: Carrinho
```java
class Carrinho {
    private Long id;
    private Usuario cliente;
    private List<Veiculo> veiculos;
    private LocalDateTime dataCriacao;
    private Duration tempoValidade; // 1 minuto
}
```

---

## 🔐 Segurança

- Utilizamos **Spring Security** para proteger todos os endpoints.
- A documentação da API (`OpenAPI`) só é visível após login.
- Autorização baseada em **perfil de acesso** (`VENDEDOR`, `CLIENTE`).
- Dados sensíveis criptografados (usando BCrypt/AES, conforme aplicável).

---

## 📜 Regras de Negócio Implementadas

| ID  | Descrição                                                                                  | Status     |
|-----|---------------------------------------------------------------------------------------------|------------|
| RN1 | Login com dois usuários diferentes simultaneamente                                          | ✅ Implementado |
| RN2 | Listar veículos apenas após login                                                           | ✅ Implementado |
| RN3 | Exibir botão “Usuários Cadastrados” apenas para VENDEDOR                                    | ✅ Implementado |
| RN4 | Não permitir acesso direto sem login                                                        | ✅ Implementado |
| RN5 | Ao adicionar veículo no carrinho, removê-lo da disponibilidade                              | ✅ Implementado |
| RN6 | Após tempo de validade (1 min), veículo volta ao estoque se não for comprado                | 🔄 Em progresso |
| RN7 | VENDEDOR pode realizar venda física                                                         | 🔄 Em progresso |
| RN8 | CLIENTE pode realizar compra online                                                         | ✅ Implementado |
| RN9 | Tela de listagem de usuários apenas para vendedores                                         | ✅ Implementado |
| RN10| Criptografia de dados sensíveis no banco                                                    | ✅ Implementado |

---

## 🧪 Testes Automatizados

### Back-end:
- Usamos `JUnit + Mockito`
- Pasta: `src/test/java`
- Cada teste contém comentários explicativos sobre a regra de negócio testada.

### Front-end:
- Usamos `Cypress`
- Testes simulam interações nas páginas: login, carrinho, listagem de veículos.

---

## 🖼️ Telas do Sistema

1. `login.html`: Tela de login
2. `tela_principal_listagem_veiculos.html`: Lista veículos
3. `tela_detalhamento_veiculo.html`: Exibe detalhes do veículo e botões
4. `tela_carrinho.html`: Mostra veículos no carrinho
5. `tela_secundaria_listagem_usuarios.html`: Apenas para VENDEDOR

---

## ⚙️ Configuração

### Arquivo: `application-config.properties`
```properties
# Tempo de validade do item no carrinho
TEMPO-VALIDADE-ITEMS-CARRINHO=60

# Configuração de erro visível
MENSAGEM-ERRO-TEMPO-ESGOTADO="Tempo expirado. Item voltou ao estoque."
```

---

## 🧪 TDD em Ação

### Exemplo de teste JUnit:
```java
@Test
void deveAdicionarVeiculoNoCarrinhoEAtualizarEstoque() {
    // Regra RN5: ao adicionar, deve tirar do estoque
    ...
}
```

### Exemplo de teste Cypress:
```javascript
describe('Tela Principal', () => {
  it('só exibe veículos após login', () => {
    ...
  });
});
```

---

## 🗄️ Banco de Dados (Docker)

```bash
# docker-compose.yml
services:
  mysql:
    image: mysql:8
    volumes:
      - ./data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: senha123
```


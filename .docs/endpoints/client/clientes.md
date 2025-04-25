# API de Clientes

## Visão Geral

A API de Clientes permite gerenciar informações de clientes no sistema. Ela oferece endpoints para listar, buscar, criar, atualizar e excluir clientes.

## Endpoints

### Listar Clientes

Retorna uma lista paginada de clientes.

**URL:** `/api/clients`

**Método:** `GET`

**Autenticação:** Requerida

**Parâmetros de Query:**

| Parâmetro | Tipo    | Descrição                                  | Padrão | Obrigatório |
|-----------|---------|-------------------------------------------|--------|-------------|
| page      | integer | Número da página                           | 1      | Não         |
| limit     | integer | Quantidade de registros por página (1-100) | 10     | Não         |

**Exemplo de Requisição:**
```
GET /api/clients?page=2&limit=15
```

**Exemplo de Resposta:**
```json
{
  "data": [
    {
      "id": 16,
      "name": "Cliente Exemplo",
      "email": "cliente@exemplo.com",
      "phone": "11999999999",
      "address": "Rua Exemplo, 123",
      "created_at": "2025-04-01T10:00:00Z",
      "updated_at": "2025-04-01T10:00:00Z"
    },
    {
      "id": 17,
      "name": "Outro Cliente",
      "email": "outro@exemplo.com",
      "phone": "11988888888",
      "address": "Av. Exemplo, 456",
      "created_at": "2025-04-02T11:00:00Z",
      "updated_at": "2025-04-02T11:00:00Z"
    }
  ],
  "meta": {
    "current_page": 2,
    "per_page": 15,
    "total": 45,
    "total_pages": 3
  }
}
```

### Buscar Cliente por ID

Retorna os detalhes de um cliente específico.

**URL:** `/api/clients/:id`

**Método:** `GET`

**Autenticação:** Requerida

**Parâmetros de URL:**

| Parâmetro | Tipo    | Descrição        | Obrigatório |
|-----------|---------|------------------|-------------|
| id        | integer | ID do cliente    | Sim         |

**Exemplo de Requisição:**
```
GET /api/clients/123
```

**Exemplo de Resposta:**
```json
{
  "id": 123,
  "name": "Cliente Exemplo",
  "email": "cliente@exemplo.com",
  "phone": "11999999999",
  "address": "Rua Exemplo, 123",
  "created_at": "2025-04-01T10:00:00Z",
  "updated_at": "2025-04-01T10:00:00Z"
}
```

### Criar Cliente

Cria um novo cliente.

**URL:** `/api/clients`

**Método:** `POST`

**Autenticação:** Requerida

**Corpo da Requisição:**

| Campo   | Tipo   | Descrição                | Obrigatório |
|---------|--------|--------------------------|-------------|
| name    | string | Nome do cliente          | Sim         |
| email   | string | Email do cliente         | Sim         |
| phone   | string | Telefone do cliente      | Não         |
| address | string | Endereço do cliente      | Não         |

**Exemplo de Requisição:**
```json
{
  "name": "Novo Cliente",
  "email": "novo@cliente.com",
  "phone": "11988888888",
  "address": "Av. Nova, 456"
}
```

**Exemplo de Resposta:**
```json
{
  "id": 124,
  "name": "Novo Cliente",
  "email": "novo@cliente.com",
  "phone": "11988888888",
  "address": "Av. Nova, 456",
  "created_at": "2025-04-25T14:30:00Z",
  "updated_at": "2025-04-25T14:30:00Z"
}
```

### Atualizar Cliente

Atualiza os dados de um cliente existente.

**URL:** `/api/clients/:id`

**Método:** `PUT`

**Autenticação:** Requerida

**Parâmetros de URL:**

| Parâmetro | Tipo    | Descrição        | Obrigatório |
|-----------|---------|------------------|-------------|
| id        | integer | ID do cliente    | Sim         |

**Corpo da Requisição:**

| Campo   | Tipo   | Descrição                | Obrigatório |
|---------|--------|--------------------------|-------------|
| name    | string | Nome do cliente          | Sim         |
| email   | string | Email do cliente         | Sim         |
| phone   | string | Telefone do cliente      | Não         |
| address | string | Endereço do cliente      | Não         |

**Exemplo de Requisição:**
```json
{
  "name": "Cliente Atualizado",
  "email": "atualizado@cliente.com",
  "phone": "11977777777",
  "address": "Rua Atualizada, 789"
}
```

**Exemplo de Resposta:**
```json
{
  "id": 123,
  "name": "Cliente Atualizado",
  "email": "atualizado@cliente.com",
  "phone": "11977777777",
  "address": "Rua Atualizada, 789",
  "created_at": "2025-04-01T10:00:00Z",
  "updated_at": "2025-04-25T15:45:00Z"
}
```

### Excluir Cliente

Remove um cliente do sistema.

**URL:** `/api/clients/:id`

**Método:** `DELETE`

**Autenticação:** Requerida

**Parâmetros de URL:**

| Parâmetro | Tipo    | Descrição        | Obrigatório |
|-----------|---------|------------------|-------------|
| id        | integer | ID do cliente    | Sim         |

**Exemplo de Requisição:**
```
DELETE /api/clients/123
```

**Exemplo de Resposta:**
```json
{
  "message": "Cliente removido com sucesso"
}
```

## Códigos de Status

| Código | Descrição                                                |
|--------|---------------------------------------------------------|
| 200    | Operação realizada com sucesso                          |
| 201    | Recurso criado com sucesso                              |
| 400    | Requisição inválida (dados incorretos ou incompletos)   |
| 401    | Não autorizado (autenticação necessária)                |
| 404    | Recurso não encontrado                                  |
| 500    | Erro interno do servidor                                |

# Criar Cliente

## Descrição
Este endpoint cria um novo cliente no sistema.

## URL
```
POST /api/clients
```

## Método
`POST`

## Autenticação
Este endpoint requer autenticação. O token JWT deve ser enviado no cabeçalho da requisição.

### Cabeçalho de Autenticação
```
Authorization: Bearer {token}
```

## Parâmetros de Entrada
### Corpo da Requisição (JSON)
| Campo   | Tipo   | Obrigatório | Validação                | Descrição                |
|---------|--------|-------------|--------------------------|--------------------------|
| name    | string | Sim         | min=3, max=100           | Nome do cliente          |
| email   | string | Sim         | formato de email válido  | Email do cliente         |
| phone   | string | Não         | -                        | Telefone do cliente      |
| address | string | Não         | -                        | Endereço do cliente      |

### Exemplo de Requisição
```json
{
  "name": "Nome do Cliente",
  "email": "cliente@exemplo.com",
  "phone": "11999999999",
  "address": "Endereço do Cliente"
}
```

## Resposta
### Sucesso (201 Created)
```json
{
  "id": 1,
  "name": "Nome do Cliente",
  "email": "cliente@exemplo.com",
  "phone": "11999999999",
  "address": "Endereço do Cliente",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Erro - Dados Inválidos (400 Bad Request)
```json
{
  "error": "Dados inválidos: Key: 'CreateClientRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

### Erro - Falha na Criação (500 Internal Server Error)
```json
{
  "error": "Erro ao criar cliente: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- O campo `name` deve ter entre 3 e 100 caracteres.
- O campo `email` deve ser um email válido.
- Os campos `phone` e `address` são opcionais.
- Os timestamps `created_at` e `updated_at` são gerados automaticamente pelo banco de dados.

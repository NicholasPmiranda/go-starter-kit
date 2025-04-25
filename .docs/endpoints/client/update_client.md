# Atualizar Cliente

## Descrição
Este endpoint atualiza os dados de um cliente existente.

## URL
```
PUT /api/clients/:id
```

## Método
`PUT`

## Autenticação
Este endpoint requer autenticação. O token JWT deve ser enviado no cabeçalho da requisição.

### Cabeçalho de Autenticação
```
Authorization: Bearer {token}
```

## Parâmetros de Entrada
### Parâmetros de Rota
| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| id        | int  | ID do cliente a ser atualizado |

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
  "name": "Nome Atualizado",
  "email": "email.atualizado@exemplo.com",
  "phone": "11988888888",
  "address": "Endereço Atualizado"
}
```

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Nome Atualizado",
  "email": "email.atualizado@exemplo.com",
  "phone": "11988888888",
  "address": "Endereço Atualizado",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T12:00:00Z"
}
```

### Erro - ID Inválido (400 Bad Request)
```json
{
  "error": "ID inválido"
}
```

### Erro - Dados Inválidos (400 Bad Request)
```json
{
  "error": "Key: 'UpdateClientRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

### Erro - Falha na Atualização (500 Internal Server Error)
```json
{
  "error": "Erro ao atualizar cliente: [mensagem de erro]"
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
- O timestamp `updated_at` é atualizado automaticamente pelo banco de dados.
- Todos os campos devem ser enviados na requisição, mesmo que apenas alguns estejam sendo atualizados.

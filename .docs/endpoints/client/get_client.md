# Obter Cliente

## Descrição
Este endpoint retorna os dados de um cliente específico com base no ID fornecido.

## URL
```
GET /api/clients/:id
```

## Método
`GET`

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
| id        | int  | ID do cliente a ser consultado |

## Resposta
### Sucesso (200 OK)
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

### Erro - ID Inválido (400 Bad Request)
```json
{
  "error": "ID inválido"
}
```

### Erro - Cliente Não Encontrado (404 Not Found)
```json
{
  "error": "Cliente não encontrado"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Os campos `phone` e `address` podem ser nulos.

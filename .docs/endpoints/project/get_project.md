# Obter Projeto

## Descrição
Este endpoint retorna os dados de um projeto específico com base no ID fornecido.

## URL
```
GET /api/projects/:id
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
| id        | int  | ID do projeto a ser consultado |

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Nome do Projeto",
  "description": "Descrição do Projeto",
  "client_id": 1,
  "user_id": 1,
  "status": "em_andamento",
  "start_date": "2023-01-01",
  "end_date": "2023-12-31",
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

### Erro - Projeto Não Encontrado (404 Not Found)
```json
{
  "error": "Projeto não encontrado"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Os campos `description`, `client_id`, `user_id`, `start_date` e `end_date` podem ser nulos.

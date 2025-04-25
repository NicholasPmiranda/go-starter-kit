# Obter Tarefa

## Descrição
Este endpoint retorna os dados de uma tarefa específica com base no ID fornecido.

## URL
```
GET /api/tasks/:id
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
| id        | int  | ID da tarefa a ser consultada |

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "title": "Título da Tarefa",
  "description": "Descrição detalhada da tarefa",
  "project_id": 1,
  "assigned_to": 1,
  "status": "pendente",
  "priority": "alta",
  "due_date": "2023-12-31",
  "completed_at": null,
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

### Erro - Tarefa Não Encontrada (404 Not Found)
```json
{
  "error": "Tarefa não encontrada"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Os campos `description`, `project_id`, `assigned_to`, `due_date` e `completed_at` podem ser nulos.
- O campo `completed_at` só terá valor quando a tarefa estiver concluída.

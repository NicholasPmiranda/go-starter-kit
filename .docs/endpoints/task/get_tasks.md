# Listar Tarefas

## Descrição
Este endpoint retorna uma lista de todas as tarefas cadastradas no sistema.

## URL
```
GET /api/tasks
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
Nenhum parâmetro é necessário.

## Resposta
### Sucesso (200 OK)
```json
[
  {
    "id": 1,
    "title": "Título da Tarefa 1",
    "description": "Descrição da Tarefa 1",
    "project_id": 1,
    "assigned_to": 1,
    "status": "pendente",
    "priority": "alta",
    "due_date": "2023-12-31",
    "completed_at": null,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "title": "Título da Tarefa 2",
    "description": "Descrição da Tarefa 2",
    "project_id": 1,
    "assigned_to": 2,
    "status": "em_andamento",
    "priority": "média",
    "due_date": "2023-11-30",
    "completed_at": null,
    "created_at": "2023-01-15T00:00:00Z",
    "updated_at": "2023-01-15T00:00:00Z"
  }
]
```

### Erro (500 Internal Server Error)
```json
{
  "error": "Erro ao buscar tarefas: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Este endpoint retorna todas as tarefas sem paginação.
- Os campos `description`, `project_id`, `assigned_to`, `due_date` e `completed_at` podem ser nulos.
- O campo `completed_at` só terá valor quando a tarefa estiver concluída.
- Recomenda-se implementar paginação para grandes volumes de dados.

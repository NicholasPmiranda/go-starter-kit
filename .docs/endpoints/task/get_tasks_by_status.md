# Listar Tarefas por Status

## Descrição
Este endpoint retorna uma lista de tarefas com um status específico.

## URL
```
GET /api/tasks/by-status/:status
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
| Parâmetro | Tipo   | Descrição |
|-----------|--------|-----------|
| status    | string | Status das tarefas a serem listadas (ex: pendente, em_andamento, concluída) |

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
    "id": 5,
    "title": "Título da Tarefa 5",
    "description": "Descrição da Tarefa 5",
    "project_id": 3,
    "assigned_to": 3,
    "status": "pendente",
    "priority": "baixa",
    "due_date": "2024-01-15",
    "completed_at": null,
    "created_at": "2023-04-15T00:00:00Z",
    "updated_at": "2023-04-15T00:00:00Z"
  }
]
```

### Erro - Falha na Busca (500 Internal Server Error)
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
- Este endpoint retorna todas as tarefas com o status especificado sem paginação.
- Se não houver tarefas com o status especificado, será retornada uma lista vazia.
- Os valores comuns para o parâmetro `status` são: "pendente", "em_andamento", "concluída", mas podem variar de acordo com as regras de negócio da aplicação.
- Os campos `description`, `project_id`, `assigned_to`, `due_date` e `completed_at` podem ser nulos.
- O campo `completed_at` só terá valor quando a tarefa estiver com status "concluída".
- Recomenda-se implementar paginação para grandes volumes de dados.

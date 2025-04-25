# Listar Tarefas por Usuário

## Descrição
Este endpoint retorna uma lista de tarefas atribuídas a um usuário específico.

## URL
```
GET /api/tasks/by-user/:user_id
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
| user_id   | int  | ID do usuário cujas tarefas serão listadas |

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
    "id": 4,
    "title": "Título da Tarefa 4",
    "description": "Descrição da Tarefa 4",
    "project_id": 2,
    "assigned_to": 1,
    "status": "em_andamento",
    "priority": "média",
    "due_date": "2023-11-15",
    "completed_at": null,
    "created_at": "2023-03-15T00:00:00Z",
    "updated_at": "2023-03-15T00:00:00Z"
  }
]
```

### Erro - ID do Usuário Inválido (400 Bad Request)
```json
{
  "error": "ID do usuário inválido"
}
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
- Este endpoint retorna todas as tarefas atribuídas ao usuário especificado sem paginação.
- Se não houver tarefas atribuídas ao usuário, será retornada uma lista vazia.
- Os campos `description`, `project_id`, `due_date` e `completed_at` podem ser nulos.
- O campo `completed_at` só terá valor quando a tarefa estiver concluída.
- Recomenda-se implementar paginação para grandes volumes de dados.

# Listar Tarefas por Projeto

## Descrição
Este endpoint retorna uma lista de tarefas associadas a um projeto específico.

## URL
```
GET /api/tasks/by-project/:project_id
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
| Parâmetro  | Tipo | Descrição |
|------------|------|-----------|
| project_id | int  | ID do projeto cujas tarefas serão listadas |

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
    "id": 3,
    "title": "Título da Tarefa 3",
    "description": "Descrição da Tarefa 3",
    "project_id": 1,
    "assigned_to": 2,
    "status": "concluída",
    "priority": "baixa",
    "due_date": "2023-10-15",
    "completed_at": "2023-10-10T00:00:00Z",
    "created_at": "2023-02-15T00:00:00Z",
    "updated_at": "2023-10-10T00:00:00Z"
  }
]
```

### Erro - ID do Projeto Inválido (400 Bad Request)
```json
{
  "error": "ID do projeto inválido"
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
- Este endpoint retorna todas as tarefas associadas ao projeto especificado sem paginação.
- Se não houver tarefas associadas ao projeto, será retornada uma lista vazia.
- Os campos `description`, `assigned_to`, `due_date` e `completed_at` podem ser nulos.
- O campo `completed_at` só terá valor quando a tarefa estiver concluída.
- Recomenda-se implementar paginação para grandes volumes de dados.

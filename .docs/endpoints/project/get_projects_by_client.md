# Listar Projetos por Cliente

## Descrição
Este endpoint retorna uma lista de projetos associados a um cliente específico.

## URL
```
GET /api/projects/by-client/:client_id
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
| client_id | int  | ID do cliente cujos projetos serão listados |

## Resposta
### Sucesso (200 OK)
```json
[
  {
    "id": 1,
    "name": "Nome do Projeto 1",
    "description": "Descrição do Projeto 1",
    "client_id": 1,
    "user_id": 1,
    "status": "em_andamento",
    "start_date": "2023-01-01",
    "end_date": "2023-12-31",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "name": "Nome do Projeto 3",
    "description": "Descrição do Projeto 3",
    "client_id": 1,
    "user_id": 2,
    "status": "concluído",
    "start_date": "2023-03-01",
    "end_date": "2023-06-30",
    "created_at": "2023-02-15T00:00:00Z",
    "updated_at": "2023-06-30T00:00:00Z"
  }
]
```

### Erro - ID do Cliente Inválido (400 Bad Request)
```json
{
  "error": "ID do cliente inválido"
}
```

### Erro - Falha na Busca (500 Internal Server Error)
```json
{
  "error": "Erro ao buscar projetos: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Este endpoint retorna todos os projetos associados ao cliente especificado sem paginação.
- Se não houver projetos associados ao cliente, será retornada uma lista vazia.
- Os campos `description`, `user_id`, `start_date` e `end_date` podem ser nulos.
- Recomenda-se implementar paginação para grandes volumes de dados.

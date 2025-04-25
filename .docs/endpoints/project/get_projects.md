# Listar Projetos

## Descrição
Este endpoint retorna uma lista de todos os projetos cadastrados no sistema.

## URL
```
GET /api/projects
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
    "id": 2,
    "name": "Nome do Projeto 2",
    "description": "Descrição do Projeto 2",
    "client_id": 2,
    "user_id": 1,
    "status": "planejado",
    "start_date": "2023-02-01",
    "end_date": "2023-12-31",
    "created_at": "2023-01-15T00:00:00Z",
    "updated_at": "2023-01-15T00:00:00Z"
  }
]
```

### Erro (500 Internal Server Error)
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
- Este endpoint retorna todos os projetos sem paginação.
- Os campos `description`, `client_id`, `user_id`, `start_date` e `end_date` podem ser nulos.
- Recomenda-se implementar paginação para grandes volumes de dados.

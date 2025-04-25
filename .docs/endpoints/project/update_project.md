# Atualizar Projeto

## Descrição
Este endpoint atualiza os dados de um projeto existente.

## URL
```
PUT /api/projects/:id
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
| id        | int  | ID do projeto a ser atualizado |

### Corpo da Requisição (JSON)
| Campo       | Tipo   | Obrigatório | Descrição                                |
|-------------|--------|-------------|------------------------------------------|
| name        | string | Sim         | Nome do projeto                          |
| description | string | Não         | Descrição detalhada do projeto           |
| client_id   | int    | Não         | ID do cliente associado ao projeto       |
| user_id     | int    | Não         | ID do usuário responsável pelo projeto   |
| status      | string | Sim         | Status do projeto (ex: planejado, em_andamento, concluído) |
| start_date  | string | Não         | Data de início do projeto (formato YYYY-MM-DD) |
| end_date    | string | Não         | Data de término do projeto (formato YYYY-MM-DD) |

### Exemplo de Requisição
```json
{
  "name": "Projeto Atualizado",
  "description": "Descrição atualizada do projeto",
  "client_id": 2,
  "user_id": 3,
  "status": "em_andamento",
  "start_date": "2023-06-15",
  "end_date": "2024-01-31"
}
```

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Projeto Atualizado",
  "description": "Descrição atualizada do projeto",
  "client_id": 2,
  "user_id": 3,
  "status": "em_andamento",
  "start_date": "2023-06-15",
  "end_date": "2024-01-31",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-05-20T14:45:00Z"
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
  "error": "Dados inválidos: [mensagem de erro]"
}
```

### Erro - Falha na Atualização (500 Internal Server Error)
```json
{
  "error": "Erro ao atualizar projeto: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Os campos `description`, `client_id`, `user_id`, `start_date` e `end_date` são opcionais.
- O campo `status` deve conter um valor válido de acordo com as regras de negócio da aplicação.
- As datas devem ser fornecidas no formato YYYY-MM-DD.
- O timestamp `updated_at` é atualizado automaticamente pelo banco de dados.
- Todos os campos devem ser enviados na requisição, mesmo que apenas alguns estejam sendo atualizados.

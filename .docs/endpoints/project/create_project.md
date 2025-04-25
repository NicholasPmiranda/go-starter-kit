# Criar Projeto

## Descrição
Este endpoint cria um novo projeto no sistema.

## URL
```
POST /api/projects
```

## Método
`POST`

## Autenticação
Este endpoint requer autenticação. O token JWT deve ser enviado no cabeçalho da requisição.

### Cabeçalho de Autenticação
```
Authorization: Bearer {token}
```

## Parâmetros de Entrada
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
  "name": "Novo Projeto",
  "description": "Descrição detalhada do novo projeto",
  "client_id": 1,
  "user_id": 1,
  "status": "planejado",
  "start_date": "2023-06-01",
  "end_date": "2023-12-31"
}
```

## Resposta
### Sucesso (201 Created)
```json
{
  "id": 3,
  "name": "Novo Projeto",
  "description": "Descrição detalhada do novo projeto",
  "client_id": 1,
  "user_id": 1,
  "status": "planejado",
  "start_date": "2023-06-01",
  "end_date": "2023-12-31",
  "created_at": "2023-05-15T10:30:00Z",
  "updated_at": "2023-05-15T10:30:00Z"
}
```

### Erro - Dados Inválidos (400 Bad Request)
```json
{
  "error": "Dados inválidos: [mensagem de erro]"
}
```

### Erro - Falha na Criação (500 Internal Server Error)
```json
{
  "error": "Erro ao criar projeto: [mensagem de erro]"
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
- Os timestamps `created_at` e `updated_at` são gerados automaticamente pelo banco de dados.

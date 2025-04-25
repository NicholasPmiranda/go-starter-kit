# Listar Clientes

## Descrição
Este endpoint retorna uma lista de todos os clientes cadastrados no sistema.

## URL
```
GET /api/clients
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
    "name": "Nome do Cliente 1",
    "email": "cliente1@exemplo.com",
    "phone": "11999999999",
    "address": "Endereço do Cliente 1",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "name": "Nome do Cliente 2",
    "email": "cliente2@exemplo.com",
    "phone": "11988888888",
    "address": "Endereço do Cliente 2",
    "created_at": "2023-01-02T00:00:00Z",
    "updated_at": "2023-01-02T00:00:00Z"
  }
]
```

### Erro (500 Internal Server Error)
```json
{
  "error": "Erro ao buscar clientes: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Este endpoint retorna todos os clientes sem paginação.
- Recomenda-se implementar paginação para grandes volumes de dados.

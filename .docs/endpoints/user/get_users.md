# Listar Usuários

## Descrição
Este endpoint retorna uma lista de todos os usuários cadastrados no sistema.

## URL
```
GET /api/users
```

## Método
`GET`

## Parâmetros de Entrada
Nenhum parâmetro é necessário.

## Resposta
### Sucesso (200 OK)
```json
[
  {
    "id": 1,
    "name": "Nome do Usuário 1",
    "email": "usuario1@exemplo.com",
    "password": "[hash da senha]"
  },
  {
    "id": 2,
    "name": "Nome do Usuário 2",
    "email": "usuario2@exemplo.com",
    "password": "[hash da senha]"
  }
]
```

### Erro (500 Internal Server Error)
```json
{
  "error": "Erro ao buscar usuários: [mensagem de erro]"
}
```

## Observações
- Este endpoint retorna todos os usuários sem paginação.
- Por questões de segurança, em um ambiente de produção, o campo `password` não deveria ser retornado na resposta.
- Recomenda-se implementar paginação para grandes volumes de dados.

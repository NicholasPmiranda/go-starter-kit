# Obter Usuário

## Descrição
Este endpoint retorna os dados de um usuário específico com base no ID fornecido.

## URL
```
GET /api/users/:id
```

## Método
`GET`

## Parâmetros de Entrada
### Parâmetros de Rota
| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| id        | int  | ID do usuário a ser consultado |

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Nome do Usuário",
  "email": "usuario@exemplo.com",
  "password": "[hash da senha]"
}
```

### Erro - ID Inválido (400 Bad Request)
```json
{
  "error": "ID inválido"
}
```

### Erro - Usuário Não Encontrado (404 Not Found)
```json
{
  "error": "Usuário não encontrado"
}
```

## Observações
- Por questões de segurança, em um ambiente de produção, o campo `password` não deveria ser retornado na resposta.

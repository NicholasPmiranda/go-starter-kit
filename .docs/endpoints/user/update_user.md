# Atualizar Usuário

## Descrição
Este endpoint atualiza os dados de um usuário existente.

## URL
```
PUT /api/users/:id
```

## Método
`PUT`

## Parâmetros de Entrada
### Parâmetros de Rota
| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| id        | int  | ID do usuário a ser atualizado |

### Corpo da Requisição (JSON)
| Campo    | Tipo   | Obrigatório | Descrição                |
|----------|--------|-------------|--------------------------|
| name     | string | Sim         | Nome do usuário          |
| email    | string | Sim         | Email do usuário         |
| password | string | Sim         | Senha do usuário         |

### Exemplo de Requisição
```json
{
  "name": "Nome Atualizado",
  "email": "email.atualizado@exemplo.com",
  "password": "novaSenha123"
}
```

## Resposta
### Sucesso (200 OK)
```json
{
  "id": 1,
  "name": "Nome Atualizado",
  "email": "email.atualizado@exemplo.com",
  "password": "[hash da senha]"
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
  "error": "Erro ao atualizar usuário: [mensagem de erro]"
}
```

## Observações
- Todos os campos devem ser enviados na requisição, mesmo que apenas alguns estejam sendo atualizados.
- O campo `password` é atualizado como um hash no banco de dados.
- Por questões de segurança, em um ambiente de produção, o campo `password` não deveria ser retornado na resposta.
- Recomenda-se implementar validações adicionais para garantir a segurança da senha (comprimento mínimo, caracteres especiais, etc.).

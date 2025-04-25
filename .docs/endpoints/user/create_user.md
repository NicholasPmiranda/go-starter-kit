# Criar Usuário

## Descrição
Este endpoint cria um novo usuário no sistema.

## URL
```
POST /api/users
```

## Método
`POST`

## Parâmetros de Entrada
### Corpo da Requisição (JSON)
| Campo    | Tipo   | Obrigatório | Descrição                |
|----------|--------|-------------|--------------------------|
| name     | string | Sim         | Nome do usuário          |
| email    | string | Sim         | Email do usuário         |
| password | string | Sim         | Senha do usuário         |

### Exemplo de Requisição
```json
{
  "name": "Nome do Usuário",
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

## Resposta
### Sucesso (201 Created)
```json
{
  "id": 1,
  "name": "Nome do Usuário",
  "email": "usuario@exemplo.com",
  "password": "[hash da senha]"
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
  "error": "Erro ao criar usuário: [mensagem de erro]"
}
```

## Observações
- O campo `password` é armazenado como um hash no banco de dados, mas é retornado na resposta.
- Por questões de segurança, em um ambiente de produção, o campo `password` não deveria ser retornado na resposta.
- Recomenda-se implementar validações adicionais para garantir a segurança da senha (comprimento mínimo, caracteres especiais, etc.).

# Login

## Descrição
Este endpoint realiza a autenticação do usuário no sistema, validando suas credenciais e retornando um token JWT para ser utilizado nas requisições subsequentes.

## URL
```
POST /api/login
```

## Método
`POST`

## Parâmetros de Entrada
### Corpo da Requisição (JSON)
| Campo    | Tipo   | Obrigatório | Descrição                |
|----------|--------|-------------|--------------------------|
| email    | string | Sim         | Email do usuário         |
| password | string | Sim         | Senha do usuário         |

### Exemplo de Requisição
```json
{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

## Resposta
### Sucesso (200 OK)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Erro (400 Bad Request)
```json
{
  "error": "usuario e ou senha incorretos"
}
```

### Erro de Validação (400 Bad Request)
```json
{
  "error": "Key: 'email' Error:Field validation for 'email' failed on the 'required' tag"
}
```

## Observações
- O token JWT gerado tem validade de 24 horas.
- Este token deve ser incluído no cabeçalho `Authorization` das requisições subsequentes, no formato `Bearer {token}`.

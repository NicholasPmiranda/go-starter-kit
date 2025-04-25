# Perfil do Usuário

## Descrição
Este endpoint retorna as informações do perfil do usuário autenticado.

## URL
```
GET /api/profile
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
Nenhum parâmetro adicional é necessário.

## Resposta
### Sucesso (200 OK)
```json
{
  "user": {
    "id": 1,
    "name": "Nome do Usuário",
    "email": "usuario@exemplo.com",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Este endpoint só pode ser acessado por usuários autenticados.
- O objeto `user` contém todas as informações do usuário autenticado, exceto a senha.

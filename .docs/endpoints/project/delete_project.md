# Excluir Projeto

## Descrição
Este endpoint remove um projeto do sistema.

## URL
```
DELETE /api/projects/:id
```

## Método
`DELETE`

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
| id        | int  | ID do projeto a ser excluído |

## Resposta
### Sucesso (200 OK)
```json
{
  "message": "Projeto removido com sucesso"
}
```

### Erro - ID Inválido (400 Bad Request)
```json
{
  "error": "ID inválido"
}
```

### Erro - Falha na Exclusão (500 Internal Server Error)
```json
{
  "error": "Erro ao remover projeto: [mensagem de erro]"
}
```

### Erro de Autenticação (401 Unauthorized)
```json
{
  "error": "Token inválido ou expirado"
}
```

## Observações
- Esta operação é irreversível. Uma vez excluído, o projeto não pode ser recuperado.
- Recomenda-se implementar uma exclusão lógica (soft delete) em vez de uma exclusão física dos dados, especialmente em ambientes de produção.
- A exclusão de um projeto pode afetar outros registros relacionados, como tarefas, subtarefas, comentários e anexos associados a este projeto.

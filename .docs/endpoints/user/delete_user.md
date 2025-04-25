# Excluir Usuário

## Descrição
Este endpoint remove um usuário do sistema.

## URL
```
DELETE /api/users/:id
```

## Método
`DELETE`

## Parâmetros de Entrada
### Parâmetros de Rota
| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| id        | int  | ID do usuário a ser excluído |

## Resposta
### Sucesso (200 OK)
```json
{
  "message": "Usuário removido com sucesso"
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
  "error": "Erro ao remover usuário: [mensagem de erro]"
}
```

## Observações
- Esta operação é irreversível. Uma vez excluído, o usuário não pode ser recuperado.
- Recomenda-se implementar uma exclusão lógica (soft delete) em vez de uma exclusão física dos dados, especialmente em ambientes de produção.

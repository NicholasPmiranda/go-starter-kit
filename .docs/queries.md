# Como Criar Queries SQL

## Visão Geral

O Go Starter Kit utiliza [SQLC](https://sqlc.dev/) para gerar código Go a partir de consultas SQL. Isso proporciona uma camada de acesso a dados fortemente tipada e eficiente, eliminando a necessidade de escrever código boilerplate para mapear resultados de consultas para structs Go.

## Estrutura de Arquivos

As queries SQL são organizadas da seguinte forma:

```
database/
├── migrations/       # Migrações de banco de dados
├── query/           # Arquivos de consultas SQL
│   ├── user.sql     # Consultas relacionadas a usuários
│   └── shipments.sql # Consultas relacionadas a envios
├── schema/          # Definições de esquema
└── seeds/           # Dados iniciais
```

## Como Definir Novas Queries

### 1. Criar ou Editar um Arquivo SQL

Para adicionar novas consultas, você deve criar ou editar um arquivo SQL no diretório `database/query/`:

```sql
-- name: FindById :one
select * from users where id = $1;

-- name: FindByEmail :one
select * from users WHERE email = $1 limit 1;

-- name: FindMany :many
select * from users;

-- name: CreateUser :one
insert into users (name, email, password)
 values ($1, $2, $3) returning *;
```

### 2. Anotações SQLC

Cada consulta deve ter uma anotação SQLC que define:

1. **Nome da função**: O nome da função Go que será gerada
2. **Tipo de retorno**: O tipo de retorno da função
   - `:one` - Retorna um único registro
   - `:many` - Retorna múltiplos registros
   - `:exec` - Não retorna registros (para INSERT, UPDATE, DELETE sem RETURNING)
   - `:execrows` - Retorna o número de linhas afetadas
   - `:execresult` - Retorna o resultado completo da execução

Exemplo:
```sql
-- name: FindById :one
select * from users where id = $1;
```

Esta anotação gerará uma função Go chamada `FindById` que retorna um único registro.

### 3. Parâmetros de Consulta

Os parâmetros de consulta são definidos usando a sintaxe `$n`, onde `n` é o número do parâmetro:

```sql
-- name: FindByEmailAndStatus :many
select * from users where email = $1 and status = $2;
```

### 4. Gerar o Código Go

Após definir suas consultas, você deve gerar o código Go correspondente usando o comando:

```bash
make sqlc
```

ou diretamente:

```bash
sqlc generate
```

Isso irá gerar arquivos Go no diretório `internal/database/` com funções que correspondem às suas consultas SQL.

## Tipos de Consultas

### Consulta que Retorna um Único Registro

```sql
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
```

Função Go gerada:
```go
func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
    // ...
}
```

### Consulta que Retorna Múltiplos Registros

```sql
-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;
```

Função Go gerada:
```go
func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
    // ...
}
```

### Consulta de Inserção com Retorno

```sql
-- name: CreateUser :one
INSERT INTO users (
    name, email, password
) VALUES (
    $1, $2, $3
)
RETURNING *;
```

Função Go gerada:
```go
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
    // ...
}
```

### Consulta de Atualização

```sql
-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3
WHERE id = $1
RETURNING *;
```

Função Go gerada:
```go
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
    // ...
}
```

### Consulta de Exclusão

```sql
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
```

Função Go gerada:
```go
func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
    // ...
}
```

## Como Usar as Queries Geradas

### 1. Obter uma Instância de Queries

```go
import (
    "context"
    "database/sql"
    "boilerPlate/internal/database"
)

func exemplo(db *sql.DB) {
    queries := database.New(db)

    // Agora você pode usar as funções geradas
    // ...
}
```

### 2. Executar Consultas

```go
// Buscar um usuário por ID
user, err := queries.FindById(context.Background(), 123)
if err != nil {
    // Tratar erro
}

// Criar um novo usuário
newUser, err := queries.CreateUser(context.Background(), database.CreateUserParams{
    Name:     "João Silva",
    Email:    "joao@exemplo.com",
    Password: "senha_hash",
})
if err != nil {
    // Tratar erro
}

// Listar todos os usuários
users, err := queries.FindMany(context.Background())
if err != nil {
    // Tratar erro
}
```

### 3. Usar Transações

```go
// Iniciar uma transação
tx, err := db.BeginTx(context.Background(), nil)
if err != nil {
    return err
}
defer tx.Rollback()

// Criar queries para a transação
qtx := database.New(tx)

// Executar consultas dentro da transação
user, err := qtx.CreateUser(context.Background(), database.CreateUserParams{
    Name:     "Maria Silva",
    Email:    "maria@exemplo.com",
    Password: "senha_hash",
})
if err != nil {
    return err
}

// Mais operações...

// Commit da transação
if err := tx.Commit(); err != nil {
    return err
}
```

## Configuração do SQLC

A configuração do SQLC é definida no arquivo `sqlc.yaml` na raiz do projeto:

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "database/query/"
    schema: "database/schema/"
    gen:
      go:
        package: "database"
        out: "internal/database"
        emit_json_tags: true
        emit_prepared_queries: true
        emit_interface: true
        emit_exact_table_names: false
```

## Boas Práticas

1. **Nomeação**: Use nomes descritivos para suas consultas que reflitam a operação sendo realizada.
2. **Organização**: Agrupe consultas relacionadas no mesmo arquivo SQL.
3. **Comentários**: Adicione comentários para explicar consultas complexas.
4. **Parâmetros Nomeados**: Para consultas complexas, considere usar parâmetros nomeados (suportados pelo SQLC).
5. **Validação**: Valide os dados antes de passá-los para as consultas para evitar injeção de SQL.
6. **Transações**: Use transações para operações que envolvem múltiplas consultas que precisam ser atômicas.
7. **Paginação**: Implemente paginação para consultas que podem retornar muitos registros.

## SQLx vs SQLC: Quando Usar Cada Um

### O que é SQLx?

[SQLx](https://github.com/jmoiron/sqlx) é uma biblioteca que estende o pacote `database/sql` padrão do Go, adicionando funcionalidades para facilitar o trabalho com bancos de dados. Diferentemente do SQLC, que gera código a partir de consultas SQL, o SQLx é uma biblioteca de tempo de execução que oferece métodos auxiliares para mapear resultados de consultas para structs Go.

### Como o SQLx é Usado no Projeto

O Go Starter Kit inclui suporte para SQLx através da função `ConnectDBX()` no pacote `internal/database`. Esta função estabelece uma conexão com o banco de dados PostgreSQL usando o driver pgx e retorna um objeto `*sqlx.DB`.

Exemplo de conexão com SQLx:

```go
// Conecta ao banco de dados usando sqlx
db := database.ConnectDBX()
defer db.Close()
```

### Principais Recursos do SQLx

1. **Mapeamento Automático**: Mapeia resultados de consultas para structs Go usando tags de struct.
2. **Métodos Auxiliares**: Oferece métodos como `Get`, `Select`, `QueryRowx`, e `StructScan` para facilitar o trabalho com resultados de consultas.
3. **Preparação de Consultas**: Suporta consultas preparadas para melhor desempenho e segurança.
4. **Transações**: Facilita o trabalho com transações.

### Exemplo de Uso do SQLx

```go
// Inserção com retorno de dados
var novoUsuario database.User
err := db.QueryRowx(
    "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *",
    "Novo Usuário", "usuario@exemplo.com", "senha123",
).StructScan(&novoUsuario)

// Consulta de múltiplos registros
var usuarios []database.User
err = db.Select(&usuarios, "SELECT id, name, email, password FROM users")
```

### Quando Usar SQLC vs SQLx

#### Use SQLC quando:

1. **Consultas Conhecidas em Tempo de Compilação**: Suas consultas SQL são conhecidas antecipadamente e não mudam dinamicamente.
2. **Verificação de Tipo em Tempo de Compilação**: Você deseja verificação de tipo em tempo de compilação para suas consultas SQL.
3. **Geração de Código**: Você prefere ter código Go gerado automaticamente a partir de suas consultas SQL.
4. **Consultas Complexas**: Você tem consultas SQL complexas que se beneficiariam de verificação de sintaxe em tempo de compilação.
5. **Manutenção de API**: Você deseja uma API estável e bem definida para acesso a dados.

#### Use SQLx quando:

1. **Consultas Dinâmicas**: Você precisa construir consultas SQL dinamicamente em tempo de execução.
2. **Flexibilidade**: Você precisa de mais flexibilidade no mapeamento de resultados para structs.
3. **Prototipagem Rápida**: Você está prototipando rapidamente e não quer gerar código a cada alteração de consulta.
4. **Consultas Ad-hoc**: Você precisa executar consultas ad-hoc que não são conhecidas em tempo de compilação.
5. **Integração com Código Existente**: Você está integrando com código existente que já usa `database/sql`.

### Combinando SQLC e SQLx

Em muitos projetos, é benéfico usar tanto SQLC quanto SQLx:

- Use **SQLC** para consultas comuns e bem definidas que são conhecidas em tempo de compilação.
- Use **SQLx** para consultas dinâmicas ou ad-hoc que precisam ser construídas em tempo de execução.

O Go Starter Kit suporta ambas as abordagens, permitindo que você escolha a ferramenta certa para cada caso de uso.

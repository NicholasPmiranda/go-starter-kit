# Go Starter Kit com Docker

Este projeto utiliza Docker para facilitar o desenvolvimento e execução da aplicação. Abaixo estão as instruções para utilizar os comandos disponíveis.

## Requisitos

- Docker
- Docker Compose
- Make

## Configuração Inicial

1. Clone o repositório
2. Configure o arquivo `.env` com suas variáveis de ambiente
3. entre no arquivo setup.sh troque o nome do projeto e execute o arquivo
4. Execute o comando para construir as imagens Docker:

```bash
make build
```

4. Inicie os containers:

```bash
make up
```

## Comandos Disponíveis

### Comandos Docker

| Comando | Descrição |
|---------|-----------|
| `make up` | Inicia os containers Docker em modo detached (background) |
| `make down` | Para os containers Docker e remove os recursos criados |
| `make restart` | Reinicia os containers Docker sem recriá-los |
| `make build` | Constrói ou reconstrói as imagens Docker conforme definido no docker-compose.yaml |
| `make logs` | Exibe os logs dos containers em tempo real (use Ctrl+C para sair) |
| `make shell` | Acessa o shell (sh) do container da aplicação para executar comandos manualmente |
| `make docker-exec [comando]` | Executa um comando específico no container da aplicação |

### Comandos de Desenvolvimento

| Comando | Descrição |
|---------|-----------|
| `make docker-dev` | Executa a aplicação com hot reload usando Air dentro do container (alterações no código são aplicadas automaticamente) |
| `make docker-run` | Executa a aplicação sem hot reload dentro do container (precisa reiniciar manualmente após alterações) |
| `make docker-sqlc` | Executa o SQLc para gerar código Go a partir das consultas SQL definidas em database/query |

### Comandos de Migração

| Comando | Descrição |
|---------|-----------|
| `make migration [nome]` | Cria uma nova migração com o nome especificado (ex: `make migration create_users_table`) |
| `make migrate` | Aplica todas as migrações pendentes no banco de dados local |
| `make migrate:fresh` | Dropa todo o banco de dados e aplica todas as migrações novamente no ambiente local |
| `make migrate:rollback` | Reverte a última migração aplicada no ambiente local |
| `make docker-migrate` | Executa todas as migrações pendentes dentro do container Docker |
| `make docker-migrate:fresh` | Dropa o banco e executa todas as migrações dentro do container Docker |
| `make docker-migrate:rollback` | Reverte a última migração aplicada dentro do container Docker |

### Comandos de Instalação

| Comando | Descrição |
|---------|-----------|
| `make install-migrate` | Instala a ferramenta golang-migrate para gerenciamento de migrações |
| `make install-air` | Instala o Air para hot reload durante o desenvolvimento |
| `make install-sqlc` | Instala o SQLc para geração de código a partir de SQL |

### Comandos de Execução Local

| Comando | Descrição |
|---------|-----------|
| `make dev` | Executa a aplicação localmente com hot reload usando Air |
| `make run` | Executa a aplicação localmente sem hot reload |

## Estrutura do Projeto

O projeto segue a estrutura recomendada para aplicações Go:

```
/
├── cmd/                    # Pontos de entrada da aplicação
│   ├── server/             # Servidor principal
│   └── worker/             # Worker para processamento em background
├── internal/               # Código específico da aplicação
│   ├── domain/             # Modelos de domínio e regras de negócio
│   ├── service/            # Serviços da aplicação
│   ├── repository/         # Implementações de acesso a dados
│   ├── handler/            # Manipuladores HTTP
│   ├── database/           # Código gerado pelo SQLc
│   └── usecase/            # Casos de uso da aplicação
├── database/               # Código de suporte para banco de dados
│   ├── migrations/         # Migrações
│   ├── query/              # Consultas SQL para o SQLc
│   ├── schema/             # Definições de esquema para o SQLc
│   └── seeds/              # Dados iniciais
├── config/                 # Configurações da aplicação
├── helpers/                # Funções auxiliares
├── routes/                 # Definição de rotas
├── storage/                # Armazenamento de arquivos
└── template/               # Templates HTML
```

### Arquivos de Configuração Importantes

- `.env`: Variáveis de ambiente para configuração local
- `docker-compose.yaml`: Configuração dos serviços Docker
- `Dockerfile`: Instruções para construir a imagem Docker da aplicação
- `sqlc.yaml`: Configuração do SQLc para geração de código
- `Makefile`: Comandos para facilitar o desenvolvimento

## Ferramentas Incluídas

O container Docker da aplicação inclui as seguintes ferramentas:

- **Go**: Linguagem de programação
- **Air**: Para hot-reload durante o desenvolvimento
- **SQLc**: Para gerar código Go a partir de consultas SQL
- **Golang Migrations**: Para gerenciar migrações de banco de dados

## Guia Detalhado

### Como Iniciar o Desenvolvimento

Para iniciar o desenvolvimento com hot-reload:

```bash
# Inicia os containers
make up

# Executa a aplicação com hot-reload
make docker-dev
```

A aplicação estará disponível em `http://localhost:8080` (ou na porta definida na variável de ambiente APP_PORT).

### Como Trabalhar com Migrações

As migrações permitem versionar seu banco de dados. O projeto usa a ferramenta `golang-migrate`.

#### Criando uma Nova Migração

```bash
# Cria uma nova migração (substitua 'nome_da_migracao' pelo nome desejado)
make migration nome_da_migracao
```

Isso criará dois arquivos na pasta `database/migrations`:
- `YYYYMMDDHHMMSS_nome_da_migracao.up.sql`: Para aplicar a migração
- `YYYYMMDDHHMMSS_nome_da_migracao.down.sql`: Para reverter a migração

Edite esses arquivos para adicionar suas instruções SQL.

#### Aplicando Migrações

```bash
# Dentro do container Docker (recomendado)
make docker-migrate

# Ou localmente
make migrate
```

#### Revertendo Migrações

```bash
# Reverte a última migração (dentro do container)
make docker-migrate:rollback

# Ou localmente
make migrate:rollback
```

#### Recriando o Banco de Dados

```bash
# Dropa o banco e aplica todas as migrações (dentro do container)
make docker-migrate:fresh

# Ou localmente
make migrate:fresh
```

### Como Usar o SQLc

O SQLc gera código Go a partir de consultas SQL, oferecendo type safety e melhor desempenho.

#### Estrutura de Arquivos do SQLc

- `database/schema/`: Contém os arquivos de definição de esquema (CREATE TABLE, etc.)
- `database/query/`: Contém as consultas SQL que serão convertidas em código Go
- `internal/database/`: Onde o código gerado será armazenado

#### Criando uma Nova Consulta

1. Crie ou atualize o esquema em `database/schema/`
2. Crie um arquivo SQL em `database/query/` com suas consultas

Exemplo de consulta em `database/query/user.sql`:

```sql
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;
```

3. Gere o código Go:

```bash
make docker-sqlc
```

4. O código gerado estará disponível em `internal/database/` e poderá ser importado em seu código.

#### Usando o Código Gerado

```go
import "github.com/seu-usuario/seu-projeto/internal/database"

// Em algum lugar do seu código
user, err := queries.GetUserByID(ctx, userID)
```

### Executando a Aplicação

```bash
# Com hot-reload (desenvolvimento)
make docker-dev

# Sem hot-reload (produção)
make docker-run

# Localmente com hot-reload
make dev

# Localmente sem hot-reload
make run
```

## Fluxo de Trabalho Recomendado

Para um desenvolvimento eficiente com este projeto, recomendamos o seguinte fluxo de trabalho:

1. **Configuração Inicial**:
   - Clone o repositório
   - Configure o arquivo `.env`
   - Execute `make build` e `make up` para iniciar os containers

2. **Desenvolvimento**:
   - Execute `make docker-dev` para iniciar a aplicação com hot-reload
   - Desenvolva suas funcionalidades
   - Use `make docker-sqlc` quando modificar arquivos SQL
   - Use `make migration nome_migracao` para criar novas migrações
   - Use `make docker-migrate` para aplicar migrações

3. **Testes**:
   - Escreva testes para suas funcionalidades
   - Execute os testes localmente ou no container

4. **Implantação**:
   - Use `make build` para reconstruir as imagens
   - Implante conforme necessário

## Dicas e Solução de Problemas

### Problemas Comuns

1. **Erro de conexão com o banco de dados**:
   - Verifique se o container do PostgreSQL está em execução: `docker ps`
   - Verifique as variáveis de ambiente no arquivo `.env`
   - Tente reiniciar os containers: `make restart`

2. **Erro ao executar migrações**:
   - Verifique se o banco de dados está acessível
   - Verifique a sintaxe SQL nas migrações
   - Use `make migrate:fresh` para recriar o banco de dados

3. **Erro ao gerar código com SQLc**:
   - Verifique a sintaxe SQL nos arquivos de consulta
   - Verifique se os esquemas estão corretos
   - Consulte a documentação do SQLc para mais informações

### Documentação Interna

O projeto inclui documentação detalhada na pasta `.docs/` sobre vários aspectos do sistema:

- `emails.md`: Como disparar emails usando o provedor de email
- `autenticacao.md`: Sistema de autenticação
- `handlers.md`: Como criar e usar handlers
- `jobs.md`: Trabalhando com jobs em background
- `logs.md`: Sistema de logs
- `middlewares.md`: Criação e uso de middlewares
- `provedores.md`: Integração com serviços externos
- `queries.md`: Trabalhando com consultas SQL
- `rotas.md`: Configuração de rotas

### Recursos Adicionais

- [Documentação do Go](https://golang.org/doc/)
- [Documentação do SQLc](https://docs.sqlc.dev/)
- [Documentação do Golang Migrate](https://github.com/golang-migrate/migrate)
- [Documentação do Air](https://github.com/air-verse/air)

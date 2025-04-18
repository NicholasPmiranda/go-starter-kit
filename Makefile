include .env
export

# Função para montar a URL do banco de dados
db_url = postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable

#
# COMANDOS DOCKER
#

# Inicia os containers Docker
up:
	@docker compose up -d

# Para os containers Docker
down:
	@docker compose down

# Reinicia os containers Docker
restart:
	@docker compose restart

# Constrói as imagens Docker
build:
	@docker compose build

# Exibe os logs dos containers
logs:
	@docker compose logs -f

# Executa um comando no container da aplicação
docker-exec:
	@docker compose exec app $(filter-out $@, $(MAKECMDGOALS))

# Acessa o shell do container da aplicação
shell:
	@docker compose exec app sh

# Executa a aplicação com hot reload usando Air dentro do container
docker-dev:
	@docker compose exec app air

# Executa a aplicação sem hot reload dentro do container
docker-run:
	@docker compose exec app go run ./cmd/server/main.go

# Executa o SQLc para gerar código a partir de SQL dentro do container
docker-sqlc:
	@docker compose exec app sqlc generate

#
# COMANDOS DE MIGRAÇÃO
#

# Cria uma nova migration
migration:
	@migrate create -ext sql -dir database/migrations $(filter-out $@, $(MAKECMDGOALS))

# Aplica todas as migrations
migrate:
	@migrate -path database/migrations -database "$(db_url)" up

# Dropa todo o banco e roda as migrations novamente
migrate\:fresh:
	@migrate -path database/migrations -database "$(db_url)" drop
	@migrate -path database/migrations -database "$(db_url)" up

# Dá rollback na última migration
migrate\:rollback:
	@migrate -path database/migrations -database "$(db_url)" down 1

# Cria diretório de migrations se não existir
database/migrations:
	@mkdir -p database/migrations

# Executa migrations dentro do container Docker
docker-migrate:
	@docker compose exec app migrate -path database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@pgsql:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" up

# Executa fresh migrations dentro do container Docker
docker-migrate\:fresh:
	@docker compose exec app migrate -path database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@pgsql:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" drop
	@docker compose exec app migrate -path database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@pgsql:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" up

# Executa rollback da última migration dentro do container Docker
docker-migrate\:rollback:
	@docker compose exec app migrate -path database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@pgsql:$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" down 1

#
# COMANDOS DE INSTALAÇÃO
#

# Comando para instalar o golang-migrate CLI
install-migrate:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Instala o Air para hot reload
install-air:
	@go install github.com/air-verse/air@latest

# Instala o SQLc
install-sqlc:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

#
# COMANDOS DE EXECUÇÃO LOCAL
#

# Executa a aplicação com hot reload usando Air
dev:
	@air

# Executa a aplicação sem hot reload
run:
	@go run ./cmd/server/main.go

include .env
export

# Função para montar a URL do banco de dados
db_url = postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable

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

# Comando para instalar o golang-migrate CLI
install-migrate:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Instala o Air para hot reload
install-air:
	@go install github.com/cosmtrek/air@latest

# Executa a aplicação com hot reload usando Air
dev:
	@air

# Executa a aplicação sem hot reload
run:
	@go run ./cmd/server/main.go

FROM golang:1.24-alpine

# Instalar dependências do sistema
RUN apk add --no-cache git curl make gcc libc-dev

# Instalar Air para hot-reload
RUN go install github.com/air-verse/air@latest

# Instalar SQLc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Instalar Golang Migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Configurar diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o restante dos arquivos
COPY . .

# Expor a porta da aplicação
EXPOSE 3030

# Comando padrão para executar a aplicação com hot-reload
CMD ["air"]

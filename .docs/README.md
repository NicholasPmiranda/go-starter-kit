# Documentação do Projeto Go Starter Kit

## Visão Geral

Este projeto é um kit inicial (boilerplate) para desenvolvimento de aplicações em Go, seguindo boas práticas e padrões de projeto. Ele fornece uma estrutura organizada e componentes pré-configurados para acelerar o desenvolvimento de novas aplicações.

## Estrutura do Projeto

```
/
├── cmd/                    # Pontos de entrada da aplicação
│   ├── server/             # Servidor principal
│   └── worker/             # Worker para processamento em background
├── config/                 # Configurações e provedores externos
│   ├── emailProvider/      # Provedor de serviço de e-mail
│   ├── looger/             # Configuração de logs
│   ├── queue/              # Configuração de filas
│   └── storageProvider/    # Provedor de armazenamento
├── database/               # Código de suporte para banco de dados
│   ├── migrations/         # Migrações
│   ├── query/              # Consultas SQL
│   ├── schema/             # Definições de esquema
│   └── seeds/              # Dados iniciais
├── docs/                   # Documentação do projeto
├── helpers/                # Funções auxiliares
│   └── authHelper/         # Auxiliares para autenticação
├── internal/               # Código específico da aplicação
│   ├── database/           # Conexão e modelos de banco de dados
│   ├── http/               # Componentes HTTP
│   │   ├── handler/        # Manipuladores de requisições
│   │   ├── middleware/     # Middlewares
│   │   └── request/        # Modelos de requisição
│   └── jobs/               # Definição de jobs assíncronos
├── pkg/                    # Bibliotecas reutilizáveis
├── routes/                 # Definição de rotas
├── storage/                # Armazenamento de arquivos
│   ├── app/                # Arquivos da aplicação
│   ├── log/                # Logs
│   └── planilha/           # Planilhas importadas
└── template/               # Templates HTML
```

## Componentes Principais

### Servidor HTTP

O servidor HTTP é implementado usando o framework Gin e é o ponto de entrada principal da aplicação. Ele é responsável por receber requisições HTTP, roteá-las para os handlers apropriados e retornar respostas.

### Sistema de Jobs

O sistema de jobs assíncronos é implementado usando a biblioteca Asynq. Ele permite que tarefas sejam enfileiradas para processamento em background, melhorando a responsividade da aplicação.

### Banco de Dados

O projeto utiliza PostgreSQL como banco de dados e SQLC para gerar código Go a partir de consultas SQL. Isso proporciona uma camada de acesso a dados fortemente tipada e eficiente.

### Autenticação

O sistema de autenticação é baseado em tokens JWT (JSON Web Tokens) e inclui middlewares para proteger rotas que requerem autenticação.

### Provedores Externos

O projeto inclui provedores para integração com serviços externos, como envio de e-mails e armazenamento de arquivos.

## Documentação Específica

Para informações mais detalhadas sobre como usar os diferentes componentes do projeto, consulte os seguintes documentos:

- [Como Disparar Emails](./emails.md)
- [Como Criar e Disparar Jobs](./jobs.md)
- [Como Criar Rotas](./rotas.md)
- [Como Criar Queries](./queries.md)
- [Como Usar Handlers](./handlers.md)
- [Como Usar Middlewares](./middlewares.md)
- [Como Usar Provedores](./provedores.md)
- [Como Usar o Sistema de Logs](./logs.md)
- [Sistema de Autenticação](./autenticacao.md)
- [Como Usar o Hot Reload](./hot-reload.md)
- [API de Clientes](endpoints/client/clientes.md)

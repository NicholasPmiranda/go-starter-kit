---
description: Convenções de código e práticas de desenvolvimento
globs: 
alwaysApply: true
---

sempre responda tudo em portugues e nunca  em ingles

## Convenções para Go

### 1. Estrutura de Pastas e Arquivos
- **camelCase**: Para arquivos Go, use nomes em letras minúsculas, sem underscores.
- **CamelCase**: Para nomes exportados (funções, structs, interfaces), siga a convenção CamelCase.
- **lowercase**: Para pacotes (diretórios), use nomes em minúsculas sem underscores.
- **camelCase**: Para nomes de variáveis, use camelCase iniciado com letra minúscula.


### 2. Estrutura do Projeto (baseada em golang-standards/project-layout)
```
/
├── cmd/                    # Pontos de entrada da aplicação
│   ├── server/             # Servidor principal
│   └── worker/             # Worker para processamento em background
├── internal/               # Código específico da aplicação (como o "app" do Laravel)
│   ├── domain/             # Modelos de domínio e regras de negócio
│   ├── service/            # Serviços da aplicação
│   ├── repository/         # Implementações de acesso a dados
│   ├── handler/            # Manipuladores HTTP
│   └── usecase/            # Casos de uso da aplicação
├── pkg/                    # Bibliotecas reutilizáveis
│   └── workerPool/         # Implementação de pool de workers
├── api/                    # Definições de API (swagger, protobuf)
├── configs/                # Arquivos de configuração
├── database/               # Código de suporte para banco de dados
│   ├── migrations/         # Migrações
│   ├── query/              # Consultas SQL
│   ├── schema/             # Definições de esquema
│   └── seeds/              # Dados iniciais
├── providers/              # Integrações com serviços externos
│   ├── emailProvider/      # Serviço de e-mail
│   └── storageProvider/    # Serviço de armazenamento
├── routes/                 # Definição de rotas
├── helpers/                # Funções auxiliares
├── storage/                # Armazenamento de arquivos
└── template/               # Templates
```

### 3. Nomes de Tipos e Interfaces
- **PascalCase** para tipos exportados (UserService, OrderRepository).
- **camelCase** para variáveis e funções não exportadas (calculatePrice, getUserByID).
- Evite redundância: Não use UserServiceStruct, apenas UserService.
- Interfaces com um único método geralmente são nomeadas com o sufixo "er" (Reader, Writer).

### 4. Nomes de Variáveis e Funções
- **camelCase** para variáveis e funções não exportadas (calculatePrice, getUserByID).
- **PascalCase** para variáveis e funções exportadas (ProcessPayment, SendEmail).
- Use nomes claros e descritivos, evitando abreviações excessivas (GetUserData e não gud).
- Acrônimos em nomes devem ser tratados como uma palavra (HttpServer → HTTPServer, Api → API).
- Nunca use Impl como sufixo para dizer que está implementando uma interface.

### 5. Convenções para Testes
- Arquivos de teste têm o sufixo _test.go (user_service_test.go).
- Funções de teste começam com Test seguido do nome da função testada (TestUserService_Process).
- Benchmarks começam com Benchmark (BenchmarkUserService_Process).
- Exemplos começam com Example (ExampleUserService_Process).
- Os testes ficam no mesmo pacote que o código testado.

### 6. Importações
- Organize as importações em grupos:
    1. Pacotes da biblioteca padrão
    2. Pacotes de terceiros
    3. Pacotes internos do projeto
- Use o caminho completo de importação baseado no módulo Go.

Exemplo:
```go
import (
    "context"
    "fmt"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    
    "github.com/seu-usuario/seu-projeto/internal/domain"
)
```

### 7. Tratamento de Erros
- Retorne erros explicitamente em vez de usar panics.
- Use pacotes como "errors" ou "github.com/pkg/errors" para criar e enriquecer erros.
- Verifique erros imediatamente após a chamada que pode gerá-los.
- Evite usar _ para ignorar erros, a menos que seja absolutamente necessário.

### 8. Documentação
- Todos os pacotes e funções/tipos exportados devem ter comentários de documentação.
- Comentários de documentação começam com o nome do elemento que estão documentando em inglês.
- Use frases completas com ponto final.

Exemplo:
```go
// UserService providers methods...
type UserService struct {
    // ...
}

// Process process the payment
// Returns a response
func (s *UserService) Process(ctx context.Context, req Request) (Response, error) {
    // ...
}
```


### 9. Organização da Lógica de Aplicação

- **Toda a lógica da aplicação deve estar contida dentro da pasta `internal/`** (equivalente à pasta `app/` do Laravel).
- **Pastas fora de `internal/` são destinadas a infraestrutura, configuração, documentação, integração e recursos externos.**


### 10. Banco de dados

- Postgresql.
- Sqlc


### 11. Services

 - sempre crie services  pra abtrair os comportamentos  e nao  deixa toda logica no controller
 - sempre crie  os tipos  dentro de internal/types
 - sempre  cria agrupamento pra mehor etendimento  ao inves de criar 3 services soltos como categoria post e user coloca todos dentro um /blog pra  agrupar



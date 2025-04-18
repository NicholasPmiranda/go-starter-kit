# Como Usar o Sistema de Logs

## Visão Geral

O sistema de logs do Go Starter Kit permite registrar informações, avisos e erros durante a execução da aplicação. Isso é essencial para monitoramento, depuração e auditoria. O sistema utiliza um logger configurável que pode direcionar os logs para diferentes destinos (console, arquivos, serviços externos).

## Estrutura

O sistema de logs é implementado no diretório:

```
config/looger/
└── looger.go       # Configuração e inicialização do logger
```

## Configuração

O logger é configurado no arquivo `config/looger/looger.go` e pode ser ajustado de acordo com as necessidades do projeto. Por padrão, os logs são armazenados no arquivo `storage/log/app.log`.

## Como Usar

### 1. Importar o Pacote

```go
import (
    "boilerPlate/config/looger"
)
```

### 2. Obter uma Instância do Logger

```go
// Obter o logger global
logger := looger.GetLogger()
```

### 3. Registrar Logs

```go
// Log de informação
logger.Info("Operação realizada com sucesso", "user_id", 123, "action", "login")

// Log de aviso
logger.Warn("Tentativa de acesso suspeita", "ip", "192.168.1.1", "attempts", 5)

// Log de erro
logger.Error("Falha ao processar pagamento", "error", err, "payment_id", paymentID)

// Log de depuração (apenas em ambiente de desenvolvimento)
logger.Debug("Valores intermediários", "step", "validation", "data", data)
```

### 4. Logs Estruturados

O logger suporta logs estruturados, onde você pode adicionar pares chave-valor para fornecer contexto adicional:

```go
logger.Info("Usuário criado",
    "user_id", user.ID,
    "email", user.Email,
    "role", user.Role,
    "created_at", time.Now(),
)
```

### 5. Logs com Contexto

Você pode criar um logger com contexto adicional que será incluído em todas as mensagens:

```go
// Criar um logger com contexto
requestLogger := logger.With(
    "request_id", requestID,
    "user_id", userID,
    "ip", clientIP,
)

// Usar o logger com contexto
requestLogger.Info("Requisição recebida", "path", "/api/users", "method", "POST")
requestLogger.Info("Processando dados", "step", "validation")
requestLogger.Info("Requisição concluída", "status", 200, "duration_ms", 45)
```

## Níveis de Log

O sistema de logs suporta diferentes níveis, que podem ser configurados de acordo com o ambiente:

1. **Debug**: Informações detalhadas para depuração (apenas em ambiente de desenvolvimento)
2. **Info**: Informações gerais sobre o funcionamento da aplicação
3. **Warn**: Avisos que não impedem o funcionamento, mas merecem atenção
4. **Error**: Erros que afetam uma operação específica
5. **Fatal**: Erros graves que impedem o funcionamento da aplicação (encerra a aplicação)

## Exemplo em um Handler

```go
package userhandler

import (
    "boilerPlate/config/looger"
    "boilerPlate/internal/database"
    "context"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetUser(c *gin.Context) {
    logger := looger.GetLogger()
    requestID := c.GetHeader("X-Request-ID")
    
    // Logger com contexto da requisição
    reqLogger := logger.With(
        "request_id", requestID,
        "handler", "GetUser",
    )
    
    // Log de início da requisição
    reqLogger.Info("Iniciando processamento da requisição")
    
    // Validação de entrada
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        reqLogger.Warn("ID inválido fornecido", "id_str", idStr, "error", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
        })
        return
    }
    
    // Processamento
    db := c.MustGet("db").(*database.Queries)
    user, err := db.FindById(context.Background(), id)
    if err != nil {
        reqLogger.Error("Erro ao buscar usuário", "error", err, "user_id", id)
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Usuário não encontrado",
        })
        return
    }
    
    // Log de sucesso
    reqLogger.Info("Usuário encontrado com sucesso", "user_id", user.ID)
    
    // Resposta
    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}
```

## Exemplo em um Job

```go
package jobs

import (
    "boilerPlate/config/looger"
    "boilerPlate/internal/http/request/RequestModel"
    "context"
    "encoding/json"
    "github.com/hibiken/asynq"
)

func Execute() asynq.HandlerFunc {
    logger := looger.GetLogger()
    
    return func(ctx context.Context, task *asynq.Task) error {
        // Logger com contexto do job
        jobLogger := logger.With(
            "job", "ExampleJob",
            "task_id", task.ID(),
        )
        
        jobLogger.Info("Iniciando processamento do job")
        
        var payload RequestModel.Pessoa
        if err := json.Unmarshal(task.Payload(), &payload); err != nil {
            jobLogger.Error("Erro ao deserializar payload", "error", err)
            return err
        }
        
        jobLogger.Info("Payload deserializado com sucesso", "nome", payload.Nome)
        
        // Lógica de processamento do job
        // ...
        
        jobLogger.Info("Job processado com sucesso")
        return nil
    }
}
```

## Configuração Avançada

### Configurar Nível de Log

Você pode configurar o nível de log de acordo com o ambiente:

```go
// Em ambiente de desenvolvimento
logger.SetLevel(looger.DebugLevel)

// Em ambiente de produção
logger.SetLevel(looger.InfoLevel)
```

### Configurar Destinos de Log

Você pode configurar múltiplos destinos para os logs:

```go
// Configurar log para console e arquivo
logger.SetOutput([]io.Writer{os.Stdout, fileWriter})
```

### Rotação de Arquivos de Log

Para evitar que os arquivos de log cresçam indefinidamente, você pode implementar rotação de logs:

```go
// Configurar rotação de logs
rotateWriter := &lumberjack.Logger{
    Filename:   "storage/log/app.log",
    MaxSize:    10,    // megabytes
    MaxBackups: 3,     // número de backups
    MaxAge:     28,    // dias
    Compress:   true,  // compactar backups
}

logger.SetOutput(rotateWriter)
```

## Boas Práticas

1. **Consistência**: Use o mesmo formato e estilo de logs em toda a aplicação.
2. **Contexto**: Sempre inclua contexto suficiente para entender o log.
3. **Níveis Apropriados**: Use o nível de log apropriado para cada mensagem.
4. **Informações Sensíveis**: Nunca registre informações sensíveis como senhas ou tokens.
5. **Performance**: Evite logs excessivos em código de alto desempenho.
6. **Estruturação**: Prefira logs estruturados (chave-valor) em vez de mensagens de texto simples.
7. **Identificadores**: Inclua identificadores únicos (IDs de requisição, usuário, etc.) para facilitar a correlação de logs.
8. **Rotação**: Implemente rotação de logs para gerenciar o tamanho dos arquivos.

# Documentação do Scheduler

## Introdução

O Scheduler é um mecanismo de agendamento de tarefas que permite executar jobs em intervalos de tempo específicos. Ele é integrado com a API principal e é iniciado automaticamente quando a aplicação é iniciada.

## Configuração

O Scheduler está configurado no arquivo `config/scheduler/scheduler.go`. A configuração padrão usa o Redis como backend para armazenar as tarefas agendadas.

### Configuração do Redis

```go
const redisAddr = "localhost:6379"
```

Você pode alterar o endereço do Redis conforme necessário.

## Como Registrar Tarefas

As tarefas são registradas na função `registerTasks` no arquivo `config/scheduler/scheduler.go`. Veja um exemplo:

```go
func registerTasks(ts *TaskScheduler) {
    // Exemplo de tarefa agendada para executar a cada minuto
    examplePayload := RequestModel.Pessoa{
        Nome:  "Exemplo",
        Email: "exemplo@exemplo.com",
        Idade: 30,
    }

    task, _ := jobs.NewJobModel(examplePayload)
    ts.Register(task).EveryMinute()
}
```

### Criando um Novo Job

Para criar um novo job, você precisa:

1. Criar uma função que retorne um `*asynq.Task` com o payload necessário
2. Criar uma função que execute o job quando for agendado
3. Registrar o job no scheduler

Exemplo:

```go
// No arquivo internal/jobs/meuJob.go
package jobs

import (
    "context"
    "encoding/json"
    "github.com/hibiken/asynq"
    "log"
)

const MeuJobName = "meuJob"

// Payload do job
type MeuJobPayload struct {
    ID   string `json:"id"`
    Nome string `json:"nome"`
}

// Função para criar um novo job
func NewMeuJob(payload MeuJobPayload) (*asynq.Task, error) {
    data, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    return asynq.NewTask(MeuJobName, data), nil
}

// Função para executar o job
func ExecuteMeuJob() asynq.HandlerFunc {
    return func(ctx context.Context, task *asynq.Task) error {
        var payload MeuJobPayload
        if err := json.Unmarshal(task.Payload(), &payload); err != nil {
            return err
        }
        
        log.Printf("Executando job para %s (ID: %s)", payload.Nome, payload.ID)
        
        // Lógica do job aqui
        
        return nil
    }
}
```

### Registrando o Job no Worker

Você também precisa registrar o handler do job no worker para que ele possa ser executado. Isso é feito no arquivo `config/worker/worker.go`:

```go
// No arquivo config/worker/worker.go
mux.HandleFunc(jobs.MeuJobName, jobs.ExecuteMeuJob())
```

### Registrando o Job no Scheduler

Finalmente, você precisa registrar o job no scheduler para que ele seja agendado:

```go
// No arquivo config/scheduler/scheduler.go
func registerTasks(ts *TaskScheduler) {
    // ... outros jobs ...
    
    // Meu job personalizado
    meuJobPayload := jobs.MeuJobPayload{
        ID:   "123",
        Nome: "Tarefa Importante",
    }
    
    meuJob, _ := jobs.NewMeuJob(meuJobPayload)
    ts.Register(meuJob).Daily() // Executa diariamente à meia-noite
}
```

## Padrões de Tempo Disponíveis

O scheduler suporta os seguintes padrões de tempo:

### Intervalos de Tempo

- **EveryMinute()** — Executa a tarefa a cada minuto.
- **EveryTwoMinutes()** — A cada 2 minutos.
- **EveryFiveMinutes()** — A cada 5 minutos.
- **EveryTenMinutes()** — A cada 10 minutos.
- **EveryFifteenMinutes()** — A cada 15 minutos.
- **EveryThirtyMinutes()** — A cada 30 minutos.
- **Hourly()** — A cada hora.
- **HourlyAt(17)** — A cada hora, no minuto 17.
- **Daily()** — Todos os dias à meia-noite.
- **DailyAt("13:00")** — Todos os dias às 13h.
- **TwiceDaily(1, 13)** — Duas vezes por dia, às 1h e 13h.
- **Weekly()** — Toda semana.
- **WeeklyOn(1, "8:00")** — Toda segunda-feira às 8h.
- **Monthly()** — Todo mês.
- **MonthlyOn(4, "15:00")** — Todo dia 4 do mês às 15h.
- **Quarterly()** — A cada trimestre.
- **Yearly()** — Todo ano.

### Restrições de Dias

- **Weekdays()** — Executa apenas nos dias úteis (segunda a sexta).
- **Weekends()** — Executa apenas nos finais de semana (sábado e domingo).
- **Mondays()** — Executa apenas nas segundas-feiras.
- **Tuesdays()** — Executa apenas nas terças-feiras.
- **Wednesdays()** — Executa apenas nas quartas-feiras.
- **Thursdays()** — Executa apenas nas quintas-feiras.
- **Fridays()** — Executa apenas nas sextas-feiras.
- **Saturdays()** — Executa apenas nos sábados.
- **Sundays()** — Executa apenas nos domingos.

## Exemplos de Uso

### Executar uma tarefa diariamente às 8h

```go
ts.Register(task).DailyAt("8:00")
```

### Executar uma tarefa toda segunda-feira às 9h

```go
ts.Register(task).WeeklyOn(1, "9:00") // 1 = segunda-feira (0 = domingo, 6 = sábado)
```

### Executar uma tarefa a cada 15 minutos

```go
ts.Register(task).EveryFifteenMinutes()
```

### Executar uma tarefa no primeiro dia de cada mês

```go
ts.Register(task).MonthlyOn(1, "00:00")
```

### Executar uma tarefa apenas nos dias úteis

```go
ts.Register(task).Weekdays()
```

## Considerações Finais

- O scheduler usa o Redis como backend, então certifique-se de que o Redis está em execução.
- As tarefas são executadas pelo worker, então certifique-se de que o worker está em execução.
- Se você precisar de padrões de tempo mais complexos, você pode usar expressões cron diretamente.

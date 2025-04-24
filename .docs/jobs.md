# Como Criar e Disparar Jobs

## Visão Geral

O sistema de jobs do Go Starter Kit utiliza a biblioteca [Asynq](https://github.com/hibiken/asynq) para processamento assíncrono de tarefas. Isso permite que operações demoradas sejam executadas em background, melhorando a responsividade da aplicação.

## Estrutura de um Job

Um job no Go Starter Kit consiste em dois componentes principais:

1. **Definição do Job**: Define o nome do job e como criar uma nova tarefa.
2. **Executor do Job**: Define a função que será executada quando o job for processado.

### Exemplo de Definição de Job

```go
package jobs

import (
	"sixTask/internal/http/request/RequestModel"
	"encoding/json"
	"github.com/hibiken/asynq"
)

// Nome único do job
const JobName = "jobModelo"

// Função para criar uma nova tarefa
func NewJobModel(request RequestModel.Pessoa) (*asynq.Task, error) {
	// Serializa o payload para JSON
	payload, _ := json.Marshal(&request)

	// Cria e retorna uma nova tarefa com o nome do job e o payload
	return asynq.NewTask(JobName, payload), nil
}

// Função que será executada quando o job for processado
func Execute() asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		// Deserializa o payload
		var payload RequestModel.Pessoa
		json.Unmarshal(task.Payload(), &payload)

		// Lógica de processamento do job
		// ...

		return nil
	}
}
```

## Como Registrar um Job

Para que o worker possa processar o job, é necessário registrá-lo no servidor Asynq. Isso é feito no arquivo `cmd/worker/main.go`:

```go
func main() {
	// ...

	// Cria um novo servidor Asynq
	srv := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"default": 10,
		},
	})

	// Registra os handlers para cada tipo de job
	mux := asynq.NewServeMux()
	mux.HandleFunc(jobs.JobName, jobs.Execute())

	// Inicia o servidor
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
```

## Como Disparar um Job

Existem duas maneiras de disparar um job:

### 1. Através de um Handler HTTP

```go
package JobHandler

import (
	"sixTask/config/queue"
	"sixTask/internal/http/request/RequestModel"
	"sixTask/internal/jobs"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func DisparJob(c *gin.Context) {
	// Recebe os dados da requisição
	var request RequestModel.Pessoa
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Cria uma nova tarefa
	task, err := jobs.NewJobModel(request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Conecta ao cliente de fila
	queueCliente := queue.Conect()
	defer queueCliente.Close()

	// Enfileira a tarefa
	queueCliente.Enqueue(task, asynq.Queue("default"))

	c.JSON(200, gin.H{
		"message": "processamento concluido",
	})
}
```

### 2. Diretamente no Código

```go
package exemplo

import (
	"sixTask/config/queue"
	"sixTask/internal/http/request/RequestModel"
	"sixTask/internal/jobs"
	"github.com/hibiken/asynq"
)

func EnfileirarJob() {
	// Cria os dados para o job
	pessoa := RequestModel.Pessoa{
		Nome: "João",
		Idade: 30,
	}

	// Cria uma nova tarefa
	task, _ := jobs.NewJobModel(pessoa)

	// Conecta ao cliente de fila
	queueCliente := queue.Conect()
	defer queueCliente.Close()

	// Enfileira a tarefa
	queueCliente.Enqueue(task, asynq.Queue("default"))
}
```

## Monitoramento de Jobs

O Go Starter Kit inclui uma interface web para monitoramento de jobs, acessível através da rota `/monitor`. Esta interface permite:

- Visualizar jobs em execução
- Visualizar jobs enfileirados
- Visualizar jobs concluídos
- Visualizar jobs com erro
- Reprocessar jobs com erro

## Boas Práticas

1. **Nomeação**: Use nomes descritivos para seus jobs, preferencialmente com um sufixo "Job" (ex: `ImportacaoJob`).
2. **Idempotência**: Sempre que possível, implemente jobs idempotentes (que podem ser executados múltiplas vezes sem efeitos colaterais).
3. **Tratamento de Erros**: Implemente tratamento adequado de erros e considere estratégias de retry para jobs que podem falhar temporariamente.
4. **Logging**: Adicione logs adequados para facilitar o diagnóstico de problemas.
5. **Timeout**: Configure timeouts apropriados para evitar que jobs fiquem executando indefinidamente.

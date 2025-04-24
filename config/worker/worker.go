package worker

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sixTask/
	"syscall"
	"time"

	"github.com/hibiken/asynq"
)

// Configuração do Redis para o worker
const redisAddr = "localhost:6379"

// SetupWorker configura e inicia o worker
func SetupWorker() {
	// Cria um novo cliente Asynq
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// Cria um novo servidor Asynq
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"planilhas": 6,
				"default":   2,
			},
		},
	)

	// Cria um novo multiplexador para registrar os handlers
	mux := asynq.NewServeMux()

	// Registra o handler para o job de exemplo
	mux.HandleFunc(jobs.JobName, jobs.Execute())

	// Configura o canal para capturar sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Inicia o servidor em uma goroutine
	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("Erro no worker: %v", err)
		}
	}()

	// Aguarda o sinal de interrupção em uma goroutine separada
	go func() {
		<-quit
		fmt.Println("Finalizando worker...")

		// Cria um canal para sinalizar que o worker foi encerrado
		done := make(chan struct{})

		// Inicia o encerramento do worker em uma goroutine
		go func() {
			srv.Shutdown()
			close(done)
		}()

		// Espera o encerramento do worker ou o timeout
		select {
		case <-done:
			log.Println("Worker encerrado com sucesso")
		case <-time.After(5 * time.Second):
			log.Println("Timeout ao encerrar o worker, forçando encerramento")
			// Aqui poderíamos usar os.Exit(1) para forçar o encerramento,
			// mas isso pode causar perda de dados. Em vez disso, apenas logamos o evento.
		}
	}()

	log.Println("Worker iniciado com sucesso")
}

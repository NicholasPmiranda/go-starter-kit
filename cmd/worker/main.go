package main

import (
	logger "boilerPlate/config/looger"
	"boilerPlate/internal/jobs"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	logger.SetupLogger()

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"planilhas": 6,
				"default":   2,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(jobs.JobName, jobs.Execute())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Roda o servidor em segundo plano
	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("Erro no worker: %v", err)
		}
	}()

	log.Println("Worker iniciado. Pressione Ctrl+C para sair.")

	// Espera Ctrl+C
	<-quit

	log.Println("Finalizando aplicação...")

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
	}
}

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
}

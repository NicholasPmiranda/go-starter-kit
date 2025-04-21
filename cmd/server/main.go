package main

import (
	logger "boilerPlate/config/looger"
	schedulerConfig "boilerPlate/config/scheduler"
	workerConfig "boilerPlate/config/worker"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"boilerPlate/database/seeds"
	"boilerPlate/routes"
)

func main() {
	logger.SetupLogger()

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	seeds.Run()

	// Inicializa o scheduler
	schedulerConfig.SetupScheduler()

	// Inicializa o worker
	workerConfig.SetupWorker()

	router := routes.SetupRoutes()

	// Cria um servidor HTTP com configurações personalizadas
	srv := &http.Server{
		Addr:    ":3030",
		Handler: router,
	}

	// Inicia o servidor em uma goroutine separada
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	// Configura o canal para capturar sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Aguarda o sinal de interrupção
	<-quit
	log.Println("Encerrando servidor...")

	// Cria um contexto com timeout para o encerramento
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tenta encerrar o servidor HTTP com o contexto de timeout
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao encerrar o servidor: %v", err)
	}

	log.Println("Servidor encerrado com sucesso")
}

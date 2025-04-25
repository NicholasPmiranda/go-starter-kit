package scheduler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Configuração do Redis para o scheduler
const redisAddr = "localhost:6379"

// SetupScheduler configura e inicia o scheduler
func SetupScheduler() {
	// Cria uma nova instância do scheduler
	taskScheduler := NewScheduler(redisAddr)

	// Registra as tarefas agendadas
	registerTasks(taskScheduler)

	// Inicia o scheduler em uma goroutine
	go func() {
		if err := taskScheduler.Start(); err != nil {
			log.Fatalf("Erro ao iniciar o scheduler: %v", err)
		}
	}()

	// Configura o canal para capturar sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Aguarda o sinal de interrupção em uma goroutine separada
	go func() {
		<-quit
		log.Println("Finalizando scheduler...")

		// Cria um canal para sinalizar que o scheduler foi encerrado
		done := make(chan struct{})

		// Inicia o encerramento do scheduler em uma goroutine
		go func() {
			taskScheduler.Stop()
			close(done)
		}()

		// Espera o encerramento do scheduler ou o timeout
		select {
		case <-done:
			log.Println("Scheduler encerrado com sucesso")
		case <-time.After(5 * time.Second):
			log.Println("Timeout ao encerrar o scheduler, forçando encerramento")
			// Aqui poderíamos usar os.Exit(1) para forçar o encerramento,
			// mas isso pode causar perda de dados. Em vez disso, apenas logamos o evento.
		}
	}()

	log.Println("Scheduler iniciado com sucesso")
}

// registerTasks registra todas as tarefas agendadas
func registerTasks(ts *TaskScheduler) {
	//// Exemplo de tarefa agendada para executar a cada minuto
	//examplePayload := RequestModel.Pessoa{
	//	Nome:  "Exemplo",
	//	Email: "exemplo@exemplo.com",
	//	Idade: 30,
	//}
	//
	//task, _ := jobs.NewJobModel(examplePayload)
	//ts.Register(task).EveryThirtyMinutes()

	// Aqui você pode registrar outras tarefas com diferentes intervalos
	// Exemplos:
	// ts.Register(task2).EveryFiveMinutes()
	// ts.Register(task3).Daily()
	// ts.Register(task4).WeeklyOn(1, "8:00") // Segunda-feira às 8h
	// ts.Register(task5).MonthlyOn(1, "00:00") // Dia 1 à meia-noite
}

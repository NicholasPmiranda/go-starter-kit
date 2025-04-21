package scheduler

import (
	"fmt"

	"github.com/hibiken/asynq"
)

// Schedule representa uma configuração de agendamento para uma tarefa
type Schedule struct {
	task     *asynq.Task
	cronSpec string
	options  []asynq.Option
	ts       *TaskScheduler
}

// TaskScheduler é a interface para agendar tarefas
type TaskScheduler struct {
	client    *asynq.Client
	scheduler *asynq.Scheduler
}

// NewScheduler cria uma nova instância do TaskScheduler
func NewScheduler(redisAddr string) *TaskScheduler {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: redisAddr},
		&asynq.SchedulerOpts{},
	)

	return &TaskScheduler{
		client:    client,
		scheduler: scheduler,
	}
}

// Register registra uma tarefa para ser agendada
func (ts *TaskScheduler) Register(task *asynq.Task, opts ...asynq.Option) *Schedule {
	return &Schedule{
		task:    task,
		options: opts,
		ts:      ts,
	}
}

// Start inicia o scheduler
func (ts *TaskScheduler) Start() error {
	return ts.scheduler.Start()
}

// Stop para o scheduler
func (ts *TaskScheduler) Stop() {
	ts.scheduler.Shutdown()
	ts.client.Close()
}

// EveryMinute agenda a tarefa para executar a cada minuto
func (s *Schedule) EveryMinute() *Schedule {
	s.cronSpec = "@every 1m"
	s.register()
	return s
}

// EveryTwoMinutes agenda a tarefa para executar a cada 2 minutos
func (s *Schedule) EveryTwoMinutes() *Schedule {
	s.cronSpec = "@every 2m"
	s.register()
	return s
}

// EveryFiveMinutes agenda a tarefa para executar a cada 5 minutos
func (s *Schedule) EveryFiveMinutes() *Schedule {
	s.cronSpec = "@every 5m"
	s.register()
	return s
}

// EveryTenMinutes agenda a tarefa para executar a cada 10 minutos
func (s *Schedule) EveryTenMinutes() *Schedule {
	s.cronSpec = "@every 10m"
	s.register()
	return s
}

// EveryFifteenMinutes agenda a tarefa para executar a cada 15 minutos
func (s *Schedule) EveryFifteenMinutes() *Schedule {
	s.cronSpec = "@every 15m"
	s.register()
	return s
}

// EveryThirtyMinutes agenda a tarefa para executar a cada 30 minutos
func (s *Schedule) EveryThirtyMinutes() *Schedule {
	s.cronSpec = "@every 30m"
	s.register()
	return s
}

// Hourly agenda a tarefa para executar a cada hora
func (s *Schedule) Hourly() *Schedule {
	s.cronSpec = "@hourly"
	s.register()
	return s
}

// HourlyAt agenda a tarefa para executar a cada hora no minuto especificado
func (s *Schedule) HourlyAt(minute int) *Schedule {
	if minute < 0 || minute > 59 {
		return s
	}
	s.cronSpec = fmt.Sprintf("%d * * * *", minute)
	s.register()
	return s
}

// Daily agenda a tarefa para executar todos os dias à meia-noite
func (s *Schedule) Daily() *Schedule {
	s.cronSpec = "@daily"
	s.register()
	return s
}

// DailyAt agenda a tarefa para executar todos os dias no horário especificado (formato "HH:MM")
func (s *Schedule) DailyAt(time string) *Schedule {
	var hour, minute int
	_, err := fmt.Sscanf(time, "%d:%d", &hour, &minute)
	if err != nil || hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return s
	}
	s.cronSpec = fmt.Sprintf("%d %d * * *", minute, hour)
	s.register()
	return s
}

// TwiceDaily agenda a tarefa para executar duas vezes por dia nos horários especificados
func (s *Schedule) TwiceDaily(hour1, hour2 int) *Schedule {
	if hour1 < 0 || hour1 > 23 || hour2 < 0 || hour2 > 23 {
		return s
	}
	s.cronSpec = fmt.Sprintf("0 %d,%d * * *", hour1, hour2)
	s.register()
	return s
}

// Weekly agenda a tarefa para executar uma vez por semana (domingo à meia-noite)
func (s *Schedule) Weekly() *Schedule {
	s.cronSpec = "@weekly"
	s.register()
	return s
}

// WeeklyOn agenda a tarefa para executar no dia da semana e horário especificados
// day: 0 (domingo) a 6 (sábado)
func (s *Schedule) WeeklyOn(day int, time string) *Schedule {
	if day < 0 || day > 6 {
		return s
	}

	var hour, minute int
	_, err := fmt.Sscanf(time, "%d:%d", &hour, &minute)
	if err != nil || hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return s
	}

	s.cronSpec = fmt.Sprintf("%d %d * * %d", minute, hour, day)
	s.register()
	return s
}

// Monthly agenda a tarefa para executar uma vez por mês (dia 1 à meia-noite)
func (s *Schedule) Monthly() *Schedule {
	s.cronSpec = "@monthly"
	s.register()
	return s
}

// MonthlyOn agenda a tarefa para executar no dia do mês e horário especificados
func (s *Schedule) MonthlyOn(day int, time string) *Schedule {
	if day < 1 || day > 31 {
		return s
	}

	var hour, minute int
	_, err := fmt.Sscanf(time, "%d:%d", &hour, &minute)
	if err != nil || hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return s
	}

	s.cronSpec = fmt.Sprintf("%d %d %d * *", minute, hour, day)
	s.register()
	return s
}

// Quarterly agenda a tarefa para executar uma vez por trimestre
func (s *Schedule) Quarterly() *Schedule {
	s.cronSpec = "0 0 1 1,4,7,10 *"
	s.register()
	return s
}

// Yearly agenda a tarefa para executar uma vez por ano (1º de janeiro à meia-noite)
func (s *Schedule) Yearly() *Schedule {
	s.cronSpec = "@yearly"
	s.register()
	return s
}

// Weekdays agenda a tarefa para executar nos dias úteis (segunda a sexta)
func (s *Schedule) Weekdays() *Schedule {
	s.cronSpec = "0 0 * * 1-5"
	s.register()
	return s
}

// Weekends agenda a tarefa para executar nos finais de semana (sábado e domingo)
func (s *Schedule) Weekends() *Schedule {
	s.cronSpec = "0 0 * * 0,6"
	s.register()
	return s
}

// Mondays agenda a tarefa para executar todas as segundas-feiras
func (s *Schedule) Mondays() *Schedule {
	s.cronSpec = "0 0 * * 1"
	s.register()
	return s
}

// Tuesdays agenda a tarefa para executar todas as terças-feiras
func (s *Schedule) Tuesdays() *Schedule {
	s.cronSpec = "0 0 * * 2"
	s.register()
	return s
}

// Wednesdays agenda a tarefa para executar todas as quartas-feiras
func (s *Schedule) Wednesdays() *Schedule {
	s.cronSpec = "0 0 * * 3"
	s.register()
	return s
}

// Thursdays agenda a tarefa para executar todas as quintas-feiras
func (s *Schedule) Thursdays() *Schedule {
	s.cronSpec = "0 0 * * 4"
	s.register()
	return s
}

// Fridays agenda a tarefa para executar todas as sextas-feiras
func (s *Schedule) Fridays() *Schedule {
	s.cronSpec = "0 0 * * 5"
	s.register()
	return s
}

// Saturdays agenda a tarefa para executar todos os sábados
func (s *Schedule) Saturdays() *Schedule {
	s.cronSpec = "0 0 * * 6"
	s.register()
	return s
}

// Sundays agenda a tarefa para executar todos os domingos
func (s *Schedule) Sundays() *Schedule {
	s.cronSpec = "0 0 * * 0"
	s.register()
	return s
}

// register registra a tarefa no scheduler
func (s *Schedule) register() {
	if s.cronSpec != "" && s.ts != nil && s.task != nil {
		// Ignora o retorno do método Register
		s.ts.scheduler.Register(s.cronSpec, s.task, s.options...)
	}
}

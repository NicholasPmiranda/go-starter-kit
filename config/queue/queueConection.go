package queue

import "github.com/hibiken/asynq"

func Conect() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
}

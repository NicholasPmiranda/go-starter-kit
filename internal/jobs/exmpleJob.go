package jobs

import (
	"boilerPlate/internal/http/request/RequestModel"
	"fmt"

	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

const JobName = "jobModelo"

func NewJobModel(request RequestModel.Pessoa) (*asynq.Task, error) {

	payload, _ := json.Marshal(&request)

	return asynq.NewTask(JobName, payload), nil
}

func Execute() asynq.HandlerFunc {

	return func(ctx context.Context, task *asynq.Task) error {

		log.Println("processou")
		var payload RequestModel.Pessoa

		json.Unmarshal(task.Payload(), &payload)

		fmt.Println(payload)
		return nil
	}
}

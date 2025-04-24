package JobHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"sixTask/con
	"sixTask/config/queue"
	"sixTask/internal/http/request/RequestModel"
	"github.com/hibiken/asynq"
)

func DisparJob(c *gin.Context) {

	var request RequestModel.Pessoa

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	task, err := jobs.NewJobModel(request)
	queueCliente := queue.Conect()

	defer queueCliente.Close()

	if err != nil {
		c.JSON(400, gin.H{
			"meessage": err.Error(),
		})
	}

	queueCliente.Enqueue(task, asynq.Queue("default"))

	c.JSON(200, gin.H{
		"message": "processamento concluido",
	})

}

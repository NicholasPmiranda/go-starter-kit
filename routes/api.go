package routes

import (
	"boilerPlate/internal/http/handler"
	"boilerPlate/internal/http/handler/JobHandler"
	authhandler "boilerPlate/internal/http/handler/authHandler"
	filehandler "boilerPlate/internal/http/handler/fileHandler"
	authmiddleware "boilerPlate/internal/middleware/authMiddleware"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"*"})

	redisOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	monitor := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitor",
		RedisConnOpt: redisOpt,
	})

	router.Any("/monitor/*path", gin.WrapH(monitor))

	// Storage file serving route
	router.GET("/storage/*filepath", filehandler.GetFileHandler())

	api := router.Group("/api")
	{
		api.POST("disparar-job", JobHandler.DisparJob)

		api.POST("/", handler.StartHandler)
		api.POST("/login", authhandler.Login)
		api.POST("/upload", filehandler.UploadFileExample)

		authenticated := api.Group("/")
		authenticated.Use(authmiddleware.AuthMiddleware())
		{
			authenticated.GET("/profile", authhandler.Profile)
		}
	}

	return router
}

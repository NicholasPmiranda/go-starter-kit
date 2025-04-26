package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"sixTask/internal/http/handler"
	"sixTask/internal/http/handler/JobHandler"
	attachmenthandler "sixTask/internal/http/handler/attachmentHandler"
	authhandler "sixTask/internal/http/handler/authHandler"
	clienthandler "sixTask/internal/http/handler/clientHandler"
	commenthandler "sixTask/internal/http/handler/commentHandler"
	filehandler "sixTask/internal/http/handler/fileHandler"
	notificationhandler "sixTask/internal/http/handler/notificationHandler"
	projecthandler "sixTask/internal/http/handler/projectHandler"
	subtaskhandler "sixTask/internal/http/handler/subtaskHandler"
	taskhandler "sixTask/internal/http/handler/taskHandler"
	userhandler "sixTask/internal/http/handler/userHandler"
	"sixTask/internal/http/validator"
	
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Inicializa o validador com traduções em português
	validator.InitValidator()

	// Configurar o tamanho máximo de upload para 1GB
	router.MaxMultipartMemory = 1 << 30 // 1GB

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
		// Rotas públicas
		api.POST("disparar-job", JobHandler.DisparJob)
		api.POST("/", handler.StartHandler)
		api.POST("/login", authhandler.Login)
		api.POST("/upload", filehandler.UploadFileExample)

		// Rotas de usuário
		api.GET("/users", userhandler.GetUsers)
		api.GET("/users/:id", userhandler.GetUser)
		api.POST("/users", userhandler.CreateUser)
		api.PUT("/users/:id", userhandler.UpdateUser)
		api.DELETE("/users/:id", userhandler.DeleteUser)

		// Rotas autenticadas
		authenticated := api.Group("/")
		authenticated.Use(authmiddleware.AuthMiddleware())
		{
			authenticated.GET("/profile", authhandler.Profile)

			// Rotas de cliente
			authenticated.GET("/clients", clienthandler.GetClients)
			authenticated.GET("/clients/:id", clienthandler.GetClient)
			authenticated.POST("/clients", clienthandler.CreateClient)
			authenticated.PUT("/clients/:id", clienthandler.UpdateClient)
			authenticated.DELETE("/clients/:id", clienthandler.DeleteClient)

			// Rotas de projeto
			authenticated.GET("/projects", projecthandler.GetProjects)
			authenticated.GET("/projects/:id", projecthandler.GetProject)
			authenticated.GET("/projects/by-client/:client_id", projecthandler.GetProjectsByClient)
			authenticated.GET("/projects/by-user/:user_id", projecthandler.GetProjectsByUser)
			authenticated.POST("/projects", projecthandler.CreateProject)
			authenticated.PUT("/projects/:id", projecthandler.UpdateProject)
			authenticated.DELETE("/projects/:id", projecthandler.DeleteProject)

			// Rotas de tarefa
			authenticated.GET("/tasks", taskhandler.GetTasks)
			authenticated.GET("/tasks/:id", taskhandler.GetTask)
			authenticated.GET("/tasks/by-project/:project_id", taskhandler.GetTasksByProject)
			authenticated.GET("/tasks/by-user/:user_id", taskhandler.GetTasksByAssignedTo)
			authenticated.GET("/tasks/by-status/:status", taskhandler.GetTasksByStatus)
			authenticated.GET("/tasks/by-priority/:priority", taskhandler.GetTasksByPriority)
			authenticated.POST("/tasks", taskhandler.CreateTask)
			authenticated.PUT("/tasks/:id", taskhandler.UpdateTask)
			authenticated.PUT("/tasks/:id/complete", taskhandler.CompleteTask)
			authenticated.DELETE("/tasks/:id", taskhandler.DeleteTask)

			// Rotas de subtarefa
			authenticated.GET("/subtasks", subtaskhandler.GetSubtasks)
			authenticated.GET("/subtasks/:id", subtaskhandler.GetSubtask)
			authenticated.GET("/subtasks/by-task/:task_id", subtaskhandler.GetSubtasksByTask)
			authenticated.GET("/subtasks/by-user/:user_id", subtaskhandler.GetSubtasksByAssignedTo)
			authenticated.GET("/subtasks/by-status/:status", subtaskhandler.GetSubtasksByStatus)
			authenticated.POST("/subtasks", subtaskhandler.CreateSubtask)
			authenticated.PUT("/subtasks/:id", subtaskhandler.UpdateSubtask)
			authenticated.PUT("/subtasks/:id/complete", subtaskhandler.CompleteSubtask)
			authenticated.DELETE("/subtasks/:id", subtaskhandler.DeleteSubtask)

			// Rotas de comentário
			authenticated.GET("/comments", commenthandler.GetComments)
			authenticated.GET("/comments/:id", commenthandler.GetComment)
			authenticated.GET("/comments/user/:user_id", commenthandler.GetCommentsByUser)
			authenticated.GET("/comments/by-commentable/:commentable_type/:commentable_id", commenthandler.GetCommentsByCommentable)
			authenticated.POST("/comments", commenthandler.CreateComment)
			authenticated.PUT("/comments/:id", commenthandler.UpdateComment)
			authenticated.DELETE("/comments/:id", commenthandler.DeleteComment)

			// Rotas de anexo
			authenticated.GET("/attachments", attachmenthandler.GetAttachments)
			authenticated.GET("/attachments/:id", attachmenthandler.GetAttachment)
			authenticated.GET("/attachments/user/:user_id", attachmenthandler.GetAttachmentsByUser)
			authenticated.GET("/attachments/by-attachable/:attachable_type/:attachable_id", attachmenthandler.GetAttachmentsByAttachable)
			authenticated.POST("/attachments", attachmenthandler.CreateAttachment)
			authenticated.PUT("/attachments/:id", attachmenthandler.UpdateAttachment)
			authenticated.DELETE("/attachments/:id", attachmenthandler.DeleteAttachment)

			// Rotas de notificação
			authenticated.GET("/notifications", notificationhandler.GetNotifications)
			authenticated.GET("/notifications/:id", notificationhandler.GetNotification)
			authenticated.GET("/notifications/user/:user_id", notificationhandler.GetNotificationsByUser)
			authenticated.GET("/notifications/user/:user_id/unread", notificationhandler.GetUnreadNotificationsByUser)
			authenticated.GET("/notifications/by-notifiable/:notifiable_type/:notifiable_id", notificationhandler.GetNotificationsByNotifiable)
			authenticated.POST("/notifications", notificationhandler.CreateNotification)
			authenticated.PUT("/notifications/:id", notificationhandler.UpdateNotification)
			authenticated.PUT("/notifications/:id/read", notificationhandler.MarkNotificationAsRead)
			authenticated.PUT("/notifications/user/:user_id/read-all", notificationhandler.MarkAllNotificationsAsRead)
			authenticated.DELETE("/notifications/:id", notificationhandler.DeleteNotification)
		}
	}

	return router
}

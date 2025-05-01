package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"toDoList/internal/db"
	"toDoList/internal/handlers"
	"toDoList/internal/repo"
)

func TodoRoutes(router *gin.Engine) {

	psql := db.Psql{}
	psqlHandler := &handlers.TodoHandlers{Repo: &repo.PsqlRepo{DB: psql.Connect()}}

	mongo := db.Mng{}
	mongoHandler := &handlers.TodoHandlers{Repo: &repo.MongoRepo{DB: mongo.Connect()}}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todos := router.Group("/todos")
	{
		todos.POST("", psqlHandler.CreateHandler)
		todos.GET("", psqlHandler.GetTasksHandler)
		todos.GET("/get/:id", psqlHandler.GetTaskByIDHandler)
		todos.PUT("/update", psqlHandler.UpdateTaskHandler)
		todos.DELETE("/delete", psqlHandler.DeleteTaskHandler)
	}

	mongoTodos := router.Group("/mongoTodos")
	{
		mongoTodos.POST("", mongoHandler.CreateHandler)
		mongoTodos.GET("", mongoHandler.GetTasksHandler)
		mongoTodos.GET("/get/:id", mongoHandler.GetTaskByIDHandler)
		mongoTodos.PUT("/update", mongoHandler.UpdateTaskHandler)
		mongoTodos.DELETE("/delete", mongoHandler.DeleteTaskHandler)
	}

}

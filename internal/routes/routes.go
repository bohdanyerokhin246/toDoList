package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"toDoList/internal/handlers"
)

func TodoRoutes(db *sql.DB, router *gin.Engine) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todos := router.Group("/todos")
	{
		todos.POST("", handlers.CreateTodoHandler(db))
		todos.GET("", handlers.GetTodosHandler(db))
		todos.GET("/get/:id", handlers.GetTodoByIDHandler(db))
		todos.PUT("/update", handlers.UpdateTodoHandler(db))
		todos.DELETE("/delete", handlers.DeleteTodoHandler(db))
	}
}

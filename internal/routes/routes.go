package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"toDoList/internal/handlers"
)

func TodoRoutes(db *sql.DB, router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.POST("", handlers.CreateTodoHandler(db))
		todos.GET("/all", handlers.GetAllTodosHandler(db))
		todos.GET("/get/:id", handlers.GetTodoByIDHandler(db))
		todos.PUT("/update", handlers.UpdateTodoHandler(db))
		todos.DELETE("/delete", handlers.DeleteTodoHandler(db))
	}
}

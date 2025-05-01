package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"toDoList/internal/routes"

	_ "toDoList/docs"
)

// Run
// @title Todo List API
// @version 1.0
// @description A simple API for managing todo items
// @host localhost:8080
// @BasePath /
func Run() {
	r := gin.Default()

	routes.TodoRoutes(r)

	log.Println("Starting Todo List Service on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

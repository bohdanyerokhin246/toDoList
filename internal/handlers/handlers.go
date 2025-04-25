package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toDoList/internal/models"
	"toDoList/internal/repo"
)

// CreateTodoHandler godoc
// @Summary Create task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo for creating"
// @Success 201 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]
func CreateTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		todo, err := repo.CreateTodo(db, todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error inserting todo to database. Error: %v", err)})
			return
		}
		c.JSON(http.StatusCreated, todo)
	}
}

// GetAllTodosHandler godoc
// @Summary Get all tasks
// @Tags todos
// @Produce json
// @Success 200 {object} []models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/all [get]
func GetAllTodosHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := repo.GetAllTodos(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading todos from database. Error: %v", err)})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}

// GetTodoByIDHandler godoc
// @Summary Get task by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID of the Todo task" // ID Transmitted as part of the path
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/get/{id} [get]
func GetTodoByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		todo, err := repo.GetTodoByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error getting todo by ID. Error: %v", err)})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateTodoHandler godoc
// @Summary Update task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo for updating"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/update [put]
func UpdateTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		updatedTodo, err := repo.UpdateTodo(db, todo)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error updatinh todo. Error: %v", err)})
			return
		}
		c.JSON(http.StatusOK, updatedTodo)
	}
}

// DeleteTodoHandler godoc
// @Summary Delete task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo for deleting"
// @Success 204 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/delete [delete]
func DeleteTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := repo.DeleteTodo(db, todo); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error deleting todo. Error: %v", err)})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

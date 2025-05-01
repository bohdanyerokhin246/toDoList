package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toDoList/internal/models"
)

// CreateHandler godoc
// @Summary Create task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.CreateTodoInput true "Todo for creating"
// @Success 201 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]
func (th *TodoHandlers) CreateHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo, err := th.Repo.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error inserting todo to database. Error: %v", err)})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// GetTasksHandler godoc
// @Summary Get all tasks with filtering, sorting and pagination
// @Tags todos
// @Produce json
// @Param status query string false "Filter by status (optional)"
// @Param sortBy query string false "Sort by field: id, description, status, created_at, updated_at (default: created_at)"
// @Param sortOrder query string false "Sort order: ASC or DESC (default: ASC)"
// @Param limit query int false "Limit number of results (default: 10)"
// @Param offset query int false "Offset for pagination (default: 0)"
// @Success 200 {array} models.Todo
// @Failure 500 {object} map[string]string
// @Router /todos [get]
func (th *TodoHandlers) GetTasksHandler(c *gin.Context) {
	// Извлекаем параметры из запроса
	filter := models.TodoFilter{
		Status: c.DefaultQuery("status", ""),
	}

	// Sorting parameters
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "ASC")

	// Pagination parameters
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	todos, err := th.Repo.GetAll(filter, sortBy, sortOrder, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching todos: %v", err)})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTaskByIDHandler godoc
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
func (th *TodoHandlers) GetTaskByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := th.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error fetching todo by ID. Error: %v", err)})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTaskHandler godoc
// @Summary Update task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.UpdateTodoInput true "Todo for updating"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/update [put]
func (th *TodoHandlers) UpdateTaskHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedTodo, err := th.Repo.Update(todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error updatinh todo. Error: %v", err)})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

// DeleteTaskHandler godoc
// @Summary Delete task
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.DeleteTodoInput true "Todo for deleting"
// @Success 204 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/delete [delete]
func (th *TodoHandlers) DeleteTaskHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := th.Repo.Delete(todo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Error deleting todo. Error: %v", err)})
		return
	}
	c.Status(http.StatusNoContent)
}

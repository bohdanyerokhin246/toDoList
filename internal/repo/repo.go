package repo

import (
	"toDoList/internal/models"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	GetAll(filter models.TodoFilter, sortBy, sortOrder string, limit, offset int) ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(todo models.Todo) error
}

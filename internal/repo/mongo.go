package repo

import (
	"database/sql"
	"fmt"
	"toDoList/internal/models"
)

type MongoRepo struct {
	DB *sql.DB
}

func (mr *MongoRepo) Create(_ models.Todo) (models.Todo, error) {
	return models.Todo{Status: "Data created"}, nil
}

func (mr *MongoRepo) GetAll(_ models.TodoFilter, _ string, _ string, _ int, _ int) ([]models.Todo, error) {
	return []models.Todo{{Status: "Data selected"}, {Status: "Data selected"}}, nil
}

func (mr *MongoRepo) GetByID(_ int) (models.Todo, error) {
	return models.Todo{Status: "Data by ID selected"}, nil
}

func (mr *MongoRepo) Update(todo models.Todo) (models.Todo, error) {
	return models.Todo{Status: "Data updated"}, nil
}

func (mr *MongoRepo) Delete(_ models.Todo) error {
	fmt.Println("Data deleted")
	return nil
}

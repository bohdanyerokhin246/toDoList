package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoFilter struct {
	Status string `json:"status"`
}

type CreateTodoInput struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateTodoInput struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type DeleteTodoInput struct {
	ID int `json:"id"`
}

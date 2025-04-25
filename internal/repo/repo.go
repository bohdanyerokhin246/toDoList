package repo

import (
	"database/sql"
	"fmt"
	"time"
	"toDoList/internal/models"
)

func CreateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	query := `
        INSERT INTO todos (description, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, description, status, created_at, updated_at
    `

	err := db.QueryRow(query, todo.Description, todo.Status, time.Now(), time.Now()).Scan(
		&todo.ID, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to insert todo: %w", err)
	}

	return todo, nil
}

func GetTodoByID(db *sql.DB, id int) (models.Todo, error) {
	var todo models.Todo

	query := `
        SELECT id, description, status, created_at, updated_at
        FROM todos
        WHERE id = $1
    `

	err := db.QueryRow(query, id).Scan(
		&todo.ID, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Todo{}, fmt.Errorf("query error: %w", err)
		}
		return models.Todo{}, fmt.Errorf("query error: %w", err)
	}

	return todo, nil
}

func GetAllTodos(db *sql.DB) ([]models.Todo, error) {
	var todos []models.Todo

	query := "SELECT id, description, status, created_at, updated_at FROM todos"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows iteration: %v", err)
	}

	return todos, nil
}

func UpdateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	query := `
		UPDATE todos 
		SET 
			description = COALESCE(NULLIF($1, ''), description),
			status = COALESCE(NULLIF($2, ''), status),
			updated_at = $3
		WHERE id = $4
		RETURNING id, description, status, created_at, updated_at
	`

	err := db.QueryRow(query, todo.Description, todo.Status, time.Now(), todo.ID).
		Scan(
			&todo.ID,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

	if err == sql.ErrNoRows {
		return models.Todo{}, err
	} else if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(db *sql.DB, todo models.Todo) error {
	query := "DELETE FROM todos WHERE id = $1"

	result, err := db.Exec(query, todo.ID)
	if err != nil {
		return fmt.Errorf("error executing delete query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("noting affected: %v", err)
	}

	return nil
}

package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"toDoList/internal/models"
)

func CreateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	query := `
        INSERT INTO todos (description, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, description, status, created_at, updated_at
    `

	err := db.
		QueryRow(
			query,
			todo.Description,
			todo.Status,
			time.Now(),
			time.Now()).
		Scan(
			&todo.ID,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt,
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

	err := db.
		QueryRow(
			query,
			id).
		Scan(
			&todo.ID,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

	if err != nil {
		return models.Todo{}, fmt.Errorf("error scanning todo: %v", err)
	}

	return todo, nil
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

	err := db.
		QueryRow(
			query,
			todo.Description,
			todo.Status,
			time.Now(),
			todo.ID).
		Scan(
			&todo.ID,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt)

	if err != nil {
		return models.Todo{}, fmt.Errorf("error updating todo: %v", err)
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

func GetTodos(
	db *sql.DB,
	filter models.TodoFilter,
	sortBy string,
	sortOrder string,
	limit int,
	offset int) ([]models.Todo, error) {

	// Protection against SQL injection for sortBy and sortOrder
	allowedSortFields := map[string]bool{
		"id":          true,
		"description": true,
		"status":      true,
		"created_at":  true,
		"updated_at":  true,
	}
	if !allowedSortFields[sortBy] {
		sortBy = "created_at" // default sort
	}

	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "ASC" // default order
	}

	query := fmt.Sprintf(`
		SELECT id, description, status, created_at, updated_at
		FROM todos
		WHERE ($1::text IS NULL OR status = $1)
		ORDER BY %s %s
		LIMIT $2 OFFSET $3;
	`, sortBy, sortOrder)

	// If filter.Status is empty, pass NULL instead of a string
	var statusFilter interface{}
	if filter.Status == "" {
		statusFilter = nil
	} else {
		statusFilter = filter.Status
	}

	rows, err := db.
		Query(
			query,
			statusFilter,
			limit,
			offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching todos: %v", err)
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.
			Scan(
				&todo.ID,
				&todo.Description,
				&todo.Status,
				&todo.CreatedAt,
				&todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning todo: %v", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

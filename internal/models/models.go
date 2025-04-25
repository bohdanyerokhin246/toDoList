package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // Например: "pending", "completed"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

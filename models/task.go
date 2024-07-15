package models

import "time"

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AssignedTo  Member    `json:"member"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"Updated_at"`
}

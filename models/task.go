package models

import "time"

type Task struct {
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	AssignedToMember Member    `json:"member"`
	AssignedToBoard  Board     `json:"assigned-to-board"`
	DueDate          time.Time `json:"due_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"Updated_at"`
}

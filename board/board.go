package board

import "KanbanBoard/models"

// NewBoard is board constructor
func NewBoard() *models.Board {
	name := "default"
	return &models.Board{Name: name, Columns: []string{"ToDo", "Doing", "Review", "Done"}}
}

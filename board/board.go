package board

import "KanbanBoard/models"

// NewBoard is board constructor
func NewBoard() *models.Board {
	boardName := "default"
	return &models.Board{Name: boardName, Columns: []string{"ToDo", "Doing", "Done"}}
}

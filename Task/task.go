package Task

import "KanbanBoard/models"

func NewTask(title string, description string,

// member string, board string, column string
) *models.Task {

	return &models.Task{
		Title:            title,
		Description:      description,
		AssignedToColumn: "ToDo",
	}

}

package main

import (
	"KanbanBoard/Task"
	"KanbanBoard/board"
	"fmt"
)

func main() {
	defaultBoard := board.NewBoard()
	testTask := Task.NewTask("Test", "This is a test task")
	fmt.Println(defaultBoard)
	fmt.Println(testTask)

}

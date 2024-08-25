package handlers

import (
	"fmt"
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"github.com/joaolucassilva/task-track-cli/internal/infra/database"
)

func ListAll(filterStatus string) {
	if filterStatus != "" && filterStatus != "done" && filterStatus != "todo" && filterStatus != "in-progress" {
		fmt.Println("Filter invalid")
		return
	}

	tasks := database.GetTasksFromFile()

	if tasks == nil {
		fmt.Println("No tasks found or error reading file")
		return
	}

	filter := entities.Status(filterStatus)
	for _, task := range tasks {
		if filter == "" {
			printList(&task)
			continue
		}

		if task.Status == filter {
			printList(&task)
		}
	}
}

func printList(task *entities.Task) {
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Description: %s\n", task.Description)
	fmt.Printf("Status: %s\n", task.Status)
	fmt.Printf("CreatedAt: %s\n", task.CreatedAt)
	fmt.Printf("UpdatedAt: %s\n", task.UpdatedAt)
	fmt.Println()
}

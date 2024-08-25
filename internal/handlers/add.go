package handlers

import (
	"fmt"
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"github.com/joaolucassilva/task-track-cli/internal/infra/database"
)

func Add(name string) {
	tasks := database.GetTasksFromFile()

	if CheckTaskNameExists(name, tasks) {
		fmt.Printf("There is already a task with this name.")
		return
	}

	newTask := entities.NewTask(
		getNextId(tasks),
		name,
	)

	tasks = append(tasks, *newTask)

	database.WriteTaskToFile(tasks)

	fmt.Printf("Task added successfully (ID: %d).\n", newTask.ID)
}

func getNextId(tasks []entities.Task) int64 {
	if len(tasks) == 0 {
		return 1
	}

	maxId := tasks[0].ID
	for _, task := range tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return maxId + 1
}

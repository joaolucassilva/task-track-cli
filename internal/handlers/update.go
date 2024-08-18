package handlers

import (
	"fmt"
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"strconv"
)

func UpdateDescription(ID string, description string) {
	updated := false
	tasks := GetTasksFromFile()

	if CheckTaskNameExists(description, tasks) {
		fmt.Printf("There is already a task with this name.")
		return
	}

	var idInt, _ = strconv.ParseInt(ID, 10, 64)

	for i, task := range tasks {
		if task.ID == idInt {
			err := task.UpdateDescription(description)
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			tasks[i] = task
			updated = true
			break
		}
	}

	if updated {
		WriteTaskToFile(tasks)
		return
	}

	fmt.Printf("Task with ID %s not found", ID)
}

func UpdateStatus(ID string, status entities.Status) {
	tasks := GetTasksFromFile()

	var idInt, _ = strconv.ParseInt(ID, 10, 64)
	updated := false
	for i, task := range tasks {
		if task.ID == idInt {
			err := task.UpdateStatus(status)
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			tasks[i] = task
			updated = true
			break
		}
	}

	if updated {
		WriteTaskToFile(tasks)
		return
	}

	fmt.Printf("Task with ID %s not found", ID)
}

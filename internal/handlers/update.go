package handlers

import (
	"fmt"
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

	if !updated {
		fmt.Printf("Task with ID %s not found", ID)
	}

	WriteTaskToFile(tasks)
}

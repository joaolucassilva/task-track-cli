package handlers

import "github.com/joaolucassilva/task-track-cli/internal/entities"

func CheckTaskNameExists(name string, tasks []entities.Task) bool {
	for _, task := range tasks {
		if task.Description == name {
			return true
		}
	}
	return false
}

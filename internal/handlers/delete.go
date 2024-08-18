package handlers

import (
	"fmt"
	"strconv"
)

func DeleteById(ID string) {
	tasks := GetTasksFromFile()
	var idInt, _ = strconv.ParseInt(ID, 10, 64)

	indexSearch := -1
	for i, task := range tasks {
		if task.ID == idInt {
			indexSearch = i
			break
		}
	}

	if indexSearch == -1 {
		fmt.Printf("No task found with ID %s\n", ID)
		return
	}

	tasks = append(tasks[:indexSearch], tasks[indexSearch+1:]...)

	WriteTaskToFile(tasks)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"io"
	"os"
)

var (
	fileName = "tasks.json"
)

func GetTasksFromFile() []entities.Task {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []entities.Task{}
		}

		fmt.Println("Error opening file:", err)
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	var tasks []entities.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling file:", err)
		return nil
	}

	return tasks
}

func WriteTaskToFile(tasks []entities.Task) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling file:", err)
		return
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

package entities

import (
	"errors"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusDone       Status = "done"
	StatusInProgress Status = "in-progress"
	StatusDefault    Status = ""
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewTask(id int64, description string) *Task {
	currentTime := time.Now().Format(time.RFC3339)
	return &Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
}

func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("description is empty")
	}

	t.Description = description
	t.UpdatedAt = time.Now().Format(time.RFC3339)
	return nil
}

func (t *Task) UpdateStatus(status Status) error {
	t.Status = status
	t.UpdatedAt = time.Now().Format(time.RFC3339)
	return nil
}

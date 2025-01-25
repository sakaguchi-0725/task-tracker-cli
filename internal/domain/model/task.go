package model

import "fmt"

type Task struct {
	ID          TaskID
	Title       string
	Description string
	Status      Status
}

func NewTask(id TaskID, title, desc string, status Status) (Task, error) {
	if title == "" {
		return Task{}, fmt.Errorf("title is required")
	}

	return Task{
		ID:          id,
		Title:       title,
		Description: desc,
		Status:      status,
	}, nil
}

func RecreateTask(id TaskID, title, desc string, status Status) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: desc,
		Status:      status,
	}
}

package dto

import "github.com/sakaguchi-0725/task-tracker/internal/domain/model"

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTaskDTO(task model.Task) Task {
	return Task{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status.String(),
	}
}

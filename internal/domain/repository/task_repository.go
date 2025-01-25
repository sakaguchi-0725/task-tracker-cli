//go:generate mockgen -source=task_repository.go -destination=../../mock/repository/mock_task_repository.go -package=mock
package repository

import "github.com/sakaguchi-0725/task-tracker/internal/domain/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	FindByID(id model.TaskID) (model.Task, error)
	Create(task model.Task) error
	Update(task model.Task) error
	DeleteByID(id model.TaskID) error
}

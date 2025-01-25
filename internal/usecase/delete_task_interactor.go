package usecase

import (
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/repository"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type deleteTaskInteractor struct {
	repo repository.TaskRepository
}

func (d *deleteTaskInteractor) Execute(id string, output output.DeleteTaskOutputPort) {
	taskID, err := model.NewTaskID(id)
	if err != nil {
		output.RenderError(err)
		return
	}

	if err := d.repo.DeleteByID(taskID); err != nil {
		output.RenderError(err)
		return
	}

	output.Render()
}

func NewDeleteTaskInteractor(repo repository.TaskRepository) input.DeleteTaskInputPort {
	return &deleteTaskInteractor{repo}
}

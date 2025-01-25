package usecase

import (
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/repository"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type updateTaskInteractor struct {
	repo repository.TaskRepository
}

func (u *updateTaskInteractor) Execute(input input.UpdateTaskInput, output output.UpdateTaskOutputPort) {
	id, err := model.NewTaskID(input.ID)
	if err != nil {
		output.RenderError(err)
		return
	}

	task, err := u.repo.FindByID(id)
	if err != nil {
		output.RenderError(err)
		return
	}

	if input.Title != "" {
		task.Title = input.Title
	}

	if input.Description != "" {
		task.Description = input.Description
	}

	if input.Status != "" {
		status, err := model.NewStatus(input.Status)
		if err != nil {
			output.RenderError(err)
			return
		}
		task.Status = status
	}

	if err := u.repo.Update(task); err != nil {
		output.RenderError(err)
		return
	}

	output.Render()
}

func NewUpdateTaskInteractor(repo repository.TaskRepository) input.UpdateTaskInputPort {
	return &updateTaskInteractor{repo}
}

package usecase

import (
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/repository"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type createTaskInteractor struct {
	repo repository.TaskRepository
}

func (c createTaskInteractor) Execute(input input.CreateTaskInput, output output.CreateTaskOutputPort) {
	id := model.GenerateTaskID()
	status := model.GenerateStatus()
	task, err := model.NewTask(id, input.Title, input.Description, status)
	if err != nil {
		output.RenderError(err)
		return
	}

	if err := c.repo.Create(task); err != nil {
		output.RenderError(err)
		return
	}

	output.Render()
}

func NewCreateTaskInteractor(repo repository.TaskRepository) input.CreateTaskInputPort {
	return createTaskInteractor{repo}
}

package usecase

import (
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/repository"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type getTaskListInteractor struct {
	repo repository.TaskRepository
}

func NewGetTaskListUsecase(repo repository.TaskRepository) input.GetTaskListInputPort {
	return &getTaskListInteractor{repo}
}

func (t *getTaskListInteractor) Execute(output output.GetTaskListOutputPort) {
	tasks, err := t.repo.FindAll()
	if err != nil {
		output.RenderError(err)
		return
	}

	output.Render(makeOutput(tasks))
}

func makeOutput(tasks []model.Task) []output.Task {
	outputs := make([]output.Task, len(tasks))
	for i, task := range tasks {
		outputs[i] = output.Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status.String(),
		}
	}

	return outputs
}

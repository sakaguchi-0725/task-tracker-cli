//go:generate mockgen -source=update_task.go -destination=../../mock/usecase/input/mock_update_task.go -package=mock
package input

import "github.com/sakaguchi-0725/task-tracker/internal/usecase/output"

type UpdateTaskInput struct {
	ID          string
	Title       string
	Description string
	Status      string
}

type UpdateTaskInputPort interface {
	Execute(input UpdateTaskInput, output output.UpdateTaskOutputPort)
}

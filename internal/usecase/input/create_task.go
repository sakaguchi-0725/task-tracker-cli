//go:generate mockgen -source=create_task.go -destination=../../mock/usecase/input/mock_create_task.go -package=mock
package input

import "github.com/sakaguchi-0725/task-tracker/internal/usecase/output"

type CreateTaskInput struct {
	Title       string
	Description string
}

type CreateTaskInputPort interface {
	Execute(input CreateTaskInput, output output.CreateTaskOutputPort)
}

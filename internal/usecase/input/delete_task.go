//go:generate mockgen -source=delete_task.go -destination=../../mock/usecase/input/mock_delete_task.go -package=mock
package input

import "github.com/sakaguchi-0725/task-tracker/internal/usecase/output"

type DeleteTaskInputPort interface {
	Execute(id string, output output.DeleteTaskOutputPort)
}

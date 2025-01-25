//go:generate mockgen -source=get_task_list.go -destination=../../mock/usecase/input/mock_get_task_list.go -package=mock
package input

import "github.com/sakaguchi-0725/task-tracker/internal/usecase/output"

type GetTaskListInputPort interface {
	Execute(output output.GetTaskListOutputPort)
}

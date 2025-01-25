//go:generate mockgen -source=get_task_list.go -destination=../../mock/usecase/output/mock_get_task_list.go -package=mock
package output

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
}

type GetTaskListOutputPort interface {
	Render([]Task)
	RenderError(err error)
}

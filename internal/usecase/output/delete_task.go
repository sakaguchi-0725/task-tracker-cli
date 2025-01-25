//go:generate mockgen -source=delete_task.go -destination=../../mock/usecase/output/mock_delete_task.go -package=mock
package output

type DeleteTaskOutputPort interface {
	Render()
	RenderError(err error)
}

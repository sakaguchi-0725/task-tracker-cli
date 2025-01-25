//go:generate mockgen -source=update_task.go -destination=../../mock/usecase/output/mock_update_task.go -package=mock
package output

type UpdateTaskOutputPort interface {
	Render()
	RenderError(err error)
}

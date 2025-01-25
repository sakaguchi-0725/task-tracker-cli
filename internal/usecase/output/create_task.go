//go:generate mockgen -source=create_task.go -destination=../../mock/usecase/output/mock_create_task.go -package=mock
package output

type CreateTaskOutputPort interface {
	Render()
	RenderError(err error)
}

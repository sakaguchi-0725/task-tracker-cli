package presenter

import (
	"github.com/fatih/color"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type deleteTaskPresenter struct{}

func (d *deleteTaskPresenter) Render() {
	color.Green("Task deleted successfully 🎉")
}

func (d *deleteTaskPresenter) RenderError(err error) {
	color.Red(err.Error())
}

func NewDeleteTaskPresenter() output.DeleteTaskOutputPort {
	return &deleteTaskPresenter{}
}

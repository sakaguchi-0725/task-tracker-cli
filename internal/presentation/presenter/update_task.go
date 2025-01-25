package presenter

import (
	"github.com/fatih/color"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type updateTaskPresenter struct{}

func (u *updateTaskPresenter) Render() {
	color.Green("Task updated successfully 🎉")
}

func (u *updateTaskPresenter) RenderError(err error) {
	color.Red(err.Error())
}

func NewUpdateTaskPresenter() output.UpdateTaskOutputPort {
	return &updateTaskPresenter{}
}

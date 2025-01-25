package presenter

import (
	"github.com/fatih/color"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type createTaskPresenter struct{}

func NewCreateTaskPresenter() output.CreateTaskOutputPort {
	return &createTaskPresenter{}
}

func (c *createTaskPresenter) Render() {
	color.Green("Successfully created 🎉")
}

func (c *createTaskPresenter) RenderError(err error) {
	color.Red(err.Error())
}

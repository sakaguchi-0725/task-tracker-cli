package presenter

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
)

type getTaskListPresenter struct{}

func NewGetTaskListPresenter() output.GetTaskListOutputPort {
	return &getTaskListPresenter{}
}

func (g *getTaskListPresenter) Render(tasks []output.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "title", "description", "status"})

	for _, t := range tasks {
		table.Append([]string{t.ID, t.Title, t.Description, t.Status})
	}

	table.Render()
}

func (g *getTaskListPresenter) RenderError(err error) {
	color.Red(err.Error())
}

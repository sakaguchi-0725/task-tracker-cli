package command

import (
	"github.com/sakaguchi-0725/task-tracker/internal/presentation/presenter"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/spf13/cobra"
)

type getTaskListCommand struct {
	usecase input.GetTaskListInputPort
}

func NewGetTaskListCommand(usecase input.GetTaskListInputPort) *getTaskListCommand {
	return &getTaskListCommand{usecase}
}

func (g *getTaskListCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "show all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			g.usecase.Execute(presenter.NewGetTaskListPresenter())
		},
	}

	return cmd
}

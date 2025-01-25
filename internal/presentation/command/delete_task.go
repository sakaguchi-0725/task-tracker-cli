package command

import (
	"github.com/sakaguchi-0725/task-tracker/internal/presentation/presenter"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/spf13/cobra"
)

type deleteTaskCommand struct {
	usecase input.DeleteTaskInputPort
}

func NewDeleteTaskCommand(usecase input.DeleteTaskInputPort) *deleteTaskCommand {
	return &deleteTaskCommand{usecase}
}

func (d *deleteTaskCommand) Command() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a task by id",
		Run: func(cmd *cobra.Command, args []string) {
			d.usecase.Execute(id, presenter.NewDeleteTaskPresenter())
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID of the task")
	return cmd
}

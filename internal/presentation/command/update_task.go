package command

import (
	"github.com/sakaguchi-0725/task-tracker/internal/presentation/presenter"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/spf13/cobra"
)

type updateTaskCommand struct {
	usecase input.UpdateTaskInputPort
}

func NewUpdateTaskCommand(usecase input.UpdateTaskInputPort) *updateTaskCommand {
	return &updateTaskCommand{usecase}
}

func (u *updateTaskCommand) Command() *cobra.Command {
	var id, title, desc, status string

	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update a task",
		Run: func(cmd *cobra.Command, args []string) {
			input := input.UpdateTaskInput{
				ID:          id,
				Title:       title,
				Description: desc,
				Status:      status,
			}

			u.usecase.Execute(input, presenter.NewUpdateTaskPresenter())
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID of the task")
	cmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	cmd.Flags().StringVarP(&desc, "description", "d", "", "Description of the task")
	cmd.Flags().StringVarP(&status, "status", "s", "", "Status of the task")

	return cmd
}

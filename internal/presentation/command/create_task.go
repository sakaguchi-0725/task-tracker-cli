package command

import (
	"github.com/sakaguchi-0725/task-tracker/internal/presentation/presenter"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/spf13/cobra"
)

type createTaskCommand struct {
	usecase input.CreateTaskInputPort
}

func NewCreateTaskCommand(usecase input.CreateTaskInputPort) *createTaskCommand {
	return &createTaskCommand{usecase}
}

func (c *createTaskCommand) Command() *cobra.Command {
	var title string
	var desc string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			input := input.CreateTaskInput{
				Title:       title,
				Description: desc,
			}
			c.usecase.Execute(input, presenter.NewCreateTaskPresenter())
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	cmd.Flags().StringVarP(&desc, "description", "d", "", "Description of the task")

	return cmd
}

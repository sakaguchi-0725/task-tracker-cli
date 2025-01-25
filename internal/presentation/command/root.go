package command

import (
	"os"

	"github.com/spf13/cobra"
)

type rootCommand struct {
	cmd *cobra.Command
}

func NewRootCoomand() *rootCommand {
	rootCmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "A brief description of your application",
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return &rootCommand{cmd: rootCmd}
}

func (r *rootCommand) Execute() {
	if err := r.cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (r *rootCommand) AddCommand(cmd *cobra.Command) {
	r.cmd.AddCommand(cmd)
}

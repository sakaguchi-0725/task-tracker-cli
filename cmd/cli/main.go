package main

import (
	"fmt"
	"os"

	"github.com/sakaguchi-0725/task-tracker/internal/infra/persistence"
	"github.com/sakaguchi-0725/task-tracker/internal/presentation/command"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase"
	"github.com/sakaguchi-0725/task-tracker/internal/util"
)

func main() {
	if !util.IsJsonExists("task.json") {
		if err := util.CreateEmptyJSON("task.json"); err != nil {
			fmt.Printf("failed create json file: %v", err)
			os.Exit(1)
		}
	}

	rootCmd := command.NewRootCoomand()
	repo := persistence.NewTaskPersistence("task.json")

	createTaskUsecase := usecase.NewCreateTaskInteractor(repo)
	getTaskListUsecase := usecase.NewGetTaskListUsecase(repo)
	updateTaskUsecase := usecase.NewUpdateTaskInteractor(repo)
	deleteTaskUsecase := usecase.NewDeleteTaskInteractor(repo)

	createTaskCmd := command.NewCreateTaskCommand(createTaskUsecase)
	getTaskListCmd := command.NewGetTaskListCommand(getTaskListUsecase)
	updateTaskCmd := command.NewUpdateTaskCommand(updateTaskUsecase)
	deleteTaskCmd := command.NewDeleteTaskCommand(deleteTaskUsecase)

	rootCmd.AddCommand(createTaskCmd.Command())
	rootCmd.AddCommand(getTaskListCmd.Command())
	rootCmd.AddCommand(updateTaskCmd.Command())
	rootCmd.AddCommand(deleteTaskCmd.Command())

	rootCmd.Execute()
}

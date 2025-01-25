package usecase_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	. "github.com/sakaguchi-0725/task-tracker/internal/mock/repository"
	. "github.com/sakaguchi-0725/task-tracker/internal/mock/usecase/output"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestGetTaskListInteractor(t *testing.T) {
	usecase, repo, outputPort := setupGetTaskListTest(t)
	models, outputs := generateMockTaskList(2)

	tests := map[string]struct {
		setupMock func(*MockTaskRepository, *MockGetTaskListOutputPort)
		expected  string
	}{
		"Success": {
			setupMock: func(mtr *MockTaskRepository, gtlop *MockGetTaskListOutputPort) {
				mtr.EXPECT().FindAll().Return(models, nil)
				gtlop.EXPECT().Render(outputs).Do(func(tasks []output.Task) {
					for _, t := range tasks {
						fmt.Printf("ID: %s, Title: %s, Status: %s\n", t.ID, t.Title, t.Status)
					}
				})
			},
			expected: generateExpectedOutput(outputs),
		},
		"Failed": {
			setupMock: func(mtr *MockTaskRepository, gtlop *MockGetTaskListOutputPort) {
				mtr.EXPECT().FindAll().Return([]model.Task{}, errors.New("error"))
				gtlop.EXPECT().RenderError(errors.New("error")).Do(func(err error) {
					fmt.Print(err.Error())
				})
			},
			expected: "error",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.setupMock(repo, outputPort)

			output := captureOutput(func() {
				usecase.Execute(outputPort)
			})

			assert.Equal(t, tt.expected, output)
		})
	}
}

func generateExpectedOutput(tasks []output.Task) string {
	var sb strings.Builder
	for _, t := range tasks {
		sb.WriteString("ID: ")
		sb.WriteString(t.ID)
		sb.WriteString(", Title: ")
		sb.WriteString(t.Title)
		sb.WriteString(", Status: ")
		sb.WriteString(t.Status)
		sb.WriteString("\n")
	}
	return sb.String()
}

func generateMockTaskList(count int) ([]model.Task, []output.Task) {
	models := make([]model.Task, count)
	outputs := make([]output.Task, count)

	for i := 0; i < count; i++ {
		id := model.GenerateTaskID()
		num := strconv.Itoa(i)

		models = append(models, model.Task{ID: id, Title: fmt.Sprintf("Task%s", num), Description: "", Status: model.NotStarted})
		outputs = append(outputs, output.Task{ID: id.String(), Title: fmt.Sprintf("Task%s", num), Description: "", Status: model.NotStarted.String()})
	}

	return models, outputs
}

func setupGetTaskListTest(t *testing.T) (input.GetTaskListInputPort, *MockTaskRepository, *MockGetTaskListOutputPort) {
	ctrl := gomock.NewController(t)

	output := NewMockGetTaskListOutputPort(ctrl)
	repo := NewMockTaskRepository(ctrl)
	usecase := usecase.NewGetTaskListUsecase(repo)

	return usecase, repo, output
}

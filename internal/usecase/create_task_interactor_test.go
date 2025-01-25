package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/sakaguchi-0725/task-tracker/internal/mock/repository"
	. "github.com/sakaguchi-0725/task-tracker/internal/mock/usecase/output"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTaskInteractor(t *testing.T) {
	usecase, repo, output := setupCreateTaskTest(t)

	tests := map[string]struct {
		setupMock func(*MockTaskRepository, *MockCreateTaskOutputPort)
		input     input.CreateTaskInput
		expected  string
	}{
		"Success": {
			setupMock: func(tr *MockTaskRepository, ctop *MockCreateTaskOutputPort) {
				tr.EXPECT().Create(gomock.Any()).Return(nil)
				ctop.EXPECT().Render().Do(func() {
					fmt.Print("success")
				})
			},
			input:    input.CreateTaskInput{Title: "Test", Description: "Test"},
			expected: "success",
		},
		"Failed": {
			setupMock: func(tr *MockTaskRepository, ctop *MockCreateTaskOutputPort) {
				tr.EXPECT().Create(gomock.Any()).Return(errors.New("error"))
				ctop.EXPECT().RenderError(errors.New("error")).Do(func(err error) {
					fmt.Print(err)
				})
			},
			input:    input.CreateTaskInput{Title: "Test", Description: "Test"},
			expected: "error",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.setupMock(repo, output)
			out := captureOutput(func() {
				usecase.Execute(tt.input, output)
			})

			assert.Contains(t, tt.expected, string(out))
		})
	}
}

func setupCreateTaskTest(t *testing.T) (input.CreateTaskInputPort, *MockTaskRepository, *MockCreateTaskOutputPort) {
	ctrl := gomock.NewController(t)

	repo := NewMockTaskRepository(ctrl)
	output := NewMockCreateTaskOutputPort(ctrl)
	usecase := usecase.NewCreateTaskInteractor(repo)

	return usecase, repo, output
}

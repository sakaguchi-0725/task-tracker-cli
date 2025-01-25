package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	. "github.com/sakaguchi-0725/task-tracker/internal/mock/repository"
	. "github.com/sakaguchi-0725/task-tracker/internal/mock/usecase/output"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase"
	"github.com/sakaguchi-0725/task-tracker/internal/usecase/input"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateTaskInteractor(t *testing.T) {
	usecase, repo, output := setupUpdateTaskTest(t)
	mockID := model.GenerateTaskID()

	tests := map[string]struct {
		setupMock func(*MockTaskRepository, *MockUpdateTaskOutputPort)
		input     input.UpdateTaskInput
		expected  string
	}{
		"Success": {
			setupMock: func(mtr *MockTaskRepository, utop *MockUpdateTaskOutputPort) {
				mtr.EXPECT().FindByID(mockID).Return(
					model.Task{ID: mockID, Title: "Task", Description: "", Status: model.NotStarted}, nil,
				)
				mtr.EXPECT().Update(model.Task{ID: mockID, Title: "Update Task", Description: "", Status: model.InProgress}).Return(nil)
				utop.EXPECT().Render().Do(func() {
					fmt.Print("success")
				})
			},
			input:    input.UpdateTaskInput{ID: mockID.String(), Title: "Update Task", Description: "", Status: model.InProgress.String()},
			expected: "success",
		},
		"Failed": {
			setupMock: func(mtr *MockTaskRepository, utop *MockUpdateTaskOutputPort) {
				mtr.EXPECT().FindByID(mockID).Return(model.Task{}, errors.New("error"))
				utop.EXPECT().RenderError(errors.New("error")).Do(func(err error) {
					fmt.Print(err)
				})
			},
			input:    input.UpdateTaskInput{ID: mockID.String(), Title: "Update Failed", Description: "", Status: model.NotStarted.String()},
			expected: "error",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.setupMock(repo, output)

			out := captureOutput(func() {
				usecase.Execute(tt.input, output)
			})

			assert.Equal(t, tt.expected, out)
		})
	}
}

func setupUpdateTaskTest(t *testing.T) (input.UpdateTaskInputPort, *MockTaskRepository, *MockUpdateTaskOutputPort) {
	ctrl := gomock.NewController(t)

	output := NewMockUpdateTaskOutputPort(ctrl)
	repo := NewMockTaskRepository(ctrl)
	usecase := usecase.NewUpdateTaskInteractor(repo)

	return usecase, repo, output
}

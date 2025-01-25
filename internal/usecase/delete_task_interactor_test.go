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

func TestDeleteTaskInteractor(t *testing.T) {
	usecase, repo, output := setupDeleteTaskTest(t)
	mockID := model.GenerateTaskID()

	tests := map[string]struct {
		setupMock func(*MockTaskRepository, *MockDeleteTaskOutputPort)
		input     string
		expected  string
	}{
		"Success": {
			setupMock: func(mtr *MockTaskRepository, dtop *MockDeleteTaskOutputPort) {
				mtr.EXPECT().DeleteByID(mockID).Return(nil)
				dtop.EXPECT().Render().Do(func() {
					fmt.Print("success")
				})
			},
			input:    mockID.String(),
			expected: "success",
		},
		"Failed": {
			setupMock: func(mtr *MockTaskRepository, dtop *MockDeleteTaskOutputPort) {
				mtr.EXPECT().DeleteByID(mockID).Return(errors.New("error"))
				dtop.EXPECT().RenderError(errors.New("error")).Do(func(err error) {
					fmt.Print(err)
				})
			},
			input:    mockID.String(),
			expected: "error",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.setupMock(repo, output)
			out := captureOutput(func() {
				usecase.Execute(mockID.String(), output)
			})

			assert.Contains(t, tt.expected, string(out))
		})
	}
}

func setupDeleteTaskTest(t *testing.T) (input.DeleteTaskInputPort, *MockTaskRepository, *MockDeleteTaskOutputPort) {
	ctrl := gomock.NewController(t)

	repo := NewMockTaskRepository(ctrl)
	output := NewMockDeleteTaskOutputPort(ctrl)
	usecase := usecase.NewDeleteTaskInteractor(repo)

	return usecase, repo, output
}

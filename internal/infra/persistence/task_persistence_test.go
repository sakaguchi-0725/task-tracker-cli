package persistence_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sakaguchi-0725/task-tracker/internal/domain/model"
	"github.com/sakaguchi-0725/task-tracker/internal/infra/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskRepository(t *testing.T) {
	repo := persistence.NewTaskPersistence("test_task.json")

	t.Run("Create", func(t *testing.T) {
		defer clearJSON()

		tests := map[string]struct {
			input       model.Task
			expectedErr error
		}{
			"data can be created": {
				input: model.Task{
					ID:          model.GenerateTaskID(),
					Title:       "Task1",
					Description: "This is test",
					Status:      model.NotStarted,
				},
				expectedErr: nil,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := repo.Create(tt.input)

				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				}
			})
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		defer clearJSON()

		mockData := []model.Task{
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task1",
				Description: "Task1",
				Status:      model.NotStarted,
			},
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task2",
				Description: "Task2",
				Status:      model.InProgress,
			},
		}

		tests := map[string]struct {
			expected    []model.Task
			expectedErr error
		}{
			"Success": {
				expected:    mockData,
				expectedErr: nil,
			},
			"NoData": {
				expected:    []model.Task{},
				expectedErr: nil,
			},
		}

		for name, tt := range tests {
			if len(tt.expected) != 0 {
				for _, task := range mockData {
					err := repo.Create(task)
					require.NoError(t, err)
				}
			}

			t.Run(name, func(t *testing.T) {
				defer clearJSON()
				result, err := repo.FindAll()

				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
					if diff := cmp.Diff(tt.expected, result); diff != "" {
						t.Errorf("model.Task value is mismatch (-expected +result):%s\n", diff)
					}
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		defer clearJSON()

		mockData := []model.Task{
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task1",
				Description: "Task1",
				Status:      model.NotStarted,
			},
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task2",
				Description: "Task2",
				Status:      model.InProgress,
			},
		}

		tests := map[string]struct {
			input       model.Task
			expectedErr error
		}{
			"Success": {
				input: model.Task{
					ID:          mockData[0].ID,
					Title:       "Update",
					Description: "Updated",
					Status:      model.InProgress,
				},
				expectedErr: nil,
			},
			"ID does not exist": {
				input: model.Task{
					ID:          model.GenerateTaskID(),
					Title:       "Update failed",
					Description: "Update failed",
					Status:      model.Done,
				},
				expectedErr: errors.New("record with the specified ID not found"),
			},
		}

		t.Run("create mock data", func(t *testing.T) {
			for _, v := range mockData {
				err := repo.Create(v)
				require.NoError(t, err)
			}
		})

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := repo.Update(tt.input)

				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		defer clearJSON()

		mockData := []model.Task{
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task1",
				Description: "Task1",
				Status:      model.NotStarted,
			},
			{
				ID:          model.GenerateTaskID(),
				Title:       "Task2",
				Description: "Task2",
				Status:      model.InProgress,
			},
		}

		t.Run("create mock data", func(t *testing.T) {
			for _, v := range mockData {
				err := repo.Create(v)
				require.NoError(t, err)
			}
		})

		tests := map[string]struct {
			input       model.TaskID
			expected    model.Task
			expectedErr error
		}{
			"Success": {
				input:       mockData[0].ID,
				expected:    mockData[0],
				expectedErr: nil,
			},
			"Data not found": {
				input:       model.GenerateTaskID(),
				expected:    model.Task{},
				expectedErr: errors.New("record not found"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				task, err := repo.FindByID(tt.input)

				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
					if diff := cmp.Diff(tt.expected, task); diff != "" {
						t.Errorf("model.Task value is mismatch (-expected +result):%s\n", diff)
					}
				}
			})
		}
	})

	t.Run("DeleteByID", func(t *testing.T) {
		defer clearJSON()

		mockData := model.Task{
			ID:          model.GenerateTaskID(),
			Title:       "Task1",
			Description: "Task1",
			Status:      model.NotStarted,
		}

		t.Run("create mock data", func(t *testing.T) {
			err := repo.Create(mockData)
			require.NoError(t, err)
		})

		tests := map[string]struct {
			input       model.TaskID
			expectedErr error
		}{
			"Success": {
				input:       mockData.ID,
				expectedErr: nil,
			},
			"record not found": {
				input:       model.GenerateTaskID(),
				expectedErr: errors.New("record not found"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := repo.DeleteByID(tt.input)

				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}

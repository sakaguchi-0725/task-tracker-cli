package model

import (
	"fmt"

	"github.com/google/uuid"
)

type TaskID string

func NewTaskID(s string) (TaskID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", fmt.Errorf("taskID parse error: %v", err)
	}

	return TaskID(id.String()), nil
}

func GenerateTaskID() TaskID {
	return TaskID(uuid.NewString())
}

func (t *TaskID) String() string {
	return string(*t)
}

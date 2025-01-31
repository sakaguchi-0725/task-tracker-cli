// Code generated by MockGen. DO NOT EDIT.
// Source: delete_task.go
//
// Generated by this command:
//
//	mockgen -source=delete_task.go -destination=../../mock/usecase/input/mock_delete_task.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	output "github.com/sakaguchi-0725/task-tracker/internal/usecase/output"
	gomock "go.uber.org/mock/gomock"
)

// MockDeleteTaskInputPort is a mock of DeleteTaskInputPort interface.
type MockDeleteTaskInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteTaskInputPortMockRecorder
	isgomock struct{}
}

// MockDeleteTaskInputPortMockRecorder is the mock recorder for MockDeleteTaskInputPort.
type MockDeleteTaskInputPortMockRecorder struct {
	mock *MockDeleteTaskInputPort
}

// NewMockDeleteTaskInputPort creates a new mock instance.
func NewMockDeleteTaskInputPort(ctrl *gomock.Controller) *MockDeleteTaskInputPort {
	mock := &MockDeleteTaskInputPort{ctrl: ctrl}
	mock.recorder = &MockDeleteTaskInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteTaskInputPort) EXPECT() *MockDeleteTaskInputPortMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDeleteTaskInputPort) Execute(id string, output output.DeleteTaskOutputPort) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Execute", id, output)
}

// Execute indicates an expected call of Execute.
func (mr *MockDeleteTaskInputPortMockRecorder) Execute(id, output any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDeleteTaskInputPort)(nil).Execute), id, output)
}

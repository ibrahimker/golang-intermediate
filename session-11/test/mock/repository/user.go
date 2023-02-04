// Code generated by MockGen. DO NOT EDIT.
// Source: repository/user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockIUser) Register(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockIUserMockRecorder) Register(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIUser)(nil).Register), username, password)
}

// RegisterWithTimestamp mocks base method.
func (m *MockIUser) RegisterWithTimestamp(username, password string, createdAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterWithTimestamp", username, password, createdAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterWithTimestamp indicates an expected call of RegisterWithTimestamp.
func (mr *MockIUserMockRecorder) RegisterWithTimestamp(username, password, createdAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterWithTimestamp", reflect.TypeOf((*MockIUser)(nil).RegisterWithTimestamp), username, password, createdAt)
}

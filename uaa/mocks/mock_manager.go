// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pivotalservices/cf-mgmt/uaa (interfaces: Manager)

// Package mock_uaa is a generated GoMock package.
package mock_uaa

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uaa "github.com/pivotalservices/cf-mgmt/uaa"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateExternalUser mocks base method
func (m *MockManager) CreateExternalUser(arg0, arg1, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "CreateExternalUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExternalUser indicates an expected call of CreateExternalUser
func (mr *MockManagerMockRecorder) CreateExternalUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalUser", reflect.TypeOf((*MockManager)(nil).CreateExternalUser), arg0, arg1, arg2, arg3)
}

// ListUsers mocks base method
func (m *MockManager) ListUsers() (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListUsers")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockManagerMockRecorder) ListUsers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockManager)(nil).ListUsers))
}

// UsersByID mocks base method
func (m *MockManager) UsersByID() (map[string]uaa.User, error) {
	ret := m.ctrl.Call(m, "UsersByID")
	ret0, _ := ret[0].(map[string]uaa.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersByID indicates an expected call of UsersByID
func (mr *MockManagerMockRecorder) UsersByID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersByID", reflect.TypeOf((*MockManager)(nil).UsersByID))
}
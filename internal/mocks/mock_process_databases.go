// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: ProcessDatabaseLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241113005/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockProcessDatabaseLister is a mock of ProcessDatabaseLister interface.
type MockProcessDatabaseLister struct {
	ctrl     *gomock.Controller
	recorder *MockProcessDatabaseListerMockRecorder
}

// MockProcessDatabaseListerMockRecorder is the mock recorder for MockProcessDatabaseLister.
type MockProcessDatabaseListerMockRecorder struct {
	mock *MockProcessDatabaseLister
}

// NewMockProcessDatabaseLister creates a new mock instance.
func NewMockProcessDatabaseLister(ctrl *gomock.Controller) *MockProcessDatabaseLister {
	mock := &MockProcessDatabaseLister{ctrl: ctrl}
	mock.recorder = &MockProcessDatabaseListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessDatabaseLister) EXPECT() *MockProcessDatabaseListerMockRecorder {
	return m.recorder
}

// ProcessDatabases mocks base method.
func (m *MockProcessDatabaseLister) ProcessDatabases(arg0, arg1 string, arg2 int, arg3 *mongodbatlas.ListOptions) (*admin.PaginatedDatabase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDatabases", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.PaginatedDatabase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessDatabases indicates an expected call of ProcessDatabases.
func (mr *MockProcessDatabaseListerMockRecorder) ProcessDatabases(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDatabases", reflect.TypeOf((*MockProcessDatabaseLister)(nil).ProcessDatabases), arg0, arg1, arg2, arg3)
}

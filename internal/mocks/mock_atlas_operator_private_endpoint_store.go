// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: OperatorPrivateEndpointStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

// MockOperatorPrivateEndpointStore is a mock of OperatorPrivateEndpointStore interface.
type MockOperatorPrivateEndpointStore struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorPrivateEndpointStoreMockRecorder
}

// MockOperatorPrivateEndpointStoreMockRecorder is the mock recorder for MockOperatorPrivateEndpointStore.
type MockOperatorPrivateEndpointStoreMockRecorder struct {
	mock *MockOperatorPrivateEndpointStore
}

// NewMockOperatorPrivateEndpointStore creates a new mock instance.
func NewMockOperatorPrivateEndpointStore(ctrl *gomock.Controller) *MockOperatorPrivateEndpointStore {
	mock := &MockOperatorPrivateEndpointStore{ctrl: ctrl}
	mock.recorder = &MockOperatorPrivateEndpointStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorPrivateEndpointStore) EXPECT() *MockOperatorPrivateEndpointStoreMockRecorder {
	return m.recorder
}

// InterfaceEndpoint mocks base method.
func (m *MockOperatorPrivateEndpointStore) InterfaceEndpoint(arg0, arg1, arg2, arg3 string) (*admin.PrivateLinkEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterfaceEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.PrivateLinkEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceEndpoint indicates an expected call of InterfaceEndpoint.
func (mr *MockOperatorPrivateEndpointStoreMockRecorder) InterfaceEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceEndpoint", reflect.TypeOf((*MockOperatorPrivateEndpointStore)(nil).InterfaceEndpoint), arg0, arg1, arg2, arg3)
}

// PrivateEndpoints mocks base method.
func (m *MockOperatorPrivateEndpointStore) PrivateEndpoints(arg0, arg1 string) ([]admin.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockOperatorPrivateEndpointStoreMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockOperatorPrivateEndpointStore)(nil).PrivateEndpoints), arg0, arg1)
}

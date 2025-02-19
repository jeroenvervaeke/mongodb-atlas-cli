// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: AlertConfigurationLister,AlertConfigurationCreator,AlertConfigurationDeleter,AlertConfigurationUpdater,MatcherFieldsLister,AlertConfigurationEnabler,AlertConfigurationDisabler,AlertConfigurationDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

// MockAlertConfigurationLister is a mock of AlertConfigurationLister interface.
type MockAlertConfigurationLister struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationListerMockRecorder
}

// MockAlertConfigurationListerMockRecorder is the mock recorder for MockAlertConfigurationLister.
type MockAlertConfigurationListerMockRecorder struct {
	mock *MockAlertConfigurationLister
}

// NewMockAlertConfigurationLister creates a new mock instance.
func NewMockAlertConfigurationLister(ctrl *gomock.Controller) *MockAlertConfigurationLister {
	mock := &MockAlertConfigurationLister{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationLister) EXPECT() *MockAlertConfigurationListerMockRecorder {
	return m.recorder
}

// AlertConfigurations mocks base method.
func (m *MockAlertConfigurationLister) AlertConfigurations(arg0 *admin.ListAlertConfigurationsApiParams) (*admin.PaginatedAlertConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfigurations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedAlertConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfigurations indicates an expected call of AlertConfigurations.
func (mr *MockAlertConfigurationListerMockRecorder) AlertConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfigurations", reflect.TypeOf((*MockAlertConfigurationLister)(nil).AlertConfigurations), arg0)
}

// MockAlertConfigurationCreator is a mock of AlertConfigurationCreator interface.
type MockAlertConfigurationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationCreatorMockRecorder
}

// MockAlertConfigurationCreatorMockRecorder is the mock recorder for MockAlertConfigurationCreator.
type MockAlertConfigurationCreatorMockRecorder struct {
	mock *MockAlertConfigurationCreator
}

// NewMockAlertConfigurationCreator creates a new mock instance.
func NewMockAlertConfigurationCreator(ctrl *gomock.Controller) *MockAlertConfigurationCreator {
	mock := &MockAlertConfigurationCreator{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationCreator) EXPECT() *MockAlertConfigurationCreatorMockRecorder {
	return m.recorder
}

// CreateAlertConfiguration mocks base method.
func (m *MockAlertConfigurationCreator) CreateAlertConfiguration(arg0 *admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlertConfiguration", arg0)
	ret0, _ := ret[0].(*admin.GroupAlertsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlertConfiguration indicates an expected call of CreateAlertConfiguration.
func (mr *MockAlertConfigurationCreatorMockRecorder) CreateAlertConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationCreator)(nil).CreateAlertConfiguration), arg0)
}

// MockAlertConfigurationDeleter is a mock of AlertConfigurationDeleter interface.
type MockAlertConfigurationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationDeleterMockRecorder
}

// MockAlertConfigurationDeleterMockRecorder is the mock recorder for MockAlertConfigurationDeleter.
type MockAlertConfigurationDeleterMockRecorder struct {
	mock *MockAlertConfigurationDeleter
}

// NewMockAlertConfigurationDeleter creates a new mock instance.
func NewMockAlertConfigurationDeleter(ctrl *gomock.Controller) *MockAlertConfigurationDeleter {
	mock := &MockAlertConfigurationDeleter{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationDeleter) EXPECT() *MockAlertConfigurationDeleterMockRecorder {
	return m.recorder
}

// DeleteAlertConfiguration mocks base method.
func (m *MockAlertConfigurationDeleter) DeleteAlertConfiguration(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlertConfiguration indicates an expected call of DeleteAlertConfiguration.
func (mr *MockAlertConfigurationDeleterMockRecorder) DeleteAlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationDeleter)(nil).DeleteAlertConfiguration), arg0, arg1)
}

// MockAlertConfigurationUpdater is a mock of AlertConfigurationUpdater interface.
type MockAlertConfigurationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationUpdaterMockRecorder
}

// MockAlertConfigurationUpdaterMockRecorder is the mock recorder for MockAlertConfigurationUpdater.
type MockAlertConfigurationUpdaterMockRecorder struct {
	mock *MockAlertConfigurationUpdater
}

// NewMockAlertConfigurationUpdater creates a new mock instance.
func NewMockAlertConfigurationUpdater(ctrl *gomock.Controller) *MockAlertConfigurationUpdater {
	mock := &MockAlertConfigurationUpdater{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationUpdater) EXPECT() *MockAlertConfigurationUpdaterMockRecorder {
	return m.recorder
}

// UpdateAlertConfiguration mocks base method.
func (m *MockAlertConfigurationUpdater) UpdateAlertConfiguration(arg0 *admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlertConfiguration", arg0)
	ret0, _ := ret[0].(*admin.GroupAlertsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAlertConfiguration indicates an expected call of UpdateAlertConfiguration.
func (mr *MockAlertConfigurationUpdaterMockRecorder) UpdateAlertConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationUpdater)(nil).UpdateAlertConfiguration), arg0)
}

// MockMatcherFieldsLister is a mock of MatcherFieldsLister interface.
type MockMatcherFieldsLister struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherFieldsListerMockRecorder
}

// MockMatcherFieldsListerMockRecorder is the mock recorder for MockMatcherFieldsLister.
type MockMatcherFieldsListerMockRecorder struct {
	mock *MockMatcherFieldsLister
}

// NewMockMatcherFieldsLister creates a new mock instance.
func NewMockMatcherFieldsLister(ctrl *gomock.Controller) *MockMatcherFieldsLister {
	mock := &MockMatcherFieldsLister{ctrl: ctrl}
	mock.recorder = &MockMatcherFieldsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcherFieldsLister) EXPECT() *MockMatcherFieldsListerMockRecorder {
	return m.recorder
}

// MatcherFields mocks base method.
func (m *MockMatcherFieldsLister) MatcherFields() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatcherFields")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatcherFields indicates an expected call of MatcherFields.
func (mr *MockMatcherFieldsListerMockRecorder) MatcherFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatcherFields", reflect.TypeOf((*MockMatcherFieldsLister)(nil).MatcherFields))
}

// MockAlertConfigurationEnabler is a mock of AlertConfigurationEnabler interface.
type MockAlertConfigurationEnabler struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationEnablerMockRecorder
}

// MockAlertConfigurationEnablerMockRecorder is the mock recorder for MockAlertConfigurationEnabler.
type MockAlertConfigurationEnablerMockRecorder struct {
	mock *MockAlertConfigurationEnabler
}

// NewMockAlertConfigurationEnabler creates a new mock instance.
func NewMockAlertConfigurationEnabler(ctrl *gomock.Controller) *MockAlertConfigurationEnabler {
	mock := &MockAlertConfigurationEnabler{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationEnablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationEnabler) EXPECT() *MockAlertConfigurationEnablerMockRecorder {
	return m.recorder
}

// EnableAlertConfiguration mocks base method.
func (m *MockAlertConfigurationEnabler) EnableAlertConfiguration(arg0, arg1 string) (*admin.GroupAlertsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableAlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.GroupAlertsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableAlertConfiguration indicates an expected call of EnableAlertConfiguration.
func (mr *MockAlertConfigurationEnablerMockRecorder) EnableAlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationEnabler)(nil).EnableAlertConfiguration), arg0, arg1)
}

// MockAlertConfigurationDisabler is a mock of AlertConfigurationDisabler interface.
type MockAlertConfigurationDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationDisablerMockRecorder
}

// MockAlertConfigurationDisablerMockRecorder is the mock recorder for MockAlertConfigurationDisabler.
type MockAlertConfigurationDisablerMockRecorder struct {
	mock *MockAlertConfigurationDisabler
}

// NewMockAlertConfigurationDisabler creates a new mock instance.
func NewMockAlertConfigurationDisabler(ctrl *gomock.Controller) *MockAlertConfigurationDisabler {
	mock := &MockAlertConfigurationDisabler{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationDisabler) EXPECT() *MockAlertConfigurationDisablerMockRecorder {
	return m.recorder
}

// DisableAlertConfiguration mocks base method.
func (m *MockAlertConfigurationDisabler) DisableAlertConfiguration(arg0, arg1 string) (*admin.GroupAlertsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableAlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.GroupAlertsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableAlertConfiguration indicates an expected call of DisableAlertConfiguration.
func (mr *MockAlertConfigurationDisablerMockRecorder) DisableAlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAlertConfiguration", reflect.TypeOf((*MockAlertConfigurationDisabler)(nil).DisableAlertConfiguration), arg0, arg1)
}

// MockAlertConfigurationDescriber is a mock of AlertConfigurationDescriber interface.
type MockAlertConfigurationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockAlertConfigurationDescriberMockRecorder
}

// MockAlertConfigurationDescriberMockRecorder is the mock recorder for MockAlertConfigurationDescriber.
type MockAlertConfigurationDescriberMockRecorder struct {
	mock *MockAlertConfigurationDescriber
}

// NewMockAlertConfigurationDescriber creates a new mock instance.
func NewMockAlertConfigurationDescriber(ctrl *gomock.Controller) *MockAlertConfigurationDescriber {
	mock := &MockAlertConfigurationDescriber{ctrl: ctrl}
	mock.recorder = &MockAlertConfigurationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertConfigurationDescriber) EXPECT() *MockAlertConfigurationDescriberMockRecorder {
	return m.recorder
}

// AlertConfiguration mocks base method.
func (m *MockAlertConfigurationDescriber) AlertConfiguration(arg0, arg1 string) (*admin.GroupAlertsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.GroupAlertsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertConfiguration indicates an expected call of AlertConfiguration.
func (mr *MockAlertConfigurationDescriberMockRecorder) AlertConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertConfiguration", reflect.TypeOf((*MockAlertConfigurationDescriber)(nil).AlertConfiguration), arg0, arg1)
}

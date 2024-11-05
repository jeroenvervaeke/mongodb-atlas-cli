// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: LogsDownloader)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

// MockLogsDownloader is a mock of LogsDownloader interface.
type MockLogsDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockLogsDownloaderMockRecorder
}

// MockLogsDownloaderMockRecorder is the mock recorder for MockLogsDownloader.
type MockLogsDownloaderMockRecorder struct {
	mock *MockLogsDownloader
}

// NewMockLogsDownloader creates a new mock instance.
func NewMockLogsDownloader(ctrl *gomock.Controller) *MockLogsDownloader {
	mock := &MockLogsDownloader{ctrl: ctrl}
	mock.recorder = &MockLogsDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogsDownloader) EXPECT() *MockLogsDownloaderMockRecorder {
	return m.recorder
}

// DownloadLog mocks base method.
func (m *MockLogsDownloader) DownloadLog(arg0 *admin.GetHostLogsApiParams) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadLog", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadLog indicates an expected call of DownloadLog.
func (mr *MockLogsDownloaderMockRecorder) DownloadLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadLog", reflect.TypeOf((*MockLogsDownloader)(nil).DownloadLog), arg0)
}

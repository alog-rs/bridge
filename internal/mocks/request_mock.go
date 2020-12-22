// Code generated by MockGen. DO NOT EDIT.
// Source: /home/eindex/go/src/github.com/alog-rs/bridge/scripts/../internal/helpers/request.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHTTPRequest is a mock of HTTPRequest interface
type MockHTTPRequest struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRequestMockRecorder
}

// MockHTTPRequestMockRecorder is the mock recorder for MockHTTPRequest
type MockHTTPRequestMockRecorder struct {
	mock *MockHTTPRequest
}

// NewMockHTTPRequest creates a new mock instance
func NewMockHTTPRequest(ctrl *gomock.Controller) *MockHTTPRequest {
	mock := &MockHTTPRequest{ctrl: ctrl}
	mock.recorder = &MockHTTPRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRequest) EXPECT() *MockHTTPRequestMockRecorder {
	return m.recorder
}

// GetRuneMetricsProfile mocks base method
func (m *MockHTTPRequest) GetRuneMetricsProfile(arg0 string, arg1 int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuneMetricsProfile", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuneMetricsProfile indicates an expected call of GetRuneMetricsProfile
func (mr *MockHTTPRequestMockRecorder) GetRuneMetricsProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuneMetricsProfile", reflect.TypeOf((*MockHTTPRequest)(nil).GetRuneMetricsProfile), arg0, arg1)
}
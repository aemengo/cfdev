// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/provision (interfaces: BoshRunner)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoshRunner is a mock of BoshRunner interface
type MockBoshRunner struct {
	ctrl     *gomock.Controller
	recorder *MockBoshRunnerMockRecorder
}

// MockBoshRunnerMockRecorder is the mock recorder for MockBoshRunner
type MockBoshRunnerMockRecorder struct {
	mock *MockBoshRunner
}

// NewMockBoshRunner creates a new mock instance
func NewMockBoshRunner(ctrl *gomock.Controller) *MockBoshRunner {
	mock := &MockBoshRunner{ctrl: ctrl}
	mock.recorder = &MockBoshRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoshRunner) EXPECT() *MockBoshRunnerMockRecorder {
	return m.recorder
}

// Output mocks base method
func (m *MockBoshRunner) Output(arg0 ...string) ([]byte, error) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Output", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output
func (mr *MockBoshRunnerMockRecorder) Output(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockBoshRunner)(nil).Output), arg0...)
}

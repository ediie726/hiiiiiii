// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/manifest/loader.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ec2 "github.com/aws/copilot-cli/internal/pkg/aws/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MocksubnetIDsGetter is a mock of subnetIDsGetter interface.
type MocksubnetIDsGetter struct {
	ctrl     *gomock.Controller
	recorder *MocksubnetIDsGetterMockRecorder
}

// MocksubnetIDsGetterMockRecorder is the mock recorder for MocksubnetIDsGetter.
type MocksubnetIDsGetterMockRecorder struct {
	mock *MocksubnetIDsGetter
}

// NewMocksubnetIDsGetter creates a new mock instance.
func NewMocksubnetIDsGetter(ctrl *gomock.Controller) *MocksubnetIDsGetter {
	mock := &MocksubnetIDsGetter{ctrl: ctrl}
	mock.recorder = &MocksubnetIDsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksubnetIDsGetter) EXPECT() *MocksubnetIDsGetterMockRecorder {
	return m.recorder
}

// SubnetIDs mocks base method.
func (m *MocksubnetIDsGetter) SubnetIDs(filters ...ec2.Filter) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubnetIDs", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubnetIDs indicates an expected call of SubnetIDs.
func (mr *MocksubnetIDsGetterMockRecorder) SubnetIDs(filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubnetIDs", reflect.TypeOf((*MocksubnetIDsGetter)(nil).SubnetIDs), filters...)
}

// Mockloader is a mock of loader interface.
type Mockloader struct {
	ctrl     *gomock.Controller
	recorder *MockloaderMockRecorder
}

// MockloaderMockRecorder is the mock recorder for Mockloader.
type MockloaderMockRecorder struct {
	mock *Mockloader
}

// NewMockloader creates a new mock instance.
func NewMockloader(ctrl *gomock.Controller) *Mockloader {
	mock := &Mockloader{ctrl: ctrl}
	mock.recorder = &MockloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockloader) EXPECT() *MockloaderMockRecorder {
	return m.recorder
}

// load mocks base method.
func (m *Mockloader) load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "load")
	ret0, _ := ret[0].(error)
	return ret0
}

// load indicates an expected call of load.
func (mr *MockloaderMockRecorder) load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "load", reflect.TypeOf((*Mockloader)(nil).load))
}

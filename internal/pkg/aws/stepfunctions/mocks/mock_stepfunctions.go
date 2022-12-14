// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/stepfunctions/stepfunctions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	sfn "github.com/aws/aws-sdk-go/service/sfn"
	gomock "github.com/golang/mock/gomock"
)

// Mockapi is a mock of api interface.
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi.
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance.
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// DescribeStateMachine mocks base method.
func (m *Mockapi) DescribeStateMachine(input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStateMachine", input)
	ret0, _ := ret[0].(*sfn.DescribeStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStateMachine indicates an expected call of DescribeStateMachine.
func (mr *MockapiMockRecorder) DescribeStateMachine(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStateMachine", reflect.TypeOf((*Mockapi)(nil).DescribeStateMachine), input)
}

// StartExecution mocks base method.
func (m *Mockapi) StartExecution(input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartExecution", input)
	ret0, _ := ret[0].(*sfn.StartExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartExecution indicates an expected call of StartExecution.
func (mr *MockapiMockRecorder) StartExecution(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartExecution", reflect.TypeOf((*Mockapi)(nil).StartExecution), input)
}

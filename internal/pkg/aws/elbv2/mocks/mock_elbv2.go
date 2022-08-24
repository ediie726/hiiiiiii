// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/elbv2/elbv2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
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

// DescribeRules mocks base method.
func (m *Mockapi) DescribeRules(input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRules", input)
	ret0, _ := ret[0].(*elbv2.DescribeRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRules indicates an expected call of DescribeRules.
func (mr *MockapiMockRecorder) DescribeRules(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRules", reflect.TypeOf((*Mockapi)(nil).DescribeRules), input)
}

// DescribeTargetHealth mocks base method.
func (m *Mockapi) DescribeTargetHealth(input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetHealth", input)
	ret0, _ := ret[0].(*elbv2.DescribeTargetHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetHealth indicates an expected call of DescribeTargetHealth.
func (mr *MockapiMockRecorder) DescribeTargetHealth(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetHealth", reflect.TypeOf((*Mockapi)(nil).DescribeTargetHealth), input)
}

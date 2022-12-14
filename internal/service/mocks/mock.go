// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/onemgvv/wb-l0/internal/domain"
)

// MockOrders is a mock of Orders interface.
type MockOrders struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersMockRecorder
}

// MockOrdersMockRecorder is the mock recorder for MockOrders.
type MockOrdersMockRecorder struct {
	mock *MockOrders
}

// NewMockOrders creates a new mock instance.
func NewMockOrders(ctrl *gomock.Controller) *MockOrders {
	mock := &MockOrders{ctrl: ctrl}
	mock.recorder = &MockOrdersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrders) EXPECT() *MockOrdersMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockOrders) GetById(uid string) (domain.OrderJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", uid)
	ret0, _ := ret[0].(domain.OrderJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOrdersMockRecorder) GetById(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOrders)(nil).GetById), uid)
}

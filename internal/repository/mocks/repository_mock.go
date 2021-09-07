// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "ova-route-api/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddRoute mocks base method.
func (m *MockRepo) AddRoute(route models.Route) (models.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", route)
	ret0, _ := ret[0].(models.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockRepoMockRecorder) AddRoute(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockRepo)(nil).AddRoute), route)
}

// AddRoutes mocks base method.
func (m *MockRepo) AddRoutes(routes []models.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoutes", routes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoutes indicates an expected call of AddRoutes.
func (mr *MockRepoMockRecorder) AddRoutes(routes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoutes", reflect.TypeOf((*MockRepo)(nil).AddRoutes), routes)
}

// DescribeRoute mocks base method.
func (m *MockRepo) DescribeRoute(route models.Route) (models.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRoute", route)
	ret0, _ := ret[0].(models.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRoute indicates an expected call of DescribeRoute.
func (mr *MockRepoMockRecorder) DescribeRoute(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRoute", reflect.TypeOf((*MockRepo)(nil).DescribeRoute), route)
}

// ListRoutes mocks base method.
func (m *MockRepo) ListRoutes(limit, offset uint64) ([]models.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoutes", limit, offset)
	ret0, _ := ret[0].([]models.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoutes indicates an expected call of ListRoutes.
func (mr *MockRepoMockRecorder) ListRoutes(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoutes", reflect.TypeOf((*MockRepo)(nil).ListRoutes), limit, offset)
}

// RemoveRoute mocks base method.
func (m *MockRepo) RemoveRoute(routeID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRoute", routeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRoute indicates an expected call of RemoveRoute.
func (mr *MockRepoMockRecorder) RemoveRoute(routeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRoute", reflect.TypeOf((*MockRepo)(nil).RemoveRoute), routeID)
}

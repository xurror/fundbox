// Code generated by MockGen. DO NOT EDIT.
// Source: service/user_service.go
//
// Generated by this command:
//
//	mockgen -source=service/user_service.go -destination=service/mocks/mock_user_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	model "getting-to-go/model"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(ctx context.Context, auth0Id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, auth0Id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, auth0Id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), ctx, auth0Id)
}

// GetUserByAuth0Id mocks base method.
func (m *MockUserService) GetUserByAuth0Id(ctx context.Context, auth0Id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuth0Id", ctx, auth0Id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuth0Id indicates an expected call of GetUserByAuth0Id.
func (mr *MockUserServiceMockRecorder) GetUserByAuth0Id(ctx, auth0Id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuth0Id", reflect.TypeOf((*MockUserService)(nil).GetUserByAuth0Id), ctx, auth0Id)
}

// GetUserById mocks base method.
func (m *MockUserService) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceMockRecorder) GetUserById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserService)(nil).GetUserById), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(ctx context.Context, limit int, offset any) ([]*model.User, any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(any)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(ctx, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), ctx, limit, offset)
}

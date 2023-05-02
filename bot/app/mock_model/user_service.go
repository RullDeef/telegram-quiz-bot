// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/user_service.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
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

// ChangeUsername mocks base method.
func (m *MockUserService) ChangeUsername(username, telegramId string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUsername", username, telegramId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ChangeUsername indicates an expected call of ChangeUsername.
func (mr *MockUserServiceMockRecorder) ChangeUsername(username, telegramId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUsername", reflect.TypeOf((*MockUserService)(nil).ChangeUsername), username, telegramId)
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(username, telegramId string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", username, telegramId)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(username, telegramId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), username, telegramId)
}

// GetUserByTelegramId mocks base method.
func (m *MockUserService) GetUserByTelegramId(id string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByTelegramId", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByTelegramId indicates an expected call of GetUserByTelegramId.
func (mr *MockUserServiceMockRecorder) GetUserByTelegramId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByTelegramId", reflect.TypeOf((*MockUserService)(nil).GetUserByTelegramId), id)
}

// SetUserRole mocks base method.
func (m *MockUserService) SetUserRole(role, telegramId string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserRole", role, telegramId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetUserRole indicates an expected call of SetUserRole.
func (mr *MockUserServiceMockRecorder) SetUserRole(role, telegramId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserRole", reflect.TypeOf((*MockUserService)(nil).SetUserRole), role, telegramId)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/quiz_service.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/RullDeef/telegram-quiz-bot/model"
	gomock "github.com/golang/mock/gomock"
)

// MockQuizService is a mock of QuizService interface.
type MockQuizService struct {
	ctrl     *gomock.Controller
	recorder *MockQuizServiceMockRecorder
}

// MockQuizServiceMockRecorder is the mock recorder for MockQuizService.
type MockQuizServiceMockRecorder struct {
	mock *MockQuizService
}

// NewMockQuizService creates a new mock instance.
func NewMockQuizService(ctrl *gomock.Controller) *MockQuizService {
	mock := &MockQuizService{ctrl: ctrl}
	mock.recorder = &MockQuizServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuizService) EXPECT() *MockQuizServiceMockRecorder {
	return m.recorder
}

// AddAnswer mocks base method.
func (m *MockQuizService) AddAnswer(questionID int64, answer string, isCorrect bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAnswer", questionID, answer, isCorrect)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAnswer indicates an expected call of AddAnswer.
func (mr *MockQuizServiceMockRecorder) AddAnswer(questionID, answer, isCorrect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAnswer", reflect.TypeOf((*MockQuizService)(nil).AddAnswer), questionID, answer, isCorrect)
}

// AddQuestionToTopic mocks base method.
func (m *MockQuizService) AddQuestionToTopic(topic, question string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddQuestionToTopic", topic, question)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddQuestionToTopic indicates an expected call of AddQuestionToTopic.
func (mr *MockQuizServiceMockRecorder) AddQuestionToTopic(topic, question interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQuestionToTopic", reflect.TypeOf((*MockQuizService)(nil).AddQuestionToTopic), topic, question)
}

// CreateQuiz mocks base method.
func (m *MockQuizService) CreateQuiz(topic string) (model.Quiz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuiz", topic)
	ret0, _ := ret[0].(model.Quiz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuiz indicates an expected call of CreateQuiz.
func (mr *MockQuizServiceMockRecorder) CreateQuiz(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuiz", reflect.TypeOf((*MockQuizService)(nil).CreateQuiz), topic)
}

// SetNumQuestionsInQuiz mocks base method.
func (m *MockQuizService) SetNumQuestionsInQuiz(number int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNumQuestionsInQuiz", number)
}

// SetNumQuestionsInQuiz indicates an expected call of SetNumQuestionsInQuiz.
func (mr *MockQuizServiceMockRecorder) SetNumQuestionsInQuiz(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNumQuestionsInQuiz", reflect.TypeOf((*MockQuizService)(nil).SetNumQuestionsInQuiz), number)
}

// ViewQuestionsByTopic mocks base method.
func (m *MockQuizService) ViewQuestionsByTopic(topic string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewQuestionsByTopic", topic)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewQuestionsByTopic indicates an expected call of ViewQuestionsByTopic.
func (mr *MockQuizServiceMockRecorder) ViewQuestionsByTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewQuestionsByTopic", reflect.TypeOf((*MockQuizService)(nil).ViewQuestionsByTopic), topic)
}

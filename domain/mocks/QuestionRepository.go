// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/RullDeef/telegram-quiz-bot/domain"
	mock "github.com/stretchr/testify/mock"
)

// QuestionRepository is an autogenerated mock type for the QuestionRepository type
type QuestionRepository struct {
	mock.Mock
}

// GetById provides a mock function with given fields: ctx, id
func (_m *QuestionRepository) GetById(ctx context.Context, id int64) (domain.Question, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Question
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Question, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Question); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Question)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQuestionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewQuestionRepository creates a new instance of QuestionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuestionRepository(t mockConstructorTestingTNewQuestionRepository) *QuestionRepository {
	mock := &QuestionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

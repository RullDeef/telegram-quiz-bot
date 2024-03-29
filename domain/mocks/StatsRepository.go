// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/RullDeef/telegram-quiz-bot/domain"
	mock "github.com/stretchr/testify/mock"
)

// StatsRepository is an autogenerated mock type for the StatsRepository type
type StatsRepository struct {
	mock.Mock
}

// GetById provides a mock function with given fields: ctx, id
func (_m *StatsRepository) GetById(ctx context.Context, id int64) (domain.Stats, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Stats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Stats, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Stats); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Stats)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStatsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatsRepository creates a new instance of StatsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatsRepository(t mockConstructorTestingTNewStatsRepository) *StatsRepository {
	mock := &StatsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/BrandonReno/A.E.R/models"

// AthleteRepository is an autogenerated mock type for the AthleteRepository type
type AthleteRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, a
func (_m *AthleteRepository) Create(ctx context.Context, a *models.Athlete) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Athlete) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *AthleteRepository) GetAll(ctx context.Context) ([]*models.Athlete, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Athlete
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Athlete); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Athlete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, aid
func (_m *AthleteRepository) GetByID(ctx context.Context, aid int) (*models.Athlete, error) {
	ret := _m.Called(ctx, aid)

	var r0 *models.Athlete
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Athlete); ok {
		r0 = rf(ctx, aid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Athlete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, aid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, a
func (_m *AthleteRepository) Update(ctx context.Context, a *models.Athlete) (*models.Athlete, error) {
	ret := _m.Called(ctx, a)

	var r0 *models.Athlete
	if rf, ok := ret.Get(0).(func(context.Context, *models.Athlete) *models.Athlete); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Athlete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Athlete) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
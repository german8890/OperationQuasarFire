// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/OperationQuasarFire/internal/pkg/entity"

	mock "github.com/stretchr/testify/mock"
)

// DBRepository is an autogenerated mock type for the DBRepository type
type DBRepository struct {
	mock.Mock
}

// CreateTaskTotal provides a mock function with given fields: request
func (_m *DBRepository) CreateTaskTotal(request *entity.SatelliteRequest) (string, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for CreateTaskTotal")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.SatelliteRequest) (string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*entity.SatelliteRequest) string); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*entity.SatelliteRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDBRepository creates a new instance of DBRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBRepository {
	mock := &DBRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
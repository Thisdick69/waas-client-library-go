// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	iterator "google.golang.org/api/iterator"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// PoolIterator is an autogenerated mock type for the PoolIterator type
type PoolIterator struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *PoolIterator) Next() (*v1.Pool, error) {
	ret := _m.Called()

	var r0 *v1.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.Pool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pool)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PageInfo provides a mock function with given fields:
func (_m *PoolIterator) PageInfo() *iterator.PageInfo {
	ret := _m.Called()

	var r0 *iterator.PageInfo
	if rf, ok := ret.Get(0).(func() *iterator.PageInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iterator.PageInfo)
		}
	}

	return r0
}

// Response provides a mock function with given fields:
func (_m *PoolIterator) Response() *v1.ListPoolsResponse {
	ret := _m.Called()

	var r0 *v1.ListPoolsResponse
	if rf, ok := ret.Get(0).(func() *v1.ListPoolsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListPoolsResponse)
		}
	}

	return r0
}

// NewPoolIterator creates a new instance of PoolIterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPoolIterator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PoolIterator {
	mock := &PoolIterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
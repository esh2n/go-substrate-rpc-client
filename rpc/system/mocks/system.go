// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/esh2n/go-substrate-rpc-client/v4/types"
)

// System is an autogenerated mock type for the System type
type System struct {
	mock.Mock
}

// Chain provides a mock function with given fields:
func (_m *System) Chain() (types.Text, error) {
	ret := _m.Called()

	var r0 types.Text
	if rf, ok := ret.Get(0).(func() types.Text); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Text)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields:
func (_m *System) Health() (types.Health, error) {
	ret := _m.Called()

	var r0 types.Health
	if rf, ok := ret.Get(0).(func() types.Health); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Health)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *System) Name() (types.Text, error) {
	ret := _m.Called()

	var r0 types.Text
	if rf, ok := ret.Get(0).(func() types.Text); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Text)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkState provides a mock function with given fields:
func (_m *System) NetworkState() (types.NetworkState, error) {
	ret := _m.Called()

	var r0 types.NetworkState
	if rf, ok := ret.Get(0).(func() types.NetworkState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NetworkState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peers provides a mock function with given fields:
func (_m *System) Peers() ([]types.PeerInfo, error) {
	ret := _m.Called()

	var r0 []types.PeerInfo
	if rf, ok := ret.Get(0).(func() []types.PeerInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.PeerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Properties provides a mock function with given fields:
func (_m *System) Properties() (types.ChainProperties, error) {
	ret := _m.Called()

	var r0 types.ChainProperties
	if rf, ok := ret.Get(0).(func() types.ChainProperties); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.ChainProperties)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields:
func (_m *System) Version() (types.Text, error) {
	ret := _m.Called()

	var r0 types.Text
	if rf, ok := ret.Get(0).(func() types.Text); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Text)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewSystemT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSystem creates a new instance of System. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSystem(t NewSystemT) *System {
	mock := &System{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

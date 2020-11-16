// Code generated by mockery v2.0.3. DO NOT EDIT.

package fuse

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FuseInstaller is an autogenerated mock type for the FuseInstaller type
type FuseInstaller struct {
	mock.Mock
}

// Install provides a mock function with given fields: ctx, args
func (_m *FuseInstaller) Install(ctx context.Context, args map[string]interface{}) error {
	ret := _m.Called(ctx, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) error); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsInstalled provides a mock function with given fields: ctx
func (_m *FuseInstaller) IsInstalled(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

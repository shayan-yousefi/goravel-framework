// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	foundation "github.com/goravel/framework/contracts/foundation"
	mock "github.com/stretchr/testify/mock"
)

// ServiceProvider is an autogenerated mock type for the ServiceProvider type
type ServiceProvider struct {
	mock.Mock
}

type ServiceProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceProvider) EXPECT() *ServiceProvider_Expecter {
	return &ServiceProvider_Expecter{mock: &_m.Mock}
}

// Boot provides a mock function with given fields: app
func (_m *ServiceProvider) Boot(app foundation.Application) {
	_m.Called(app)
}

// ServiceProvider_Boot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Boot'
type ServiceProvider_Boot_Call struct {
	*mock.Call
}

// Boot is a helper method to define mock.On call
//   - app foundation.Application
func (_e *ServiceProvider_Expecter) Boot(app interface{}) *ServiceProvider_Boot_Call {
	return &ServiceProvider_Boot_Call{Call: _e.mock.On("Boot", app)}
}

func (_c *ServiceProvider_Boot_Call) Run(run func(app foundation.Application)) *ServiceProvider_Boot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(foundation.Application))
	})
	return _c
}

func (_c *ServiceProvider_Boot_Call) Return() *ServiceProvider_Boot_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceProvider_Boot_Call) RunAndReturn(run func(foundation.Application)) *ServiceProvider_Boot_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: app
func (_m *ServiceProvider) Register(app foundation.Application) {
	_m.Called(app)
}

// ServiceProvider_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type ServiceProvider_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - app foundation.Application
func (_e *ServiceProvider_Expecter) Register(app interface{}) *ServiceProvider_Register_Call {
	return &ServiceProvider_Register_Call{Call: _e.mock.On("Register", app)}
}

func (_c *ServiceProvider_Register_Call) Run(run func(app foundation.Application)) *ServiceProvider_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(foundation.Application))
	})
	return _c
}

func (_c *ServiceProvider_Register_Call) Return() *ServiceProvider_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceProvider_Register_Call) RunAndReturn(run func(foundation.Application)) *ServiceProvider_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceProvider creates a new instance of ServiceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceProvider {
	mock := &ServiceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	schedule "github.com/goravel/framework/contracts/schedule"
	mock "github.com/stretchr/testify/mock"
)

// Schedule is an autogenerated mock type for the Schedule type
type Schedule struct {
	mock.Mock
}

type Schedule_Expecter struct {
	mock *mock.Mock
}

func (_m *Schedule) EXPECT() *Schedule_Expecter {
	return &Schedule_Expecter{mock: &_m.Mock}
}

// Call provides a mock function with given fields: callback
func (_m *Schedule) Call(callback func()) schedule.Event {
	ret := _m.Called(callback)

	if len(ret) == 0 {
		panic("no return value specified for Call")
	}

	var r0 schedule.Event
	if rf, ok := ret.Get(0).(func(func()) schedule.Event); ok {
		r0 = rf(callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schedule.Event)
		}
	}

	return r0
}

// Schedule_Call_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Call'
type Schedule_Call_Call struct {
	*mock.Call
}

// Call is a helper method to define mock.On call
//   - callback func()
func (_e *Schedule_Expecter) Call(callback interface{}) *Schedule_Call_Call {
	return &Schedule_Call_Call{Call: _e.mock.On("Call", callback)}
}

func (_c *Schedule_Call_Call) Run(run func(callback func())) *Schedule_Call_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func()))
	})
	return _c
}

func (_c *Schedule_Call_Call) Return(_a0 schedule.Event) *Schedule_Call_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schedule_Call_Call) RunAndReturn(run func(func()) schedule.Event) *Schedule_Call_Call {
	_c.Call.Return(run)
	return _c
}

// Command provides a mock function with given fields: command
func (_m *Schedule) Command(command string) schedule.Event {
	ret := _m.Called(command)

	if len(ret) == 0 {
		panic("no return value specified for Command")
	}

	var r0 schedule.Event
	if rf, ok := ret.Get(0).(func(string) schedule.Event); ok {
		r0 = rf(command)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schedule.Event)
		}
	}

	return r0
}

// Schedule_Command_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Command'
type Schedule_Command_Call struct {
	*mock.Call
}

// Command is a helper method to define mock.On call
//   - command string
func (_e *Schedule_Expecter) Command(command interface{}) *Schedule_Command_Call {
	return &Schedule_Command_Call{Call: _e.mock.On("Command", command)}
}

func (_c *Schedule_Command_Call) Run(run func(command string)) *Schedule_Command_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schedule_Command_Call) Return(_a0 schedule.Event) *Schedule_Command_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schedule_Command_Call) RunAndReturn(run func(string) schedule.Event) *Schedule_Command_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: events
func (_m *Schedule) Register(events []schedule.Event) {
	_m.Called(events)
}

// Schedule_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Schedule_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - events []schedule.Event
func (_e *Schedule_Expecter) Register(events interface{}) *Schedule_Register_Call {
	return &Schedule_Register_Call{Call: _e.mock.On("Register", events)}
}

func (_c *Schedule_Register_Call) Run(run func(events []schedule.Event)) *Schedule_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]schedule.Event))
	})
	return _c
}

func (_c *Schedule_Register_Call) Return() *Schedule_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schedule_Register_Call) RunAndReturn(run func([]schedule.Event)) *Schedule_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *Schedule) Run() {
	_m.Called()
}

// Schedule_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Schedule_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *Schedule_Expecter) Run() *Schedule_Run_Call {
	return &Schedule_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *Schedule_Run_Call) Run(run func()) *Schedule_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schedule_Run_Call) Return() *Schedule_Run_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schedule_Run_Call) RunAndReturn(run func()) *Schedule_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewSchedule creates a new instance of Schedule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSchedule(t interface {
	mock.TestingT
	Cleanup(func())
}) *Schedule {
	mock := &Schedule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

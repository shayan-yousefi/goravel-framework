// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/goravel/framework/contracts/log"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Entry is an autogenerated mock type for the Entry type
type Entry struct {
	mock.Mock
}

type Entry_Expecter struct {
	mock *mock.Mock
}

func (_m *Entry) EXPECT() *Entry_Expecter {
	return &Entry_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *Entry) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Entry_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type Entry_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *Entry_Expecter) Context() *Entry_Context_Call {
	return &Entry_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *Entry_Context_Call) Run(run func()) *Entry_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Context_Call) Return(_a0 context.Context) *Entry_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Context_Call) RunAndReturn(run func() context.Context) *Entry_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Level provides a mock function with given fields:
func (_m *Entry) Level() log.Level {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Level")
	}

	var r0 log.Level
	if rf, ok := ret.Get(0).(func() log.Level); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(log.Level)
	}

	return r0
}

// Entry_Level_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Level'
type Entry_Level_Call struct {
	*mock.Call
}

// Level is a helper method to define mock.On call
func (_e *Entry_Expecter) Level() *Entry_Level_Call {
	return &Entry_Level_Call{Call: _e.mock.On("Level")}
}

func (_c *Entry_Level_Call) Run(run func()) *Entry_Level_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Level_Call) Return(_a0 log.Level) *Entry_Level_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Level_Call) RunAndReturn(run func() log.Level) *Entry_Level_Call {
	_c.Call.Return(run)
	return _c
}

// Message provides a mock function with given fields:
func (_m *Entry) Message() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Message")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Entry_Message_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Message'
type Entry_Message_Call struct {
	*mock.Call
}

// Message is a helper method to define mock.On call
func (_e *Entry_Expecter) Message() *Entry_Message_Call {
	return &Entry_Message_Call{Call: _e.mock.On("Message")}
}

func (_c *Entry_Message_Call) Run(run func()) *Entry_Message_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Message_Call) Return(_a0 string) *Entry_Message_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Message_Call) RunAndReturn(run func() string) *Entry_Message_Call {
	_c.Call.Return(run)
	return _c
}

// Time provides a mock function with given fields:
func (_m *Entry) Time() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Time")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Entry_Time_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Time'
type Entry_Time_Call struct {
	*mock.Call
}

// Time is a helper method to define mock.On call
func (_e *Entry_Expecter) Time() *Entry_Time_Call {
	return &Entry_Time_Call{Call: _e.mock.On("Time")}
}

func (_c *Entry_Time_Call) Run(run func()) *Entry_Time_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Entry_Time_Call) Return(_a0 time.Time) *Entry_Time_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Entry_Time_Call) RunAndReturn(run func() time.Time) *Entry_Time_Call {
	_c.Call.Return(run)
	return _c
}

// NewEntry creates a new instance of Entry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEntry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Entry {
	mock := &Entry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

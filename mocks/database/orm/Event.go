// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// Event is an autogenerated mock type for the Event type
type Event struct {
	mock.Mock
}

type Event_Expecter struct {
	mock *mock.Mock
}

func (_m *Event) EXPECT() *Event_Expecter {
	return &Event_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *Event) Context() context.Context {
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

// Event_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type Event_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *Event_Expecter) Context() *Event_Context_Call {
	return &Event_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *Event_Context_Call) Run(run func()) *Event_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Event_Context_Call) Return(_a0 context.Context) *Event_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_Context_Call) RunAndReturn(run func() context.Context) *Event_Context_Call {
	_c.Call.Return(run)
	return _c
}

// GetAttribute provides a mock function with given fields: key
func (_m *Event) GetAttribute(key string) interface{} {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetAttribute")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Event_GetAttribute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAttribute'
type Event_GetAttribute_Call struct {
	*mock.Call
}

// GetAttribute is a helper method to define mock.On call
//   - key string
func (_e *Event_Expecter) GetAttribute(key interface{}) *Event_GetAttribute_Call {
	return &Event_GetAttribute_Call{Call: _e.mock.On("GetAttribute", key)}
}

func (_c *Event_GetAttribute_Call) Run(run func(key string)) *Event_GetAttribute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Event_GetAttribute_Call) Return(_a0 interface{}) *Event_GetAttribute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_GetAttribute_Call) RunAndReturn(run func(string) interface{}) *Event_GetAttribute_Call {
	_c.Call.Return(run)
	return _c
}

// GetOriginal provides a mock function with given fields: key, def
func (_m *Event) GetOriginal(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOriginal")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, def...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Event_GetOriginal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOriginal'
type Event_GetOriginal_Call struct {
	*mock.Call
}

// GetOriginal is a helper method to define mock.On call
//   - key string
//   - def ...interface{}
func (_e *Event_Expecter) GetOriginal(key interface{}, def ...interface{}) *Event_GetOriginal_Call {
	return &Event_GetOriginal_Call{Call: _e.mock.On("GetOriginal",
		append([]interface{}{key}, def...)...)}
}

func (_c *Event_GetOriginal_Call) Run(run func(key string, def ...interface{})) *Event_GetOriginal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Event_GetOriginal_Call) Return(_a0 interface{}) *Event_GetOriginal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_GetOriginal_Call) RunAndReturn(run func(string, ...interface{}) interface{}) *Event_GetOriginal_Call {
	_c.Call.Return(run)
	return _c
}

// IsClean provides a mock function with given fields: columns
func (_m *Event) IsClean(columns ...string) bool {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsClean")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(columns...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Event_IsClean_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsClean'
type Event_IsClean_Call struct {
	*mock.Call
}

// IsClean is a helper method to define mock.On call
//   - columns ...string
func (_e *Event_Expecter) IsClean(columns ...interface{}) *Event_IsClean_Call {
	return &Event_IsClean_Call{Call: _e.mock.On("IsClean",
		append([]interface{}{}, columns...)...)}
}

func (_c *Event_IsClean_Call) Run(run func(columns ...string)) *Event_IsClean_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Event_IsClean_Call) Return(_a0 bool) *Event_IsClean_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_IsClean_Call) RunAndReturn(run func(...string) bool) *Event_IsClean_Call {
	_c.Call.Return(run)
	return _c
}

// IsDirty provides a mock function with given fields: columns
func (_m *Event) IsDirty(columns ...string) bool {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsDirty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(columns...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Event_IsDirty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDirty'
type Event_IsDirty_Call struct {
	*mock.Call
}

// IsDirty is a helper method to define mock.On call
//   - columns ...string
func (_e *Event_Expecter) IsDirty(columns ...interface{}) *Event_IsDirty_Call {
	return &Event_IsDirty_Call{Call: _e.mock.On("IsDirty",
		append([]interface{}{}, columns...)...)}
}

func (_c *Event_IsDirty_Call) Run(run func(columns ...string)) *Event_IsDirty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Event_IsDirty_Call) Return(_a0 bool) *Event_IsDirty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_IsDirty_Call) RunAndReturn(run func(...string) bool) *Event_IsDirty_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields:
func (_m *Event) Query() orm.Query {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Event_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type Event_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
func (_e *Event_Expecter) Query() *Event_Query_Call {
	return &Event_Query_Call{Call: _e.mock.On("Query")}
}

func (_c *Event_Query_Call) Run(run func()) *Event_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Event_Query_Call) Return(_a0 orm.Query) *Event_Query_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Event_Query_Call) RunAndReturn(run func() orm.Query) *Event_Query_Call {
	_c.Call.Return(run)
	return _c
}

// SetAttribute provides a mock function with given fields: key, value
func (_m *Event) SetAttribute(key string, value interface{}) {
	_m.Called(key, value)
}

// Event_SetAttribute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAttribute'
type Event_SetAttribute_Call struct {
	*mock.Call
}

// SetAttribute is a helper method to define mock.On call
//   - key string
//   - value interface{}
func (_e *Event_Expecter) SetAttribute(key interface{}, value interface{}) *Event_SetAttribute_Call {
	return &Event_SetAttribute_Call{Call: _e.mock.On("SetAttribute", key, value)}
}

func (_c *Event_SetAttribute_Call) Run(run func(key string, value interface{})) *Event_SetAttribute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *Event_SetAttribute_Call) Return() *Event_SetAttribute_Call {
	_c.Call.Return()
	return _c
}

func (_c *Event_SetAttribute_Call) RunAndReturn(run func(string, interface{})) *Event_SetAttribute_Call {
	_c.Call.Return(run)
	return _c
}

// NewEvent creates a new instance of Event. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *Event {
	mock := &Event{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

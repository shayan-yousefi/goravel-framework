// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/goravel/framework/contracts/cache"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

type Driver_Expecter struct {
	mock *mock.Mock
}

func (_m *Driver) EXPECT() *Driver_Expecter {
	return &Driver_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: key, value, t
func (_m *Driver) Add(key string, value interface{}, t time.Duration) bool {
	ret := _m.Called(key, value, t)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) bool); ok {
		r0 = rf(key, value, t)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Driver_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - key string
//   - value interface{}
//   - t time.Duration
func (_e *Driver_Expecter) Add(key interface{}, value interface{}, t interface{}) *Driver_Add_Call {
	return &Driver_Add_Call{Call: _e.mock.On("Add", key, value, t)}
}

func (_c *Driver_Add_Call) Run(run func(key string, value interface{}, t time.Duration)) *Driver_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}), args[2].(time.Duration))
	})
	return _c
}

func (_c *Driver_Add_Call) Return(_a0 bool) *Driver_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Add_Call) RunAndReturn(run func(string, interface{}, time.Duration) bool) *Driver_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Decrement provides a mock function with given fields: key, value
func (_m *Driver) Decrement(key string, value ...int) (int, error) {
	_va := make([]interface{}, len(value))
	for _i := range value {
		_va[_i] = value[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Decrement")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...int) (int, error)); ok {
		return rf(key, value...)
	}
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, value...)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, ...int) error); ok {
		r1 = rf(key, value...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver_Decrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrement'
type Driver_Decrement_Call struct {
	*mock.Call
}

// Decrement is a helper method to define mock.On call
//   - key string
//   - value ...int
func (_e *Driver_Expecter) Decrement(key interface{}, value ...interface{}) *Driver_Decrement_Call {
	return &Driver_Decrement_Call{Call: _e.mock.On("Decrement",
		append([]interface{}{key}, value...)...)}
}

func (_c *Driver_Decrement_Call) Run(run func(key string, value ...int)) *Driver_Decrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_Decrement_Call) Return(_a0 int, _a1 error) *Driver_Decrement_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Driver_Decrement_Call) RunAndReturn(run func(string, ...int) (int, error)) *Driver_Decrement_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields:
func (_m *Driver) Flush() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Flush")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type Driver_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
func (_e *Driver_Expecter) Flush() *Driver_Flush_Call {
	return &Driver_Flush_Call{Call: _e.mock.On("Flush")}
}

func (_c *Driver_Flush_Call) Run(run func()) *Driver_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Driver_Flush_Call) Return(_a0 bool) *Driver_Flush_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Flush_Call) RunAndReturn(run func() bool) *Driver_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Forever provides a mock function with given fields: key, value
func (_m *Driver) Forever(key string, value interface{}) bool {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Forever")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, interface{}) bool); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_Forever_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Forever'
type Driver_Forever_Call struct {
	*mock.Call
}

// Forever is a helper method to define mock.On call
//   - key string
//   - value interface{}
func (_e *Driver_Expecter) Forever(key interface{}, value interface{}) *Driver_Forever_Call {
	return &Driver_Forever_Call{Call: _e.mock.On("Forever", key, value)}
}

func (_c *Driver_Forever_Call) Run(run func(key string, value interface{})) *Driver_Forever_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *Driver_Forever_Call) Return(_a0 bool) *Driver_Forever_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Forever_Call) RunAndReturn(run func(string, interface{}) bool) *Driver_Forever_Call {
	_c.Call.Return(run)
	return _c
}

// Forget provides a mock function with given fields: key
func (_m *Driver) Forget(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Forget")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_Forget_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Forget'
type Driver_Forget_Call struct {
	*mock.Call
}

// Forget is a helper method to define mock.On call
//   - key string
func (_e *Driver_Expecter) Forget(key interface{}) *Driver_Forget_Call {
	return &Driver_Forget_Call{Call: _e.mock.On("Forget", key)}
}

func (_c *Driver_Forget_Call) Run(run func(key string)) *Driver_Forget_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Driver_Forget_Call) Return(_a0 bool) *Driver_Forget_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Forget_Call) RunAndReturn(run func(string) bool) *Driver_Forget_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key, def
func (_m *Driver) Get(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
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

// Driver_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Driver_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
//   - def ...interface{}
func (_e *Driver_Expecter) Get(key interface{}, def ...interface{}) *Driver_Get_Call {
	return &Driver_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_Get_Call) Run(run func(key string, def ...interface{})) *Driver_Get_Call {
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

func (_c *Driver_Get_Call) Return(_a0 interface{}) *Driver_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Get_Call) RunAndReturn(run func(string, ...interface{}) interface{}) *Driver_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetBool provides a mock function with given fields: key, def
func (_m *Driver) GetBool(key string, def ...bool) bool {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBool")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_GetBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBool'
type Driver_GetBool_Call struct {
	*mock.Call
}

// GetBool is a helper method to define mock.On call
//   - key string
//   - def ...bool
func (_e *Driver_Expecter) GetBool(key interface{}, def ...interface{}) *Driver_GetBool_Call {
	return &Driver_GetBool_Call{Call: _e.mock.On("GetBool",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_GetBool_Call) Run(run func(key string, def ...bool)) *Driver_GetBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]bool, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(bool)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_GetBool_Call) Return(_a0 bool) *Driver_GetBool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_GetBool_Call) RunAndReturn(run func(string, ...bool) bool) *Driver_GetBool_Call {
	_c.Call.Return(run)
	return _c
}

// GetInt provides a mock function with given fields: key, def
func (_m *Driver) GetInt(key string, def ...int) int {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInt")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Driver_GetInt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInt'
type Driver_GetInt_Call struct {
	*mock.Call
}

// GetInt is a helper method to define mock.On call
//   - key string
//   - def ...int
func (_e *Driver_Expecter) GetInt(key interface{}, def ...interface{}) *Driver_GetInt_Call {
	return &Driver_GetInt_Call{Call: _e.mock.On("GetInt",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_GetInt_Call) Run(run func(key string, def ...int)) *Driver_GetInt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_GetInt_Call) Return(_a0 int) *Driver_GetInt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_GetInt_Call) RunAndReturn(run func(string, ...int) int) *Driver_GetInt_Call {
	_c.Call.Return(run)
	return _c
}

// GetInt64 provides a mock function with given fields: key, def
func (_m *Driver) GetInt64(key string, def ...int64) int64 {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInt64")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Driver_GetInt64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInt64'
type Driver_GetInt64_Call struct {
	*mock.Call
}

// GetInt64 is a helper method to define mock.On call
//   - key string
//   - def ...int64
func (_e *Driver_Expecter) GetInt64(key interface{}, def ...interface{}) *Driver_GetInt64_Call {
	return &Driver_GetInt64_Call{Call: _e.mock.On("GetInt64",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_GetInt64_Call) Run(run func(key string, def ...int64)) *Driver_GetInt64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_GetInt64_Call) Return(_a0 int64) *Driver_GetInt64_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_GetInt64_Call) RunAndReturn(run func(string, ...int64) int64) *Driver_GetInt64_Call {
	_c.Call.Return(run)
	return _c
}

// GetString provides a mock function with given fields: key, def
func (_m *Driver) GetString(key string, def ...string) string {
	_va := make([]interface{}, len(def))
	for _i := range def {
		_va[_i] = def[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetString")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, def...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Driver_GetString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetString'
type Driver_GetString_Call struct {
	*mock.Call
}

// GetString is a helper method to define mock.On call
//   - key string
//   - def ...string
func (_e *Driver_Expecter) GetString(key interface{}, def ...interface{}) *Driver_GetString_Call {
	return &Driver_GetString_Call{Call: _e.mock.On("GetString",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_GetString_Call) Run(run func(key string, def ...string)) *Driver_GetString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_GetString_Call) Return(_a0 string) *Driver_GetString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_GetString_Call) RunAndReturn(run func(string, ...string) string) *Driver_GetString_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: key
func (_m *Driver) Has(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Driver_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type Driver_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - key string
func (_e *Driver_Expecter) Has(key interface{}) *Driver_Has_Call {
	return &Driver_Has_Call{Call: _e.mock.On("Has", key)}
}

func (_c *Driver_Has_Call) Run(run func(key string)) *Driver_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Driver_Has_Call) Return(_a0 bool) *Driver_Has_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Has_Call) RunAndReturn(run func(string) bool) *Driver_Has_Call {
	_c.Call.Return(run)
	return _c
}

// Increment provides a mock function with given fields: key, value
func (_m *Driver) Increment(key string, value ...int) (int, error) {
	_va := make([]interface{}, len(value))
	for _i := range value {
		_va[_i] = value[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Increment")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...int) (int, error)); ok {
		return rf(key, value...)
	}
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, value...)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, ...int) error); ok {
		r1 = rf(key, value...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver_Increment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Increment'
type Driver_Increment_Call struct {
	*mock.Call
}

// Increment is a helper method to define mock.On call
//   - key string
//   - value ...int
func (_e *Driver_Expecter) Increment(key interface{}, value ...interface{}) *Driver_Increment_Call {
	return &Driver_Increment_Call{Call: _e.mock.On("Increment",
		append([]interface{}{key}, value...)...)}
}

func (_c *Driver_Increment_Call) Run(run func(key string, value ...int)) *Driver_Increment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_Increment_Call) Return(_a0 int, _a1 error) *Driver_Increment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Driver_Increment_Call) RunAndReturn(run func(string, ...int) (int, error)) *Driver_Increment_Call {
	_c.Call.Return(run)
	return _c
}

// Lock provides a mock function with given fields: key, t
func (_m *Driver) Lock(key string, t ...time.Duration) cache.Lock {
	_va := make([]interface{}, len(t))
	for _i := range t {
		_va[_i] = t[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Lock")
	}

	var r0 cache.Lock
	if rf, ok := ret.Get(0).(func(string, ...time.Duration) cache.Lock); ok {
		r0 = rf(key, t...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Lock)
		}
	}

	return r0
}

// Driver_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type Driver_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
//   - key string
//   - t ...time.Duration
func (_e *Driver_Expecter) Lock(key interface{}, t ...interface{}) *Driver_Lock_Call {
	return &Driver_Lock_Call{Call: _e.mock.On("Lock",
		append([]interface{}{key}, t...)...)}
}

func (_c *Driver_Lock_Call) Run(run func(key string, t ...time.Duration)) *Driver_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]time.Duration, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(time.Duration)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Driver_Lock_Call) Return(_a0 cache.Lock) *Driver_Lock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Lock_Call) RunAndReturn(run func(string, ...time.Duration) cache.Lock) *Driver_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// Pull provides a mock function with given fields: key, def
func (_m *Driver) Pull(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Pull")
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

// Driver_Pull_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pull'
type Driver_Pull_Call struct {
	*mock.Call
}

// Pull is a helper method to define mock.On call
//   - key string
//   - def ...interface{}
func (_e *Driver_Expecter) Pull(key interface{}, def ...interface{}) *Driver_Pull_Call {
	return &Driver_Pull_Call{Call: _e.mock.On("Pull",
		append([]interface{}{key}, def...)...)}
}

func (_c *Driver_Pull_Call) Run(run func(key string, def ...interface{})) *Driver_Pull_Call {
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

func (_c *Driver_Pull_Call) Return(_a0 interface{}) *Driver_Pull_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Pull_Call) RunAndReturn(run func(string, ...interface{}) interface{}) *Driver_Pull_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: key, value, t
func (_m *Driver) Put(key string, value interface{}, t time.Duration) error {
	ret := _m.Called(key, value, t)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(key, value, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Driver_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type Driver_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - key string
//   - value interface{}
//   - t time.Duration
func (_e *Driver_Expecter) Put(key interface{}, value interface{}, t interface{}) *Driver_Put_Call {
	return &Driver_Put_Call{Call: _e.mock.On("Put", key, value, t)}
}

func (_c *Driver_Put_Call) Run(run func(key string, value interface{}, t time.Duration)) *Driver_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}), args[2].(time.Duration))
	})
	return _c
}

func (_c *Driver_Put_Call) Return(_a0 error) *Driver_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Put_Call) RunAndReturn(run func(string, interface{}, time.Duration) error) *Driver_Put_Call {
	_c.Call.Return(run)
	return _c
}

// Remember provides a mock function with given fields: key, ttl, callback
func (_m *Driver) Remember(key string, ttl time.Duration, callback func() (interface{}, error)) (interface{}, error) {
	ret := _m.Called(key, ttl, callback)

	if len(ret) == 0 {
		panic("no return value specified for Remember")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Duration, func() (interface{}, error)) (interface{}, error)); ok {
		return rf(key, ttl, callback)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration, func() (interface{}, error)) interface{}); ok {
		r0 = rf(key, ttl, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration, func() (interface{}, error)) error); ok {
		r1 = rf(key, ttl, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver_Remember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remember'
type Driver_Remember_Call struct {
	*mock.Call
}

// Remember is a helper method to define mock.On call
//   - key string
//   - ttl time.Duration
//   - callback func()(interface{} , error)
func (_e *Driver_Expecter) Remember(key interface{}, ttl interface{}, callback interface{}) *Driver_Remember_Call {
	return &Driver_Remember_Call{Call: _e.mock.On("Remember", key, ttl, callback)}
}

func (_c *Driver_Remember_Call) Run(run func(key string, ttl time.Duration, callback func() (interface{}, error))) *Driver_Remember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Duration), args[2].(func() (interface{}, error)))
	})
	return _c
}

func (_c *Driver_Remember_Call) Return(_a0 interface{}, _a1 error) *Driver_Remember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Driver_Remember_Call) RunAndReturn(run func(string, time.Duration, func() (interface{}, error)) (interface{}, error)) *Driver_Remember_Call {
	_c.Call.Return(run)
	return _c
}

// RememberForever provides a mock function with given fields: key, callback
func (_m *Driver) RememberForever(key string, callback func() (interface{}, error)) (interface{}, error) {
	ret := _m.Called(key, callback)

	if len(ret) == 0 {
		panic("no return value specified for RememberForever")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, func() (interface{}, error)) (interface{}, error)); ok {
		return rf(key, callback)
	}
	if rf, ok := ret.Get(0).(func(string, func() (interface{}, error)) interface{}); ok {
		r0 = rf(key, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, func() (interface{}, error)) error); ok {
		r1 = rf(key, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver_RememberForever_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RememberForever'
type Driver_RememberForever_Call struct {
	*mock.Call
}

// RememberForever is a helper method to define mock.On call
//   - key string
//   - callback func()(interface{} , error)
func (_e *Driver_Expecter) RememberForever(key interface{}, callback interface{}) *Driver_RememberForever_Call {
	return &Driver_RememberForever_Call{Call: _e.mock.On("RememberForever", key, callback)}
}

func (_c *Driver_RememberForever_Call) Run(run func(key string, callback func() (interface{}, error))) *Driver_RememberForever_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func() (interface{}, error)))
	})
	return _c
}

func (_c *Driver_RememberForever_Call) Return(_a0 interface{}, _a1 error) *Driver_RememberForever_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Driver_RememberForever_Call) RunAndReturn(run func(string, func() (interface{}, error)) (interface{}, error)) *Driver_RememberForever_Call {
	_c.Call.Return(run)
	return _c
}

// WithContext provides a mock function with given fields: ctx
func (_m *Driver) WithContext(ctx context.Context) cache.Driver {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WithContext")
	}

	var r0 cache.Driver
	if rf, ok := ret.Get(0).(func(context.Context) cache.Driver); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Driver)
		}
	}

	return r0
}

// Driver_WithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithContext'
type Driver_WithContext_Call struct {
	*mock.Call
}

// WithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Driver_Expecter) WithContext(ctx interface{}) *Driver_WithContext_Call {
	return &Driver_WithContext_Call{Call: _e.mock.On("WithContext", ctx)}
}

func (_c *Driver_WithContext_Call) Run(run func(ctx context.Context)) *Driver_WithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Driver_WithContext_Call) Return(_a0 cache.Driver) *Driver_WithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_WithContext_Call) RunAndReturn(run func(context.Context) cache.Driver) *Driver_WithContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewDriver creates a new instance of Driver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDriver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Driver {
	mock := &Driver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

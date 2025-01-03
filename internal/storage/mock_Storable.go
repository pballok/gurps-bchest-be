// Code generated by mockery v2.46.3. DO NOT EDIT.

package storage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockStorable is an autogenerated mock type for the Storable type
type MockStorable[K comparable, V any, F any] struct {
	mock.Mock
}

type MockStorable_Expecter[K comparable, V any, F any] struct {
	mock *mock.Mock
}

func (_m *MockStorable[K, V, F]) EXPECT() *MockStorable_Expecter[K, V, F] {
	return &MockStorable_Expecter[K, V, F]{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, item
func (_m *MockStorable[K, V, F]) Add(ctx context.Context, item V) (K, error) {
	ret := _m.Called(ctx, item)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 K
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, V) (K, error)); ok {
		return rf(ctx, item)
	}
	if rf, ok := ret.Get(0).(func(context.Context, V) K); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Get(0).(K)
	}

	if rf, ok := ret.Get(1).(func(context.Context, V) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorable_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockStorable_Add_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - item V
func (_e *MockStorable_Expecter[K, V, F]) Add(ctx interface{}, item interface{}) *MockStorable_Add_Call[K, V, F] {
	return &MockStorable_Add_Call[K, V, F]{Call: _e.mock.On("Add", ctx, item)}
}

func (_c *MockStorable_Add_Call[K, V, F]) Run(run func(ctx context.Context, item V)) *MockStorable_Add_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(V))
	})
	return _c
}

func (_c *MockStorable_Add_Call[K, V, F]) Return(_a0 K, _a1 error) *MockStorable_Add_Call[K, V, F] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorable_Add_Call[K, V, F]) RunAndReturn(run func(context.Context, V) (K, error)) *MockStorable_Add_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with given fields: ctx
func (_m *MockStorable[K, V, F]) Count(ctx context.Context) int {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockStorable_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockStorable_Count_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockStorable_Expecter[K, V, F]) Count(ctx interface{}) *MockStorable_Count_Call[K, V, F] {
	return &MockStorable_Count_Call[K, V, F]{Call: _e.mock.On("Count", ctx)}
}

func (_c *MockStorable_Count_Call[K, V, F]) Run(run func(ctx context.Context)) *MockStorable_Count_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockStorable_Count_Call[K, V, F]) Return(_a0 int) *MockStorable_Count_Call[K, V, F] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Count_Call[K, V, F]) RunAndReturn(run func(context.Context) int) *MockStorable_Count_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockStorable[K, V, F]) Delete(ctx context.Context, id K) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, K) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorable_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockStorable_Delete_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id K
func (_e *MockStorable_Expecter[K, V, F]) Delete(ctx interface{}, id interface{}) *MockStorable_Delete_Call[K, V, F] {
	return &MockStorable_Delete_Call[K, V, F]{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockStorable_Delete_Call[K, V, F]) Run(run func(ctx context.Context, id K)) *MockStorable_Delete_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(K))
	})
	return _c
}

func (_c *MockStorable_Delete_Call[K, V, F]) Return(_a0 error) *MockStorable_Delete_Call[K, V, F] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Delete_Call[K, V, F]) RunAndReturn(run func(context.Context, K) error) *MockStorable_Delete_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockStorable[K, V, F]) Get(ctx context.Context, id K) (V, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 V
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, K) (V, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, K) V); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(V)
	}

	if rf, ok := ret.Get(1).(func(context.Context, K) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorable_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockStorable_Get_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id K
func (_e *MockStorable_Expecter[K, V, F]) Get(ctx interface{}, id interface{}) *MockStorable_Get_Call[K, V, F] {
	return &MockStorable_Get_Call[K, V, F]{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockStorable_Get_Call[K, V, F]) Run(run func(ctx context.Context, id K)) *MockStorable_Get_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(K))
	})
	return _c
}

func (_c *MockStorable_Get_Call[K, V, F]) Return(_a0 V, _a1 error) *MockStorable_Get_Call[K, V, F] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorable_Get_Call[K, V, F]) RunAndReturn(run func(context.Context, K) (V, error)) *MockStorable_Get_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, filters
func (_m *MockStorable[K, V, F]) List(ctx context.Context, filters F) []V {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []V
	if rf, ok := ret.Get(0).(func(context.Context, F) []V); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]V)
		}
	}

	return r0
}

// MockStorable_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockStorable_List_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - filters F
func (_e *MockStorable_Expecter[K, V, F]) List(ctx interface{}, filters interface{}) *MockStorable_List_Call[K, V, F] {
	return &MockStorable_List_Call[K, V, F]{Call: _e.mock.On("List", ctx, filters)}
}

func (_c *MockStorable_List_Call[K, V, F]) Run(run func(ctx context.Context, filters F)) *MockStorable_List_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(F))
	})
	return _c
}

func (_c *MockStorable_List_Call[K, V, F]) Return(_a0 []V) *MockStorable_List_Call[K, V, F] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_List_Call[K, V, F]) RunAndReturn(run func(context.Context, F) []V) *MockStorable_List_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, item
func (_m *MockStorable[K, V, F]) Update(ctx context.Context, id K, item V) error {
	ret := _m.Called(ctx, id, item)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, K, V) error); ok {
		r0 = rf(ctx, id, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorable_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockStorable_Update_Call[K comparable, V any, F any] struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id K
//   - item V
func (_e *MockStorable_Expecter[K, V, F]) Update(ctx interface{}, id interface{}, item interface{}) *MockStorable_Update_Call[K, V, F] {
	return &MockStorable_Update_Call[K, V, F]{Call: _e.mock.On("Update", ctx, id, item)}
}

func (_c *MockStorable_Update_Call[K, V, F]) Run(run func(ctx context.Context, id K, item V)) *MockStorable_Update_Call[K, V, F] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(K), args[2].(V))
	})
	return _c
}

func (_c *MockStorable_Update_Call[K, V, F]) Return(_a0 error) *MockStorable_Update_Call[K, V, F] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Update_Call[K, V, F]) RunAndReturn(run func(context.Context, K, V) error) *MockStorable_Update_Call[K, V, F] {
	_c.Call.Return(run)
	return _c
}

// NewMockStorable creates a new instance of MockStorable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorable[K comparable, V any, F any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorable[K, V, F] {
	mock := &MockStorable[K, V, F]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

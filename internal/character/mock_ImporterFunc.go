// Code generated by mockery v2.46.3. DO NOT EDIT.

package character

import mock "github.com/stretchr/testify/mock"

// MockImporterFunc is an autogenerated mock type for the ImporterFunc type
type MockImporterFunc struct {
	mock.Mock
}

type MockImporterFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockImporterFunc) EXPECT() *MockImporterFunc_Expecter {
	return &MockImporterFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockImporterFunc) Execute(_a0 string, _a1 []byte) (Character, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 Character
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []byte) (Character, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, []byte) Character); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Character)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockImporterFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockImporterFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 string
//   - _a1 []byte
func (_e *MockImporterFunc_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockImporterFunc_Execute_Call {
	return &MockImporterFunc_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockImporterFunc_Execute_Call) Run(run func(_a0 string, _a1 []byte)) *MockImporterFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *MockImporterFunc_Execute_Call) Return(_a0 Character, _a1 error) *MockImporterFunc_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockImporterFunc_Execute_Call) RunAndReturn(run func(string, []byte) (Character, error)) *MockImporterFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockImporterFunc creates a new instance of MockImporterFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockImporterFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockImporterFunc {
	mock := &MockImporterFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

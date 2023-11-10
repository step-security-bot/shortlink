// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockSchemaServiceServer is an autogenerated mock type for the SchemaServiceServer type
type MockSchemaServiceServer struct {
	mock.Mock
}

type MockSchemaServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSchemaServiceServer) EXPECT() *MockSchemaServiceServer_Expecter {
	return &MockSchemaServiceServer_Expecter{mock: &_m.Mock}
}

// ReadSchema provides a mock function with given fields: _a0, _a1
func (_m *MockSchemaServiceServer) ReadSchema(_a0 context.Context, _a1 *v1.ReadSchemaRequest) (*v1.ReadSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ReadSchemaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadSchemaRequest) (*v1.ReadSchemaResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadSchemaRequest) *v1.ReadSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ReadSchemaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ReadSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchemaServiceServer_ReadSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadSchema'
type MockSchemaServiceServer_ReadSchema_Call struct {
	*mock.Call
}

// ReadSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.ReadSchemaRequest
func (_e *MockSchemaServiceServer_Expecter) ReadSchema(_a0 interface{}, _a1 interface{}) *MockSchemaServiceServer_ReadSchema_Call {
	return &MockSchemaServiceServer_ReadSchema_Call{Call: _e.mock.On("ReadSchema", _a0, _a1)}
}

func (_c *MockSchemaServiceServer_ReadSchema_Call) Run(run func(_a0 context.Context, _a1 *v1.ReadSchemaRequest)) *MockSchemaServiceServer_ReadSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ReadSchemaRequest))
	})
	return _c
}

func (_c *MockSchemaServiceServer_ReadSchema_Call) Return(_a0 *v1.ReadSchemaResponse, _a1 error) *MockSchemaServiceServer_ReadSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchemaServiceServer_ReadSchema_Call) RunAndReturn(run func(context.Context, *v1.ReadSchemaRequest) (*v1.ReadSchemaResponse, error)) *MockSchemaServiceServer_ReadSchema_Call {
	_c.Call.Return(run)
	return _c
}

// WriteSchema provides a mock function with given fields: _a0, _a1
func (_m *MockSchemaServiceServer) WriteSchema(_a0 context.Context, _a1 *v1.WriteSchemaRequest) (*v1.WriteSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.WriteSchemaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteSchemaRequest) (*v1.WriteSchemaResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteSchemaRequest) *v1.WriteSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.WriteSchemaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.WriteSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchemaServiceServer_WriteSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteSchema'
type MockSchemaServiceServer_WriteSchema_Call struct {
	*mock.Call
}

// WriteSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.WriteSchemaRequest
func (_e *MockSchemaServiceServer_Expecter) WriteSchema(_a0 interface{}, _a1 interface{}) *MockSchemaServiceServer_WriteSchema_Call {
	return &MockSchemaServiceServer_WriteSchema_Call{Call: _e.mock.On("WriteSchema", _a0, _a1)}
}

func (_c *MockSchemaServiceServer_WriteSchema_Call) Run(run func(_a0 context.Context, _a1 *v1.WriteSchemaRequest)) *MockSchemaServiceServer_WriteSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.WriteSchemaRequest))
	})
	return _c
}

func (_c *MockSchemaServiceServer_WriteSchema_Call) Return(_a0 *v1.WriteSchemaResponse, _a1 error) *MockSchemaServiceServer_WriteSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchemaServiceServer_WriteSchema_Call) RunAndReturn(run func(context.Context, *v1.WriteSchemaRequest) (*v1.WriteSchemaResponse, error)) *MockSchemaServiceServer_WriteSchema_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedSchemaServiceServer provides a mock function with given fields:
func (_m *MockSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {
	_m.Called()
}

// MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedSchemaServiceServer'
type MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedSchemaServiceServer is a helper method to define mock.On call
func (_e *MockSchemaServiceServer_Expecter) mustEmbedUnimplementedSchemaServiceServer() *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	return &MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedSchemaServiceServer")}
}

func (_c *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) Run(run func()) *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) Return() *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) RunAndReturn(run func()) *MockSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSchemaServiceServer creates a new instance of MockSchemaServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSchemaServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSchemaServiceServer {
	mock := &MockSchemaServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	v1 "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockMetadataServiceServer is an autogenerated mock type for the MetadataServiceServer type
type MockMetadataServiceServer struct {
	mock.Mock
}

type MockMetadataServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetadataServiceServer) EXPECT() *MockMetadataServiceServer_Expecter {
	return &MockMetadataServiceServer_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockMetadataServiceServer) Get(_a0 context.Context, _a1 *v1.MetadataServiceGetRequest) (*v1.MetadataServiceGetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.MetadataServiceGetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceGetRequest) (*v1.MetadataServiceGetResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceGetRequest) *v1.MetadataServiceGetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MetadataServiceGetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MetadataServiceGetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataServiceServer_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockMetadataServiceServer_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.MetadataServiceGetRequest
func (_e *MockMetadataServiceServer_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockMetadataServiceServer_Get_Call {
	return &MockMetadataServiceServer_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockMetadataServiceServer_Get_Call) Run(run func(_a0 context.Context, _a1 *v1.MetadataServiceGetRequest)) *MockMetadataServiceServer_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.MetadataServiceGetRequest))
	})
	return _c
}

func (_c *MockMetadataServiceServer_Get_Call) Return(_a0 *v1.MetadataServiceGetResponse, _a1 error) *MockMetadataServiceServer_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataServiceServer_Get_Call) RunAndReturn(run func(context.Context, *v1.MetadataServiceGetRequest) (*v1.MetadataServiceGetResponse, error)) *MockMetadataServiceServer_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *MockMetadataServiceServer) Set(_a0 context.Context, _a1 *v1.MetadataServiceSetRequest) (*v1.MetadataServiceSetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.MetadataServiceSetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceSetRequest) (*v1.MetadataServiceSetResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MetadataServiceSetRequest) *v1.MetadataServiceSetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MetadataServiceSetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MetadataServiceSetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataServiceServer_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockMetadataServiceServer_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.MetadataServiceSetRequest
func (_e *MockMetadataServiceServer_Expecter) Set(_a0 interface{}, _a1 interface{}) *MockMetadataServiceServer_Set_Call {
	return &MockMetadataServiceServer_Set_Call{Call: _e.mock.On("Set", _a0, _a1)}
}

func (_c *MockMetadataServiceServer_Set_Call) Run(run func(_a0 context.Context, _a1 *v1.MetadataServiceSetRequest)) *MockMetadataServiceServer_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.MetadataServiceSetRequest))
	})
	return _c
}

func (_c *MockMetadataServiceServer_Set_Call) Return(_a0 *v1.MetadataServiceSetResponse, _a1 error) *MockMetadataServiceServer_Set_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataServiceServer_Set_Call) RunAndReturn(run func(context.Context, *v1.MetadataServiceSetRequest) (*v1.MetadataServiceSetResponse, error)) *MockMetadataServiceServer_Set_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedMetadataServiceServer provides a mock function with given fields:
func (_m *MockMetadataServiceServer) mustEmbedUnimplementedMetadataServiceServer() {
	_m.Called()
}

// MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedMetadataServiceServer'
type MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedMetadataServiceServer is a helper method to define mock.On call
func (_e *MockMetadataServiceServer_Expecter) mustEmbedUnimplementedMetadataServiceServer() *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call {
	return &MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedMetadataServiceServer")}
}

func (_c *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call) Run(run func()) *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call) Return() *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call) RunAndReturn(run func()) *MockMetadataServiceServer_mustEmbedUnimplementedMetadataServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetadataServiceServer creates a new instance of MockMetadataServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetadataServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetadataServiceServer {
	mock := &MockMetadataServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

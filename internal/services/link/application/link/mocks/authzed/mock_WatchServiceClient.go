// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockWatchServiceClient is an autogenerated mock type for the WatchServiceClient type
type MockWatchServiceClient struct {
	mock.Mock
}

type MockWatchServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWatchServiceClient) EXPECT() *MockWatchServiceClient_Expecter {
	return &MockWatchServiceClient_Expecter{mock: &_m.Mock}
}

// Watch provides a mock function with given fields: ctx, in, opts
func (_m *MockWatchServiceClient) Watch(ctx context.Context, in *v1.WatchRequest, opts ...grpc.CallOption) (v1.WatchService_WatchClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 v1.WatchService_WatchClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WatchRequest, ...grpc.CallOption) (v1.WatchService_WatchClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WatchRequest, ...grpc.CallOption) v1.WatchService_WatchClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.WatchService_WatchClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.WatchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWatchServiceClient_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockWatchServiceClient_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.WatchRequest
//   - opts ...grpc.CallOption
func (_e *MockWatchServiceClient_Expecter) Watch(ctx interface{}, in interface{}, opts ...interface{}) *MockWatchServiceClient_Watch_Call {
	return &MockWatchServiceClient_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockWatchServiceClient_Watch_Call) Run(run func(ctx context.Context, in *v1.WatchRequest, opts ...grpc.CallOption)) *MockWatchServiceClient_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.WatchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockWatchServiceClient_Watch_Call) Return(_a0 v1.WatchService_WatchClient, _a1 error) *MockWatchServiceClient_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWatchServiceClient_Watch_Call) RunAndReturn(run func(context.Context, *v1.WatchRequest, ...grpc.CallOption) (v1.WatchService_WatchClient, error)) *MockWatchServiceClient_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWatchServiceClient creates a new instance of MockWatchServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWatchServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWatchServiceClient {
	mock := &MockWatchServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

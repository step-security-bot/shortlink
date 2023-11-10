// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockWatchService_WatchClient is an autogenerated mock type for the WatchService_WatchClient type
type MockWatchService_WatchClient struct {
	mock.Mock
}

type MockWatchService_WatchClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWatchService_WatchClient) EXPECT() *MockWatchService_WatchClient_Expecter {
	return &MockWatchService_WatchClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *MockWatchService_WatchClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWatchService_WatchClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockWatchService_WatchClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockWatchService_WatchClient_Expecter) CloseSend() *MockWatchService_WatchClient_CloseSend_Call {
	return &MockWatchService_WatchClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockWatchService_WatchClient_CloseSend_Call) Run(run func()) *MockWatchService_WatchClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchService_WatchClient_CloseSend_Call) Return(_a0 error) *MockWatchService_WatchClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchService_WatchClient_CloseSend_Call) RunAndReturn(run func() error) *MockWatchService_WatchClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *MockWatchService_WatchClient) Context() context.Context {
	ret := _m.Called()

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

// MockWatchService_WatchClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockWatchService_WatchClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockWatchService_WatchClient_Expecter) Context() *MockWatchService_WatchClient_Context_Call {
	return &MockWatchService_WatchClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockWatchService_WatchClient_Context_Call) Run(run func()) *MockWatchService_WatchClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchService_WatchClient_Context_Call) Return(_a0 context.Context) *MockWatchService_WatchClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchService_WatchClient_Context_Call) RunAndReturn(run func() context.Context) *MockWatchService_WatchClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *MockWatchService_WatchClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWatchService_WatchClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockWatchService_WatchClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockWatchService_WatchClient_Expecter) Header() *MockWatchService_WatchClient_Header_Call {
	return &MockWatchService_WatchClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockWatchService_WatchClient_Header_Call) Run(run func()) *MockWatchService_WatchClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchService_WatchClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockWatchService_WatchClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWatchService_WatchClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *MockWatchService_WatchClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *MockWatchService_WatchClient) Recv() (*v1.WatchResponse, error) {
	ret := _m.Called()

	var r0 *v1.WatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.WatchResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.WatchResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.WatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWatchService_WatchClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockWatchService_WatchClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockWatchService_WatchClient_Expecter) Recv() *MockWatchService_WatchClient_Recv_Call {
	return &MockWatchService_WatchClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockWatchService_WatchClient_Recv_Call) Run(run func()) *MockWatchService_WatchClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchService_WatchClient_Recv_Call) Return(_a0 *v1.WatchResponse, _a1 error) *MockWatchService_WatchClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWatchService_WatchClient_Recv_Call) RunAndReturn(run func() (*v1.WatchResponse, error)) *MockWatchService_WatchClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockWatchService_WatchClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWatchService_WatchClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockWatchService_WatchClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockWatchService_WatchClient_Expecter) RecvMsg(m interface{}) *MockWatchService_WatchClient_RecvMsg_Call {
	return &MockWatchService_WatchClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockWatchService_WatchClient_RecvMsg_Call) Run(run func(m interface{})) *MockWatchService_WatchClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockWatchService_WatchClient_RecvMsg_Call) Return(_a0 error) *MockWatchService_WatchClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchService_WatchClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockWatchService_WatchClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockWatchService_WatchClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWatchService_WatchClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockWatchService_WatchClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockWatchService_WatchClient_Expecter) SendMsg(m interface{}) *MockWatchService_WatchClient_SendMsg_Call {
	return &MockWatchService_WatchClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockWatchService_WatchClient_SendMsg_Call) Run(run func(m interface{})) *MockWatchService_WatchClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockWatchService_WatchClient_SendMsg_Call) Return(_a0 error) *MockWatchService_WatchClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchService_WatchClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockWatchService_WatchClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *MockWatchService_WatchClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// MockWatchService_WatchClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockWatchService_WatchClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockWatchService_WatchClient_Expecter) Trailer() *MockWatchService_WatchClient_Trailer_Call {
	return &MockWatchService_WatchClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockWatchService_WatchClient_Trailer_Call) Run(run func()) *MockWatchService_WatchClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWatchService_WatchClient_Trailer_Call) Return(_a0 metadata.MD) *MockWatchService_WatchClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWatchService_WatchClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *MockWatchService_WatchClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWatchService_WatchClient creates a new instance of MockWatchService_WatchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWatchService_WatchClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWatchService_WatchClient {
	mock := &MockWatchService_WatchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

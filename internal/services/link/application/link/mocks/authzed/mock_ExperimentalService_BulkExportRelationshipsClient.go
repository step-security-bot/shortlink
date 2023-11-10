// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockExperimentalService_BulkExportRelationshipsClient is an autogenerated mock type for the ExperimentalService_BulkExportRelationshipsClient type
type MockExperimentalService_BulkExportRelationshipsClient struct {
	mock.Mock
}

type MockExperimentalService_BulkExportRelationshipsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExperimentalService_BulkExportRelationshipsClient) EXPECT() *MockExperimentalService_BulkExportRelationshipsClient_Expecter {
	return &MockExperimentalService_BulkExportRelationshipsClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) CloseSend() *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call) RunAndReturn(run func() error) *MockExperimentalService_BulkExportRelationshipsClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsClient) Context() context.Context {
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

// MockExperimentalService_BulkExportRelationshipsClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockExperimentalService_BulkExportRelationshipsClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) Context() *MockExperimentalService_BulkExportRelationshipsClient_Context_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Context_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Context_Call) Return(_a0 context.Context) *MockExperimentalService_BulkExportRelationshipsClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Context_Call) RunAndReturn(run func() context.Context) *MockExperimentalService_BulkExportRelationshipsClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsClient) Header() (metadata.MD, error) {
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

// MockExperimentalService_BulkExportRelationshipsClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockExperimentalService_BulkExportRelationshipsClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) Header() *MockExperimentalService_BulkExportRelationshipsClient_Header_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Header_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockExperimentalService_BulkExportRelationshipsClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *MockExperimentalService_BulkExportRelationshipsClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsClient) Recv() (*v1.BulkExportRelationshipsResponse, error) {
	ret := _m.Called()

	var r0 *v1.BulkExportRelationshipsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.BulkExportRelationshipsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.BulkExportRelationshipsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.BulkExportRelationshipsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExperimentalService_BulkExportRelationshipsClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockExperimentalService_BulkExportRelationshipsClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) Recv() *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call) Return(_a0 *v1.BulkExportRelationshipsResponse, _a1 error) *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call) RunAndReturn(run func() (*v1.BulkExportRelationshipsResponse, error)) *MockExperimentalService_BulkExportRelationshipsClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockExperimentalService_BulkExportRelationshipsClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) RecvMsg(m interface{}) *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call) Run(run func(m interface{})) *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockExperimentalService_BulkExportRelationshipsClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockExperimentalService_BulkExportRelationshipsClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) SendMsg(m interface{}) *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call) Run(run func(m interface{})) *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call) Return(_a0 error) *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockExperimentalService_BulkExportRelationshipsClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *MockExperimentalService_BulkExportRelationshipsClient) Trailer() metadata.MD {
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

// MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockExperimentalService_BulkExportRelationshipsClient_Expecter) Trailer() *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call {
	return &MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call) Run(run func()) *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call) Return(_a0 metadata.MD) *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *MockExperimentalService_BulkExportRelationshipsClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExperimentalService_BulkExportRelationshipsClient creates a new instance of MockExperimentalService_BulkExportRelationshipsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExperimentalService_BulkExportRelationshipsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExperimentalService_BulkExportRelationshipsClient {
	mock := &MockExperimentalService_BulkExportRelationshipsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

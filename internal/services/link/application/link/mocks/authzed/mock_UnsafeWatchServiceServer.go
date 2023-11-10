// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import mock "github.com/stretchr/testify/mock"

// MockUnsafeWatchServiceServer is an autogenerated mock type for the UnsafeWatchServiceServer type
type MockUnsafeWatchServiceServer struct {
	mock.Mock
}

type MockUnsafeWatchServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeWatchServiceServer) EXPECT() *MockUnsafeWatchServiceServer_Expecter {
	return &MockUnsafeWatchServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedWatchServiceServer provides a mock function with given fields:
func (_m *MockUnsafeWatchServiceServer) mustEmbedUnimplementedWatchServiceServer() {
	_m.Called()
}

// MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedWatchServiceServer'
type MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedWatchServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeWatchServiceServer_Expecter) mustEmbedUnimplementedWatchServiceServer() *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	return &MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedWatchServiceServer")}
}

func (_c *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) Run(run func()) *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) Return() *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call) RunAndReturn(run func()) *MockUnsafeWatchServiceServer_mustEmbedUnimplementedWatchServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUnsafeWatchServiceServer creates a new instance of MockUnsafeWatchServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnsafeWatchServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnsafeWatchServiceServer {
	mock := &MockUnsafeWatchServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

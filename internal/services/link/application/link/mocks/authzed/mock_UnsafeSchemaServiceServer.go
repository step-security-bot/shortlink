// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import mock "github.com/stretchr/testify/mock"

// MockUnsafeSchemaServiceServer is an autogenerated mock type for the UnsafeSchemaServiceServer type
type MockUnsafeSchemaServiceServer struct {
	mock.Mock
}

type MockUnsafeSchemaServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeSchemaServiceServer) EXPECT() *MockUnsafeSchemaServiceServer_Expecter {
	return &MockUnsafeSchemaServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedSchemaServiceServer provides a mock function with given fields:
func (_m *MockUnsafeSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {
	_m.Called()
}

// MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedSchemaServiceServer'
type MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedSchemaServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeSchemaServiceServer_Expecter) mustEmbedUnimplementedSchemaServiceServer() *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	return &MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedSchemaServiceServer")}
}

func (_c *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) Run(run func()) *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) Return() *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call) RunAndReturn(run func()) *MockUnsafeSchemaServiceServer_mustEmbedUnimplementedSchemaServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUnsafeSchemaServiceServer creates a new instance of MockUnsafeSchemaServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnsafeSchemaServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnsafeSchemaServiceServer {
	mock := &MockUnsafeSchemaServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.36.1. DO NOT EDIT.

package crud

import (
	context "context"

	crud "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/repository/crud/postgres/schema/crud"
	mock "github.com/stretchr/testify/mock"

	pgconn "github.com/jackc/pgx/v5/pgconn"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// CreateLink provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateLink(ctx context.Context, arg crud.CreateLinkParams) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, crud.CreateLinkParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_CreateLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLink'
type MockQuerier_CreateLink_Call struct {
	*mock.Call
}

// CreateLink is a helper method to define mock.On call
//   - ctx context.Context
//   - arg crud.CreateLinkParams
func (_e *MockQuerier_Expecter) CreateLink(ctx interface{}, arg interface{}) *MockQuerier_CreateLink_Call {
	return &MockQuerier_CreateLink_Call{Call: _e.mock.On("CreateLink", ctx, arg)}
}

func (_c *MockQuerier_CreateLink_Call) Run(run func(ctx context.Context, arg crud.CreateLinkParams)) *MockQuerier_CreateLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(crud.CreateLinkParams))
	})
	return _c
}

func (_c *MockQuerier_CreateLink_Call) Return(_a0 error) *MockQuerier_CreateLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_CreateLink_Call) RunAndReturn(run func(context.Context, crud.CreateLinkParams) error) *MockQuerier_CreateLink_Call {
	_c.Call.Return(run)
	return _c
}

// CreateLinks provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateLinks(ctx context.Context, arg []crud.CreateLinksParams) (int64, error) {
	ret := _m.Called(ctx, arg)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []crud.CreateLinksParams) (int64, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []crud.CreateLinksParams) int64); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []crud.CreateLinksParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_CreateLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLinks'
type MockQuerier_CreateLinks_Call struct {
	*mock.Call
}

// CreateLinks is a helper method to define mock.On call
//   - ctx context.Context
//   - arg []crud.CreateLinksParams
func (_e *MockQuerier_Expecter) CreateLinks(ctx interface{}, arg interface{}) *MockQuerier_CreateLinks_Call {
	return &MockQuerier_CreateLinks_Call{Call: _e.mock.On("CreateLinks", ctx, arg)}
}

func (_c *MockQuerier_CreateLinks_Call) Run(run func(ctx context.Context, arg []crud.CreateLinksParams)) *MockQuerier_CreateLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]crud.CreateLinksParams))
	})
	return _c
}

func (_c *MockQuerier_CreateLinks_Call) Return(_a0 int64, _a1 error) *MockQuerier_CreateLinks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_CreateLinks_Call) RunAndReturn(run func(context.Context, []crud.CreateLinksParams) (int64, error)) *MockQuerier_CreateLinks_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLink provides a mock function with given fields: ctx, hash
func (_m *MockQuerier) DeleteLink(ctx context.Context, hash string) error {
	ret := _m.Called(ctx, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_DeleteLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLink'
type MockQuerier_DeleteLink_Call struct {
	*mock.Call
}

// DeleteLink is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *MockQuerier_Expecter) DeleteLink(ctx interface{}, hash interface{}) *MockQuerier_DeleteLink_Call {
	return &MockQuerier_DeleteLink_Call{Call: _e.mock.On("DeleteLink", ctx, hash)}
}

func (_c *MockQuerier_DeleteLink_Call) Run(run func(ctx context.Context, hash string)) *MockQuerier_DeleteLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_DeleteLink_Call) Return(_a0 error) *MockQuerier_DeleteLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_DeleteLink_Call) RunAndReturn(run func(context.Context, string) error) *MockQuerier_DeleteLink_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinkByHash provides a mock function with given fields: ctx, hash
func (_m *MockQuerier) GetLinkByHash(ctx context.Context, hash string) (crud.LinkLink, error) {
	ret := _m.Called(ctx, hash)

	var r0 crud.LinkLink
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (crud.LinkLink, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) crud.LinkLink); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(crud.LinkLink)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetLinkByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinkByHash'
type MockQuerier_GetLinkByHash_Call struct {
	*mock.Call
}

// GetLinkByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *MockQuerier_Expecter) GetLinkByHash(ctx interface{}, hash interface{}) *MockQuerier_GetLinkByHash_Call {
	return &MockQuerier_GetLinkByHash_Call{Call: _e.mock.On("GetLinkByHash", ctx, hash)}
}

func (_c *MockQuerier_GetLinkByHash_Call) Run(run func(ctx context.Context, hash string)) *MockQuerier_GetLinkByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_GetLinkByHash_Call) Return(_a0 crud.LinkLink, _a1 error) *MockQuerier_GetLinkByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetLinkByHash_Call) RunAndReturn(run func(context.Context, string) (crud.LinkLink, error)) *MockQuerier_GetLinkByHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinks provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) GetLinks(ctx context.Context, arg crud.GetLinksParams) ([]crud.LinkLink, error) {
	ret := _m.Called(ctx, arg)

	var r0 []crud.LinkLink
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, crud.GetLinksParams) ([]crud.LinkLink, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, crud.GetLinksParams) []crud.LinkLink); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]crud.LinkLink)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, crud.GetLinksParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinks'
type MockQuerier_GetLinks_Call struct {
	*mock.Call
}

// GetLinks is a helper method to define mock.On call
//   - ctx context.Context
//   - arg crud.GetLinksParams
func (_e *MockQuerier_Expecter) GetLinks(ctx interface{}, arg interface{}) *MockQuerier_GetLinks_Call {
	return &MockQuerier_GetLinks_Call{Call: _e.mock.On("GetLinks", ctx, arg)}
}

func (_c *MockQuerier_GetLinks_Call) Run(run func(ctx context.Context, arg crud.GetLinksParams)) *MockQuerier_GetLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(crud.GetLinksParams))
	})
	return _c
}

func (_c *MockQuerier_GetLinks_Call) Return(_a0 []crud.LinkLink, _a1 error) *MockQuerier_GetLinks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetLinks_Call) RunAndReturn(run func(context.Context, crud.GetLinksParams) ([]crud.LinkLink, error)) *MockQuerier_GetLinks_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateLink provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) UpdateLink(ctx context.Context, arg crud.UpdateLinkParams) (pgconn.CommandTag, error) {
	ret := _m.Called(ctx, arg)

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, crud.UpdateLinkParams) (pgconn.CommandTag, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, crud.UpdateLinkParams) pgconn.CommandTag); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, crud.UpdateLinkParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_UpdateLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLink'
type MockQuerier_UpdateLink_Call struct {
	*mock.Call
}

// UpdateLink is a helper method to define mock.On call
//   - ctx context.Context
//   - arg crud.UpdateLinkParams
func (_e *MockQuerier_Expecter) UpdateLink(ctx interface{}, arg interface{}) *MockQuerier_UpdateLink_Call {
	return &MockQuerier_UpdateLink_Call{Call: _e.mock.On("UpdateLink", ctx, arg)}
}

func (_c *MockQuerier_UpdateLink_Call) Run(run func(ctx context.Context, arg crud.UpdateLinkParams)) *MockQuerier_UpdateLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(crud.UpdateLinkParams))
	})
	return _c
}

func (_c *MockQuerier_UpdateLink_Call) Return(_a0 pgconn.CommandTag, _a1 error) *MockQuerier_UpdateLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_UpdateLink_Call) RunAndReturn(run func(context.Context, crud.UpdateLinkParams) (pgconn.CommandTag, error)) *MockQuerier_UpdateLink_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

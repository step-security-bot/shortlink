// Code generated by mockery v2.42.2. DO NOT EDIT.

package account_repository

import (
	context "context"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, in
func (_m *Repository) Add(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *v1.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Account) (*v1.Account, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Account) *v1.Account); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Account) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Repository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.Account
func (_e *Repository_Expecter) Add(ctx interface{}, in interface{}) *Repository_Add_Call {
	return &Repository_Add_Call{Call: _e.mock.On("Add", ctx, in)}
}

func (_c *Repository_Add_Call) Run(run func(ctx context.Context, in *v1.Account)) *Repository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Account))
	})
	return _c
}

func (_c *Repository_Add_Call) Return(_a0 *v1.Account, _a1 error) *Repository_Add_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Add_Call) RunAndReturn(run func(context.Context, *v1.Account) (*v1.Account, error)) *Repository_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Repository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Repository_Expecter) Delete(ctx interface{}, id interface{}) *Repository_Delete_Call {
	return &Repository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *Repository_Delete_Call) Run(run func(ctx context.Context, id string)) *Repository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_Delete_Call) Return(_a0 error) *Repository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Repository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *Repository) Get(ctx context.Context, id string) (*v1.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *v1.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v1.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Repository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Repository_Expecter) Get(ctx interface{}, id interface{}) *Repository_Get_Call {
	return &Repository_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *Repository_Get_Call) Run(run func(ctx context.Context, id string)) *Repository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_Get_Call) Return(_a0 *v1.Account, _a1 error) *Repository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Get_Call) RunAndReturn(run func(context.Context, string) (*v1.Account, error)) *Repository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, filter
func (_m *Repository) List(ctx context.Context, filter interface{}) ([]*v1.Account, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*v1.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) ([]*v1.Account, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) []*v1.Account); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Repository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - filter interface{}
func (_e *Repository_Expecter) List(ctx interface{}, filter interface{}) *Repository_List_Call {
	return &Repository_List_Call{Call: _e.mock.On("List", ctx, filter)}
}

func (_c *Repository_List_Call) Run(run func(ctx context.Context, filter interface{})) *Repository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Repository_List_Call) Return(_a0 []*v1.Account, _a1 error) *Repository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_List_Call) RunAndReturn(run func(context.Context, interface{}) ([]*v1.Account, error)) *Repository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, in
func (_m *Repository) Update(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *v1.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Account) (*v1.Account, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Account) *v1.Account); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Account) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Repository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.Account
func (_e *Repository_Expecter) Update(ctx interface{}, in interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", ctx, in)}
}

func (_c *Repository_Update_Call) Run(run func(ctx context.Context, in *v1.Account)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Account))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 *v1.Account, _a1 error) *Repository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Update_Call) RunAndReturn(run func(context.Context, *v1.Account) (*v1.Account, error)) *Repository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

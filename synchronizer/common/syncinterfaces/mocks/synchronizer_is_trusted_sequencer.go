// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import mock "github.com/stretchr/testify/mock"

// SynchronizerIsTrustedSequencer is an autogenerated mock type for the SynchronizerIsTrustedSequencer type
type SynchronizerIsTrustedSequencer struct {
	mock.Mock
}

type SynchronizerIsTrustedSequencer_Expecter struct {
	mock *mock.Mock
}

func (_m *SynchronizerIsTrustedSequencer) EXPECT() *SynchronizerIsTrustedSequencer_Expecter {
	return &SynchronizerIsTrustedSequencer_Expecter{mock: &_m.Mock}
}

// IsTrustedSequencer provides a mock function with given fields:
func (_m *SynchronizerIsTrustedSequencer) IsTrustedSequencer() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsTrustedSequencer")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsTrustedSequencer'
type SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call struct {
	*mock.Call
}

// IsTrustedSequencer is a helper method to define mock.On call
func (_e *SynchronizerIsTrustedSequencer_Expecter) IsTrustedSequencer() *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call {
	return &SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call{Call: _e.mock.On("IsTrustedSequencer")}
}

func (_c *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call) Run(run func()) *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call) Return(_a0 bool) *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call) RunAndReturn(run func() bool) *SynchronizerIsTrustedSequencer_IsTrustedSequencer_Call {
	_c.Call.Return(run)
	return _c
}

// NewSynchronizerIsTrustedSequencer creates a new instance of SynchronizerIsTrustedSequencer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynchronizerIsTrustedSequencer(t interface {
	mock.TestingT
	Cleanup(func())
}) *SynchronizerIsTrustedSequencer {
	mock := &SynchronizerIsTrustedSequencer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

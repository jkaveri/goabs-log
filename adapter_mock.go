package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// AdapterMock is a mock of ILogAdapter interface
type AdapterMock struct {
	ctrl     *gomock.Controller
	recorder *AdapterMockRecorder
}

// AdapterMockRecorder is the mock recorder for LogAdapterMock
type AdapterMockRecorder struct {
	mock *AdapterMock
}

// NewAdapterMock creates a new mock instance
func NewAdapterMock(ctrl *gomock.Controller) *AdapterMock {
	mock := &AdapterMock{ctrl: ctrl}
	mock.recorder = &AdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *AdapterMock) EXPECT() *AdapterMockRecorder {
	return m.recorder
}

// Log mocks base method
func (m *AdapterMock) Log(fields Fields) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", fields)
}

// Log indicates an expected call of Log
func (mr *AdapterMockRecorder) Log(fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*AdapterMock)(nil).Log), fields)
}

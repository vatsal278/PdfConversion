// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vatsal278/PdfConversion/internal/handler (interfaces: PdfConversionHandler)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPdfConversionHandler is a mock of PdfConversionHandler interface.
type MockPdfConversionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPdfConversionHandlerMockRecorder
}

// MockPdfConversionHandlerMockRecorder is the mock recorder for MockPdfConversionHandler.
type MockPdfConversionHandlerMockRecorder struct {
	mock *MockPdfConversionHandler
}

// NewMockPdfConversionHandler creates a new mock instance.
func NewMockPdfConversionHandler(ctrl *gomock.Controller) *MockPdfConversionHandler {
	mock := &MockPdfConversionHandler{ctrl: ctrl}
	mock.recorder = &MockPdfConversionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPdfConversionHandler) EXPECT() *MockPdfConversionHandlerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockPdfConversionHandler) HealthCheck() (string, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockPdfConversionHandlerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockPdfConversionHandler)(nil).HealthCheck))
}

// Ping mocks base method.
func (m *MockPdfConversionHandler) Ping(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ping", arg0, arg1)
}

// Ping indicates an expected call of Ping.
func (mr *MockPdfConversionHandlerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPdfConversionHandler)(nil).Ping), arg0, arg1)
}
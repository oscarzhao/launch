// Code generated by mockery v1.0.0. DO NOT EDIT.
package z_mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: logTag, format, v
func (_m *Logger) Debug(logTag string, format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, logTag, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: logTag, format, v
func (_m *Logger) Error(logTag string, format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, logTag, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: logTag, format, v
func (_m *Logger) Fatal(logTag string, format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, logTag, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: logTag, format, v
func (_m *Logger) Info(logTag string, format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, logTag, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: logTag, format, v
func (_m *Logger) Warn(logTag string, format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, logTag, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}
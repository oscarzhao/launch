package logging

var _ Logger = &noop{}

// noop implement Logger interface, but does nothing
type noop struct{}

// NewNoop creates a new noop object
func NewNoop() Logger {
	return &noop{}
}

// Debug ...
func (l *noop) Debug(logTag, format string, v ...interface{}) {

}

// Info ...
func (l *noop) Info(logTag, format string, v ...interface{}) {

}

// Warn ...
func (l *noop) Warn(logTag, format string, v ...interface{}) {

}

// Error ...
func (l *noop) Error(logTag, format string, v ...interface{}) {

}

// Fatal ...
func (l *noop) Fatal(logTag, format string, v ...interface{}) {

}

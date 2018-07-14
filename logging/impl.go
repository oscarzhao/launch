package logging

import (
	"io"
	"log"
	"runtime"
)

var buff [1024]byte

type standardLogger struct {
	logLevel Level
	logger   *log.Logger
}

// NewLogger creates a new standardLogger object
func NewLogger(writer io.Writer, logLevel Level) Logger {
	return &standardLogger{
		logLevel: logLevel,
		logger:   log.New(writer, "", log.Llongfile|log.LUTC),
	}
}

// Debug print debug log
func (l *standardLogger) Debug(logTag, format string, v ...interface{}) {
	if l.logLevel <= LevelDebug {
		l.logger.Printf(LevelDebug.String()+": "+logTag+": "+format+"\n", v...)
	}
}

// Info print info log
func (l *standardLogger) Info(logTag, format string, v ...interface{}) {
	if l.logLevel <= LevelInfo {
		l.logger.Printf(LevelInfo.String()+": "+logTag+": "+format+"\n", v...)
	}
}

// Warn print warn log
func (l *standardLogger) Warn(logTag, format string, v ...interface{}) {
	if l.logLevel <= LevelWarn {
		l.logger.Printf(LevelWarn.String()+": "+logTag+": "+format+"\n", v...)
	}
}

// Error print error log
func (l *standardLogger) Error(logTag, format string, v ...interface{}) {
	if l.logLevel <= LevelError {
		stackinfo := runtime.Stack(buff[:], false)
		l.logger.Printf(LevelError.String()+": "+logTag+": "+format+"\n", v...)
		l.logger.Printf("%s\n", buff[:stackinfo])
	}
}

// Fatal print fatal log
func (l *standardLogger) Fatal(logTag, format string, v ...interface{}) {
	l.logger.Panicf(LevelFatal.String()+": "+logTag+": "+format+"\n", v...)
}

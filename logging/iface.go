package logging

import (
	"encoding/json"
	"errors"
	"strings"
)

//go:generate mockery -name=Logger -case=underscore -dir=. -output=../z_mocks -outpkg=z_mocks

// Logger interface
type Logger interface {
	Debug(logTag, format string, v ...interface{})
	Info(logTag, format string, v ...interface{})
	Warn(logTag, format string, v ...interface{})
	Error(logTag, format string, v ...interface{})
	Fatal(logTag, format string, v ...interface{})
}

// Level defines log level le
type Level int

// All Log levels
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelNothing
	noLog
)

var logstrings = [noLog]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var logstr2level = map[string]Level{
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
	"FATAL": LevelFatal,
}

// String is used when formatting
func (l Level) String() string {
	return logstrings[l]
}

// MarshalJSON implements the encoding.JSON interface
func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(logstrings[l])
}

// UnmarshalJSON implements the encoding.JSON interface
func (l *Level) UnmarshalJSON(input []byte) error {
	var val string
	err := json.Unmarshal(input, &val)
	if err != nil {
		return err
	}
	var ok bool
	*l, ok = logstr2level[strings.ToUpper(val)]
	if !ok {
		return errors.New("invalid level")
	}
	return nil
}

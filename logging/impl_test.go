package logging_test

import (
	"bytes"
	"testing"

	"github.com/oscarzhao/launch/logging"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelInfo)
	require.NotNil(t, testLogger)
}

func TestLevelDebug(t *testing.T) {
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelDebug)
	require.NotNil(t, testLogger)

	testLogger.Debug("TestNewLogger", "debug_debug_debug,%s", "oooo")
	testLogger.Info("TestNewLogger", "info_info_info")

	require.Contains(t, buff.String(), "debug_debug_debug,")
	require.Contains(t, buff.String(), "oooo")
	require.Contains(t, buff.String(), "TestNewLogger")
	require.Contains(t, buff.String(), "info_info_info")
}

func TestLevelInfo(t *testing.T) {
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelInfo)
	require.NotNil(t, testLogger)

	testLogger.Debug("TestNewLoggerDebug", "debug_debug_debug,%s", "oooo")
	testLogger.Info("TestNewLoggerInfo", "info_info_info")

	require.NotContains(t, buff.String(), "TestNewLoggerDebug")
	require.NotContains(t, buff.String(), "debug_debug_debug,")
	require.NotContains(t, buff.String(), "oooo")

	require.Contains(t, buff.String(), "TestNewLoggerInfo")
	require.Contains(t, buff.String(), "info_info_info")
}

func TestLevelWarn(t *testing.T) {
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelWarn)
	require.NotNil(t, testLogger)
	testLogger.Debug("TestNewLoggerDebug", "debug_debug_debug,%s", "oooo")
	testLogger.Info("TestNewLoggerInfo", "info_info_info")
	testLogger.Warn("TestNewLoggerWarn", "warn_warn_warn")

	require.NotContains(t, buff.String(), "TestNewLoggerDebug")
	require.NotContains(t, buff.String(), "debug_debug_debug,")
	require.NotContains(t, buff.String(), "oooo")

	require.NotContains(t, buff.String(), "TestNewLoggerInfo")
	require.NotContains(t, buff.String(), "info_info_info")

	require.Contains(t, buff.String(), "TestNewLoggerWarn")
	require.Contains(t, buff.String(), "warn_warn_warn")
}

func TestLevelError(t *testing.T) {
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelError)
	require.NotNil(t, testLogger)
	testLogger.Debug("TestNewLoggerDebug", "debug_debug_debug,%s", "oooo")
	testLogger.Info("TestNewLoggerInfo", "info_info_info")
	testLogger.Warn("TestNewLoggerWarn", "warn_warn_warn")
	testLogger.Error("TestNewLoggerError", "error_error_error")

	require.Contains(t, buff.String(), "TestNewLoggerError")
	require.Contains(t, buff.String(), "error_error_error")
	require.Contains(t, buff.String(), "logging/impl_test.go:")
}

func TestLevelFatal(t *testing.T) {
	defer func() {
		err := recover()
		require.NotNil(t, err)
	}()
	buff := &bytes.Buffer{}
	testLogger := logging.NewLogger(buff, logging.LevelFatal)
	require.NotNil(t, testLogger)
	testLogger.Debug("TestNewLoggerDebug", "debug_debug_debug,%s", "oooo")
	testLogger.Info("TestNewLoggerInfo", "info_info_info")
	testLogger.Warn("TestNewLoggerWarn", "warn_warn_warn")
	testLogger.Error("TestNewLoggerError", "error_error_error")
	testLogger.Fatal("TestNewLoggerFatal", "fatal_fatal_fatal")

	require.NotContains(t, buff.String(), "TestNewLoggerFatal")
	require.NotContains(t, buff.String(), "error_error_error")
}

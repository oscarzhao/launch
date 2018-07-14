package logging

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// NewStdoutLogger creates a logger which print to stdout
func NewStdoutLogger(loglevel Level) Logger {
	return NewLogger(os.Stdout, loglevel)
}

// NewFileLogger creates a logger which print to file
func NewFileLogger(filePath string, logLevel Level) Logger {

	writer, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatalf("create log file fails, path=%s, err=%s\n", filePath, err)
	}

	go func(fp *os.File) {
		c := make(chan os.Signal, 1)
		signal.Notify(c)
		// Block until any signal is received.
		s := <-c
		if s == syscall.SIGTERM || s == syscall.SIGKILL || s == syscall.SIGQUIT || s == syscall.SIGABRT {
			_ = fp.Close()
		}
	}(writer)

	return NewLogger(writer, logLevel)
}

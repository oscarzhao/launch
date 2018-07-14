package manager

import (
	"github.com/oscarzhao/launcher/exec"
	"github.com/oscarzhao/launcher/logging"
)

//go:generate mockery -name=SessionManager -case=underscore -dir=. -output=../z_mocks -outpkg=z_mocks

// SessionManager ...
type SessionManager interface {
	RegisterProc(exec.Execer) error // register a command line task
	StartProc(name string) error    // start a process
	StopProc(name string) error     // stop a process

	Run() error  // start the daemon
	Exit() error // Stop all processes, and quit
}

// NewManager creates a default session manager instance
func NewManager(logger logging.Logger) SessionManager {
	return &defaultSessionManager{
		logger: logger,
		tasks:  make(map[string]exec.Execer),

		quit: make(chan struct{}),
	}
}

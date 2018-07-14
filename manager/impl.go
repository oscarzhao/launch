package manager

import (
	"github.com/oscarzhao/launcher/exec"
	"github.com/oscarzhao/launcher/logging"
)

type defaultSessionManager struct {
	logger logging.Logger
	tasks  map[string]exec.Execer

	quit chan struct{}
}

func (sm *defaultSessionManager) RegisterProc(exec.Execer) error {
	return nil
}

func (sm *defaultSessionManager) StartProc(name string) error {
	return nil
}

func (sm *defaultSessionManager) StopProc(name string) error {
	return nil
}

func (sm *defaultSessionManager) Run() error {
	<-sm.quit
	return nil
}

func (sm *defaultSessionManager) Exit() error {
	return nil
}

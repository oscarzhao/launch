package manager

import (
	"fmt"
	"sync"

	"github.com/oscarzhao/launcher/exec"
	"github.com/oscarzhao/launcher/logging"
)

// ErrorNoSuchTask ...
var ErrorNoSuchTask = fmt.Errorf("No such task registered")

type defaultSessionManager struct {
	logger logging.Logger
	quit   chan struct{}

	lock  sync.RWMutex
	tasks map[string]exec.Execer
}

func (sm *defaultSessionManager) RegisterProc(name string, execer exec.Execer) error {
	sm.lock.Lock()
	sm.tasks[name] = execer
	sm.lock.Unlock()
	return nil
}

func (sm *defaultSessionManager) StartProc(name string) error {
	logTag := "defaultSessionManager.StartProc"

	execer, err := sm.getExecer(name)
	if err != nil {
		sm.logger.Error(logTag, "fails to get execer, name=%s, err=%s", name, err)
		return err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				sm.logger.Error(logTag, "sub process %s crashed, err=%s", name, r)
			}
		}()
		if err := execer.Start(); err != nil {
			sm.logger.Error(logTag, "start sub process %s fails, err=%s", name, err)
		}
	}()

	return nil
}

func (sm *defaultSessionManager) StopProc(name string) error {
	logTag := "defaultSessionManager.StopProc"

	execer, err := sm.getExecer(name)
	if err != nil {
		sm.logger.Error(logTag, "fails to get execer, name=%s, err=%s", name, err)
		return err
	}

	execer.Stop()
	return nil
}

func (sm *defaultSessionManager) Run() error {
	logTag := "defaultSessionManager.Run"
	for name := range sm.tasks {
		if err := sm.StartProc(name); err != nil {
			sm.logger.Error(logTag, "fails to start proc, name=%s, err=%s", name, err)
			return err
		}
	}
	return nil
}

func (sm *defaultSessionManager) Exit() error {

	return nil
}

func (sm *defaultSessionManager) getExecer(name string) (exec.Execer, error) {
	var execer exec.Execer
	var ok bool
	sm.lock.Lock()
	execer, ok = sm.tasks[name]
	sm.lock.Unlock()
	if !ok {
		return nil, ErrorNoSuchTask
	}
	return execer, nil
}

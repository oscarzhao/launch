package main

import (
	"os"

	"github.com/oscarzhao/launcher/config"
	"github.com/oscarzhao/launcher/exec"
	"github.com/oscarzhao/launcher/logging"
	"github.com/oscarzhao/launcher/manager"
)

func main() {
	config.Initialize()

	logger := logging.NewFileLogger(config.Config.Log.Path, config.Config.Log.Level)
	mg := manager.NewManager(logger)
	for _, cmdConf := range config.Config.Commands {
		execer, err := exec.New(cmdConf.Name, cmdConf.BinaryPath, cmdConf.Args,
			exec.WithDaemon(cmdConf.IsDaemon),
			exec.WithEnvMapping(cmdConf.EnvMapping),
			exec.WithRuntimeLogger(logger),
			exec.WithWorkDir(cmdConf.WorkingDir),
		)
		if err != nil {
			logger.Error("main", "fails to create a task, name=%s, err=%s", cmdConf.Name, err)
			os.Exit(1)
		}
		if err := mg.RegisterProc(cmdConf.Name, execer); err != nil {
			logger.Error("main", "fails to register task, name=%s, err=%s", cmdConf.Name, err)
			os.Exit(1)
		}
	}

	if err := mg.StartProc("es5"); err != nil {
		logger.Error("main", "exit with error=%s", err)
		os.Exit(1)
	}
	for {
	}
}

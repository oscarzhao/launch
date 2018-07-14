package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oscarzhao/launcher/config"
	"github.com/oscarzhao/launcher/exec"
	"github.com/oscarzhao/launcher/logging"
	"github.com/oscarzhao/launcher/manager"
)

var serviceName string

func printHelp() {
	fmt.Printf("Usage: launcher start <service name>\n")
}

func init() {
	flag.Parse()
	if flag.NArg() < 2 {
		printHelp()
		os.Exit(1)
	}
	serviceName = flag.Args()[1]
}

func main() {
	config.Initialize()

	logger := logging.NewStdoutLogger(config.Config.Log.Level)
	mg := manager.NewManager(logger)
	for _, cmdConf := range config.Config.Commands {
		execer, err := exec.New(cmdConf.Name, cmdConf.BinaryPath, cmdConf.Args,
			exec.WithEnv(cmdConf.Env),
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

	if err := mg.StartProc(serviceName); err != nil {
		logger.Error("main", "exit with error=%s", err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oscarzhao/launch/config"
	"github.com/oscarzhao/launch/exec"
	"github.com/oscarzhao/launch/logging"
	"github.com/oscarzhao/launch/manager"
)

var serviceName string

func printHelp() {
	fmt.Printf("Usage of launch:\n    launch <service name>\n")
}

func init() {
	flag.Usage = printHelp
	flag.Parse()
	if flag.NArg() < 1 {
		printHelp()
		os.Exit(1)
	}
	serviceName = flag.Args()[0]
}

func main() {
	config.Initialize()

	logger := logging.NewFileLogger(config.Config.Log.Path, config.Config.Log.Level)
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
		fmt.Printf("Start %s failure: %s\n", serviceName, err)
		os.Exit(1)
	}
}

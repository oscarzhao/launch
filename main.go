package main

import (
	"github.com/oscarzhao/launcher/config"
	"github.com/oscarzhao/launcher/logging"
	"github.com/oscarzhao/launcher/manager"
)

func main() {
	config.Initialize()

	logger := logging.NewFileLogger(config.Config.Log.Path, config.Config.Log.Level)
	mg := manager.NewManager(logger)
	if err := mg.Run(); err != nil {
		logger.Error("main", "exit with error=%s", err)
	}
}

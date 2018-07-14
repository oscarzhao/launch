package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/oscarzhao/launcher/logging"
)

// Command defines what a new command line process needs to run
type Command struct {
	Name       string            `json:"name"` // trigger word
	BinaryPath string            `json:"binaryPath"`
	Args       []string          `json:"args"`
	WorkingDir string            `json:"workingDir"`
	EnvMapping map[string]string `json:"envMapping"`
	IsDaemon   bool              `json:"isDaemon"`
	AutoStart  bool              `json:"autoStart"`
}

type logConfiguration struct {
	Path     string        `json:"-"`
	FileName string        `json:"fileName"`
	Level    logging.Level `json:"level"`
}

// Configuration ...
type Configuration struct {
	Log      *logConfiguration `json:"log"`
	Commands []Command         `json:"commands"`
}

// Config ...
var Config Configuration

// Initialize read config from config file
func Initialize() {
	configDir := configDirectory()
	configPath := filepath.Join(configDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Errorf("cannot find configuration file under %s", configPath))
	}
	fileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("cannot read file %s, err=%s", configPath, err))
	}

	if err := json.Unmarshal(fileContent, &Config); err != nil {
		panic(fmt.Errorf("file %s is not valid json format", configPath))
	}

	if Config.Log == nil {
		if len(Config.Log.FileName) == 0 {
			Config.Log.FileName = "launcher.log"
		}
		Config.Log = &logConfiguration{
			Path:  filepath.Join(configDir, Config.Log.FileName),
			Level: logging.LevelInfo,
		}
	} else {
		Config.Log.Path = filepath.Join(configDir, Config.Log.FileName)
	}
}

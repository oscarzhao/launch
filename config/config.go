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
	Name       string   `json:"name"` // trigger word
	BinaryPath string   `json:"binaryPath"`
	Args       []string `json:"args"`
	WorkingDir string   `json:"workingDir"`
	Env        []string `json:"env"`
}

type logConfiguration struct {
	Path  string        `json:"path"`
	Level logging.Level `json:"level"`
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
		Config.Log = &logConfiguration{
			Path:  filepath.Join(configDir, "launcher.log"),
			Level: logging.LevelInfo,
		}
	} else if len(Config.Log.Path) == 0 {
		Config.Log.Path = filepath.Join(configDir, "launcher.log")
	}
}

package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitialize(t *testing.T) {
	configDirectoryBak := configDirectory
	configDirectory = func() string {
		return "../" // config.json is under this directory
	}
	defer func() {
		configDirectory = configDirectoryBak
	}()

	defer func() {
		err := recover()
		require.Nil(t, err)
	}()

	Initialize()

	require.NotNil(t, Config.Log)
	require.Equal(t, 2, len(Config.Commands))
}

func TestInitialize_FileNotExist(t *testing.T) {
	configDirectoryBak := configDirectory
	configDirectory = func() string {
		return os.TempDir() // config.json cannot be found
	}
	defer func() {
		configDirectory = configDirectoryBak
	}()

	defer func() {
		err := recover()
		require.NotNil(t, err)
		require.Contains(t, fmt.Sprintf("%s", err), "cannot find configuration file under")
	}()

	Initialize()
}

func TestConfigDirectory(t *testing.T) {
	dir := configDirectory()
	require.Contains(t, dir, ".launcher")
}

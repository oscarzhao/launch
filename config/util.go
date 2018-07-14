package config

import (
	"log"
	"os/user"
	"path/filepath"
)

// configDirectory check if config folder exists, if not create one
var configDirectory = func() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("fails to get current user, err=%s\n", err)
	}
	return filepath.Join(usr.HomeDir, ".launcher")
}

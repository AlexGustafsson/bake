package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/AlexGustafsson/bake/internal/version"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type lockfile struct {
	Version      string
	Dependencies []string
}

func initCommand(context *cli.Context) error {
	projectDirectory, err := os.Getwd()
	if err != nil {
		log.Error("Unable to get current working directory")
		return err
	}

	rootDirectory := path.Join(projectDirectory, ".bake")
	log.Debugf("Initializing bake in %s", rootDirectory)

	log.Debug("Creating bake home directory")
	os.Mkdir(rootDirectory, 0755)

	log.Debug("Creating lockfile")
	createLockFile(path.Join(rootDirectory, "bake.lock"))

	log.Debug("Copying binary")
	copyBinary(path.Join(rootDirectory, "bake"))

	log.Infof("Initialized bake in %s", rootDirectory)

	return nil
}

func createLockFile(path string) error {
	lockfile := lockfile{
		Version:      version.Version,
		Dependencies: []string{},
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(lockfile)

	return nil
}

func copyBinary(path string) error {
	data, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}

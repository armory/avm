package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func DeleteIfExists(path string) {
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		if err != nil {
			log.Fatalf("Failed to delete %s, err: %s", path, err.Error())
		}
	}
}

func LinkToVersion(version string) {
	log.Infof("using version: %s", version)
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable get home dir, err: %s", err.Error())
	}
	bin := filepath.Join(home, ".avm", "versions", version, "armory")
	log.Infof("ensuring that %s exists", bin)

	if _, err := os.Stat(bin); err != nil {
		log.Fatalf("%s did not exist install it with 'avm install %s'", bin, version)
	}

	binPath := filepath.Join(home, ".avm", "bin")
	err = os.MkdirAll(binPath, os.ModePerm)

	link := filepath.Join(home, ".avm", "bin", "armory")
	if err != nil {
		log.Fatalf("unable to make dirs for: %s, err: %s", link, err.Error())
	}
	DeleteIfExists(link)
	err = os.Symlink(bin, link)
	if err != nil {
		log.Fatalf("unable to link to requested version, err: %s", err.Error())
	}

	log.Infof("linked %s to %s, please ensure that %s is in your path", bin, link, binPath)
}
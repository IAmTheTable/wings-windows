//go:build linux

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pterodactyl/wings/system/osrelease"
)

func UseOpenat2() bool {
	return true
}

func EnsurePterodactylUser() error {
	return nil
}

func ConfigurePasswd() error {
	return nil
}

func EnableLogRotation() error {
	return nil
}

func ConfigureTimezone() error {
	if _config.System.Timezone != "" {
		return nil
	}
	if _, err := os.Stat("/etc/timezone"); err == nil {
		if b, err := os.ReadFile("/etc/timezone"); err == nil {
			_config.System.Timezone = strings.TrimSpace(string(b))
			return nil
		}
	}
	if p, err := filepath.EvalSymlinks("/etc/localtime"); err == nil {
		if strings.Contains(p, "zoneinfo/") {
			_config.System.Timezone = p[strings.Index(p, "zoneinfo/")+9:]
			return nil
		}
	}
	_config.System.Timezone = "UTC"
	return nil
}

func getSystemName() (string, error) {
	release, err := osrelease.Read()
	if err != nil {
		return "", err
	}
	return release["ID"], nil
}

func platformDefaults(c *Configuration) {
	// No-op for Linux
}

func platformFixups(c *Configuration) {
	// No-op for Linux
}

func (s SystemConfiguration) GetStatesPath() string {
	return filepath.Join(s.RootDirectory, "states.json")
}

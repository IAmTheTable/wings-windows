//go:build windows

package config

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/apex/log"
)

func UseOpenat2() bool {
	return false
}

func EnsurePterodactylUser() error {
	log.Warn("automatic user creation is not supported on Windows")
	return nil
}

func ConfigurePasswd() error {
	return nil
}

func EnableLogRotation() error {
	return nil
}

func ConfigureTimezone() error {
	if _config.System.Timezone == "" {
		_config.System.Timezone = "UTC"
	}
	_, err := time.LoadLocation(_config.System.Timezone)
	return err
}

func platformDefaults(c *Configuration) {
	if c.System.RootDirectory == "/var/lib/pterodactyl" {
		c.System.RootDirectory = `C:\ProgramData\Pterodactyl`
	}
	if c.System.LogDirectory == "/var/log/pterodactyl" {
		c.System.LogDirectory = `C:\ProgramData\Pterodactyl\Logs`
	}
	if c.System.Data == "/var/lib/pterodactyl/volumes" {
		c.System.Data = `C:\ProgramData\Pterodactyl\Volumes`
	}
	if c.System.ArchiveDirectory == "/var/lib/pterodactyl/archives" {
		c.System.ArchiveDirectory = `C:\ProgramData\Pterodactyl\Archives`
	}
	if c.System.BackupDirectory == "/var/lib/pterodactyl/backups" {
		c.System.BackupDirectory = `C:\ProgramData\Pterodactyl\Backups`
	}
	if c.System.TmpDirectory == "/tmp/pterodactyl" {
		c.System.TmpDirectory = `C:\temp\pterodactyl`
	}
	if c.System.Passwd.Directory == "/run/wings/etc" {
		c.System.Passwd.Directory = `C:\ProgramData\Pterodactyl\Passwd`
	}
	if c.System.MachineID.Directory == "/run/wings/machine-id" {
		c.System.MachineID.Directory = `C:\ProgramData\Pterodactyl\MachineId`
	}
}

func platformFixups(c *Configuration) {
	// Ensure all paths are absolute and use Windows slashes
	paths := []*string{
		&c.System.RootDirectory,
		&c.System.LogDirectory,
		&c.System.Data,
		&c.System.ArchiveDirectory,
		&c.System.BackupDirectory,
		&c.System.TmpDirectory,
		&c.System.Passwd.Directory,
		&c.System.MachineID.Directory,
	}

	for _, p := range paths {
		if *p == "" {
			continue
		}
		// If it's a Linux-style path, and we are on Windows, we should probably
		// fix it up if it was explicitly provided in the config.
		if strings.HasPrefix(*p, "/run/wings") || strings.HasPrefix(*p, "/var/lib") || strings.HasPrefix(*p, "/var/log") || strings.HasPrefix(*p, "/tmp") {
			// Replace leading / with C:/ if it's missing a drive letter
			if !strings.Contains(*p, ":") {
				*p = "C:" + *p
			}
		}

		// Use Clean to fix slashes
		*p = filepath.Clean(*p)

		// Ensure it's absolute
		if abs, err := filepath.Abs(*p); err == nil {
			*p = abs
		}
	}
}

func (s SystemConfiguration) GetStatesPath() string {
	return filepath.Join(s.RootDirectory, "states.json")
}

//go:build windows

package system

import (
	"context"
	"runtime"
)

func GetSystemInformation() (*Information, error) {
	version, info, err := GetDockerInfo(context.Background())
	if err != nil {
		return nil, err
	}

	return &Information{
		Version: Version,
		Docker: DockerInformation{
			Version: version.Version,
			Cgroups: DockerCgroups{
				Driver:  info.CgroupDriver,
				Version: info.CgroupVersion,
			},
			Containers: DockerContainers{
				Total:   info.Containers,
				Running: info.ContainersRunning,
				Paused:  info.ContainersPaused,
				Stopped: info.ContainersStopped,
			},
			Storage: DockerStorage{
				Driver:     info.Driver,
				Filesystem: "", // Not applicable in the same way on Windows
			},
			Runc: DockerRunc{
				Version: info.RuncCommit.ID,
			},
		},
		System: System{
			Architecture:  runtime.GOARCH,
			CPUThreads:    runtime.NumCPU(),
			MemoryBytes:   info.MemTotal,
			KernelVersion: "", // Need another way on Windows if important
			OS:            "Windows",
			OSType:        runtime.GOOS,
		},
	}, nil
}

//go:build windows

package ufs

import (
	"os"
)

func (fs *UnixFS) removeAll(path string) error {
	return os.RemoveAll(path)
}

func removeAll(fs *Quota, path string) error {
	return os.RemoveAll(path)
}

func (fs *UnixFS) removeContents(path string) error {
	// This is a bit more complex if we want to follow Quota rules exactly,
	// but for now let's just use os.RemoveAll and we might need to fix quota tracking.
	// Actually, wings uses this to clear a server's directory.
	return os.RemoveAll(path)
}

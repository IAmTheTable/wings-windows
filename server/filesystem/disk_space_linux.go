//go:build linux

package filesystem

import (
	"github.com/pterodactyl/wings/internal/ufs"
	"golang.org/x/sys/unix"
)

func (fs *Filesystem) getHardLinkInfo(info ufs.FileInfo) (uint64, bool) {
	if sysFileInfo, ok := info.Sys().(*unix.Stat_t); ok {
		if sysFileInfo.Nlink > 1 {
			return uint64(sysFileInfo.Ino), true
		}
	}
	return 0, false
}

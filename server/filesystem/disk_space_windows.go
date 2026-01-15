package filesystem

import (
	"github.com/pterodactyl/wings/internal/ufs"
)

func (fs *Filesystem) getHardLinkInfo(info ufs.FileInfo) (uint64, bool) {
	// Not easily supported on windows without more complex syscalls
	return 0, false
}

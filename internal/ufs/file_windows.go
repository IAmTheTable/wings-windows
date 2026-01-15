//go:build windows

package ufs

import "syscall"

// Re-using the same names as Go's official `unix` and `os` package do.
const (
	// O_RDONLY opens the file read-only.
	O_RDONLY = syscall.O_RDONLY
	// O_WRONLY opens the file write-only.
	O_WRONLY = syscall.O_WRONLY
	// O_RDWR opens the file read-write.
	O_RDWR = syscall.O_RDWR
	// O_APPEND appends data to the file when writing.
	O_APPEND = syscall.O_APPEND
	// O_CREATE creates a new file if it doesn't exist.
	O_CREATE = syscall.O_CREAT
	// O_EXCL is used with O_CREATE, file must not exist.
	O_EXCL = syscall.O_EXCL
	// O_SYNC open for synchronous I/O.
	O_SYNC = syscall.O_SYNC
	// O_TRUNC truncates regular writable file when opened.
	O_TRUNC = syscall.O_TRUNC
	// O_DIRECTORY opens a directory only. If the entry is not a directory an
	// error will be returned.
	//
	// On Windows this is not strictly supported in the same way, but we map it
	// to 0 or a similar flag if needed. For now, we'll leave it as 0 to avoid
	// build errors, as Windows CreateFile doesn't have an exact direct equivalent
	// flag in the same syscall package set.
	O_DIRECTORY = 0
	// O_NOFOLLOW opens the exact path given without following symlinks.
	// Windows doesn't generally support O_NOFOLLOW in the same way via standard Go syscalls.
	O_NOFOLLOW = 0
	O_CLOEXEC  = syscall.O_CLOEXEC
	// O_LARGEFILE is usually 0 on 64-bit systems or Windows where it's implicit.
	O_LARGEFILE = 0
)

// These are not supported on Windows via syscall.
const (
	AT_SYMLINK_NOFOLLOW = 0
	AT_REMOVEDIR        = 0
	AT_EMPTY_PATH       = 0
)

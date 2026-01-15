//go:build windows

package ufs

import (
	"syscall"
)

// errnoToPathError converts an errno into a proper path error.
func errnoToPathError(err syscall.Errno, op, path string) error {
	switch err {
	// File exists
	case syscall.ERROR_ALREADY_EXISTS, syscall.ERROR_FILE_EXISTS:
		return &PathError{
			Op:   op,
			Path: path,
			Err:  ErrExist,
		}
	// No such file or directory
	case syscall.ERROR_FILE_NOT_FOUND, syscall.ERROR_PATH_NOT_FOUND:
		return &PathError{
			Op:   op,
			Path: path,
			Err:  ErrNotExist,
		}
	// Operation not permitted / Access denied
	case syscall.ERROR_ACCESS_DENIED:
		return &PathError{
			Op:   op,
			Path: path,
			Err:  ErrPermission,
		}
	default:
		return &PathError{
			Op:   op,
			Path: path,
			Err:  err,
		}
	}
}

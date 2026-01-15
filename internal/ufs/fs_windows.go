//go:build windows

package ufs

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UnixFS struct {
	basePath   string
	useOpenat2 bool
}

func NewUnixFS(basePath string, useOpenat2 bool) (*UnixFS, error) {
	basePath = strings.TrimSuffix(basePath, "\\")
	basePath = strings.TrimSuffix(basePath, "/")
	fs := &UnixFS{
		basePath:   basePath,
		useOpenat2: useOpenat2,
	}
	return fs, nil
}

func (fs *UnixFS) BasePath() string {
	return fs.basePath
}

func (fs *UnixFS) Close() error {
	return nil
}

func (fs *UnixFS) Open(path string) (File, error) {
	return os.Open(filepath.Join(fs.basePath, path))
}

func (fs *UnixFS) OpenFile(path string, flag int, perm FileMode) (File, error) {
	return os.OpenFile(filepath.Join(fs.basePath, path), flag, perm)
}

func (fs *UnixFS) Touch(path string, flag int, perm FileMode) (File, error) {
	return os.OpenFile(filepath.Join(fs.basePath, path), flag|O_CREATE, perm)
}

func (fs *UnixFS) Stat(path string) (FileInfo, error) {
	return os.Stat(filepath.Join(fs.basePath, path))
}

func (fs *UnixFS) Lstat(path string) (FileInfo, error) {
	return os.Lstat(filepath.Join(fs.basePath, path))
}

func (fs *UnixFS) Remove(name string) error {
	return os.Remove(filepath.Join(fs.basePath, name))
}

func (fs *UnixFS) MkdirAll(path string, perm FileMode) error {
	return os.MkdirAll(filepath.Join(fs.basePath, path), perm)
}

func (fs *UnixFS) RemoveAll(path string) error {
	return os.RemoveAll(filepath.Join(fs.basePath, path))
}

func (fs *UnixFS) Rename(oldpath, newpath string) error {
	return os.Rename(filepath.Join(fs.basePath, oldpath), filepath.Join(fs.basePath, newpath))
}

func (fs *UnixFS) Symlink(oldpath, newpath string) error {
	return os.Symlink(oldpath, filepath.Join(fs.basePath, newpath))
}

func (fs *UnixFS) Chtimes(path string, atime, mtime time.Time) error {
	return os.Chtimes(filepath.Join(fs.basePath, path), atime, mtime)
}

func (fs *UnixFS) Chmod(path string, mode FileMode) error {
	return os.Chmod(filepath.Join(fs.basePath, path), mode)
}

func (fs *UnixFS) Lchown(path string, uid, gid string) error {
	return nil
}

func (fs *UnixFS) Lchownat(dirfd int, name string, uid, gid string) error {
	return nil
}

func (fs *UnixFS) Lstatat(dirfd int, name string) (FileInfo, error) {
	return os.Lstat(name)
}

func (fs *UnixFS) RemoveStat(name string) (FileInfo, error) {
	return os.Lstat(filepath.Join(fs.basePath, name))
}

func (fs *UnixFS) unsafePath(path string) (string, error) {
	return filepath.Join(fs.basePath, path), nil
}

func (fs *UnixFS) unlinkat(dirfd int, name string, flags int) error {
	return os.Remove(name)
}

func (fs *UnixFS) WalkDirat(dirfd int, name string, fn WalkDiratFunc) error {
	return filepath.WalkDir(name, func(path string, d os.DirEntry, err error) error {
		return fn(0, path, ".", d, err)
	})
}

func (fs *UnixFS) SafePath(path string) (int, string, func(), error) {
	return 0, filepath.Join(fs.basePath, path), func() {}, nil
}

func (fs *UnixFS) OpenFileat(dirfd int, name string, flag int, perm FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

func (fs *UnixFS) ReadDir(path string) ([]DirEntry, error) {
	return os.ReadDir(filepath.Join(fs.basePath, path))
}

package securefs

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

var _ afero.Fs = (*secureFs)(nil)

type secureFs struct {
	fs afero.Fs
}

// New creates a new secure fs, that only holds data in memory encrypted. Read
// operations read bytes in passed in byte slices, which makes the bytes leave a
// protected area. The given byte slice should also be in a protected area, but
// this is the caller's responsibility. Files that are opened and/or created
// with the returned Fs are 100% in memory.
//
//	foo.bar()
//	foo.bar()
func New(fs afero.Fs) afero.Fs {
	return &secureFs{
		fs: fs,
	}
}

func (fs *secureFs) Create(name string) (afero.File, error) {
	file, err := fs.fs.Create(name)
	if err != nil {
		return nil, err
	}
	return newSecureFile(file)
}

func (fs *secureFs) Mkdir(name string, perm os.FileMode) error {
	return fs.fs.Mkdir(name, perm)
}

func (fs *secureFs) MkdirAll(path string, perm os.FileMode) error {
	return fs.fs.MkdirAll(path, perm)
}

func (fs *secureFs) Open(name string) (afero.File, error) {
	file, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return newSecureFile(file)
}

func (fs *secureFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	file, err := fs.fs.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return newSecureFile(file)
}

func (fs *secureFs) Remove(name string) error {
	return fs.fs.Remove(name)
}

func (fs *secureFs) RemoveAll(path string) error {
	return fs.fs.RemoveAll(path)
}

func (fs *secureFs) Rename(oldname, newname string) error {
	return fs.fs.Rename(oldname, newname)
}

func (fs *secureFs) Stat(name string) (os.FileInfo, error) {
	return fs.fs.Stat(name)
}

func (fs *secureFs) Name() string {
	return "secure/" + fs.fs.Name()
}

func (fs *secureFs) Chmod(name string, mode os.FileMode) error {
	return fs.fs.Chmod(name, mode)
}

func (fs *secureFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return fs.fs.Chtimes(name, atime, mtime)
}

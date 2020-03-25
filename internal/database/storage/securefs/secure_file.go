package securefs

import (
	"fmt"
	"io"
	"os"

	"github.com/awnumar/memguard"
	"github.com/spf13/afero"
)

var _ afero.File = (*secureFile)(nil)

type secureFile struct {
	file afero.File

	enclave *memguard.Enclave
	pointer int64
	closed  bool
}

func newSecureFile(file afero.File) (*secureFile, error) {
	secureFile := &secureFile{
		file: file,
	}
	if err := secureFile.load(); err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}
	return secureFile, nil
}

// ReadAt reads len(p) bytes into p, starting from off. If EOF is reached before
// reading was finished, all read bytes are returned, together with an io.EOF
// error. BE AWARE THAT BYTES ARE COPIED FROM A SECURE AREA TO A POTENTIALLY
// INSECURE (your byte slice), AND THAT ALL READ BYTES ARE NO LONGER SECURE.
func (f *secureFile) ReadAt(p []byte, off int64) (int, error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}

	buffer, err := f.enclave.Open()
	if err != nil {
		return 0, fmt.Errorf("open enclave: %w", err)
	}
	defer func() {
		f.enclave = buffer.Seal()
	}()
	data := buffer.Bytes()

	n := copy(p, data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func (f *secureFile) WriteAt(p []byte, off int64) (int, error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}

	if f.enclave == nil || int(off)+len(p) > f.enclave.Size() {
		if err := f.grow(int(off) + len(p)); err != nil {
			return 0, fmt.Errorf("grow: %w", err)
		}
	}

	buffer, err := f.enclave.Open()
	if err != nil {
		return 0, fmt.Errorf("open enclave: %w", err)
	}
	defer func() {
		f.enclave = buffer.Seal()
	}()
	buffer.Melt()
	data := buffer.Bytes()

	n := copy(data[off:off+int64(len(p))], p)
	if n != len(p) {
		return n, fmt.Errorf("unable to write all bytes")
	}
	return n, nil
}

func (f *secureFile) Close() error {
	if f.closed {
		return nil
	}

	if err := f.Sync(); err != nil {
		return fmt.Errorf("sync: %w", err)
	}
	if err := f.file.Close(); err != nil {
		return fmt.Errorf("close underlying: %w", err)
	}
	f.enclave = nil
	f.closed = true
	return nil
}

// Reads len(p) bytes into p and returns the number of bytes read. BE AWARE THAT
// BYTES ARE COPIED FROM A SECURE AREA TO A POTENTIALLY INSECURE (your byte
// slice), AND THAT ALL READ BYTES ARE NO LONGER SECURE.
func (f *secureFile) Read(p []byte) (n int, err error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}
	return f.ReadAt(p, f.pointer)
}

func (f *secureFile) Seek(offset int64, whence int) (int64, error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}

	switch whence {
	case io.SeekCurrent:
		f.pointer += offset
	case io.SeekStart:
		f.pointer = offset
	case io.SeekEnd:
		f.pointer = int64(f.enclave.Size()) - offset
	default:
		return f.pointer, fmt.Errorf("unsupported whence: %v", whence)
	}
	return f.pointer, nil
}

func (f *secureFile) Write(p []byte) (n int, err error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}

	return f.WriteAt(p, f.pointer)
}

func (f *secureFile) Name() string {
	return f.file.Name()
}

func (f *secureFile) Readdir(count int) ([]os.FileInfo, error) {
	return f.file.Readdir(count)
}

func (f *secureFile) Readdirnames(n int) ([]string, error) {
	return f.file.Readdirnames(n)
}

func (f *secureFile) Stat() (os.FileInfo, error) {
	return f.file.Stat()
}

func (f *secureFile) Sync() error {
	if err := f.ensureOpen(); err != nil {
		return err
	}

	buffer, err := f.enclave.Open()
	if err != nil {
		return fmt.Errorf("open enclave: %w", err)
	}
	defer func() {
		f.enclave = buffer.Seal()
	}()

	if err = f.file.Truncate(0); err != nil {
		return fmt.Errorf("truncate: %w", err)
	}
	// Truncate alone doesn't work for memory files, see https://github.com/spf13/afero/issues/235
	_, err = f.file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("seek: %w", err)
	}

	_, err = buffer.Reader().WriteTo(f.file)
	if err != nil {
		return fmt.Errorf("write to: %w", err)
	}
	return nil
}

func (f *secureFile) Truncate(size int64) error {
	if err := f.grow(int(size)); err != nil {
		return fmt.Errorf("grow: %w", err)
	}
	return nil
}

func (f *secureFile) WriteString(s string) (ret int, err error) {
	if err := f.ensureOpen(); err != nil {
		return 0, err
	}

	return f.Write([]byte(s))
}

func (f *secureFile) load() error {
	buffer, err := memguard.NewBufferFromEntireReader(f.file)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}
	f.enclave = buffer.Seal()
	return nil
}

func (f *secureFile) grow(newSize int) error {
	if f.enclave == nil {
		f.enclave = memguard.NewBuffer(newSize).Seal()
		return nil
	}

	oldBuffer, err := f.enclave.Open()
	if err != nil {
		return fmt.Errorf("open enclave: %w", err)
	}
	defer oldBuffer.Destroy()

	// allocate new memory and copy old data
	newBuffer := memguard.NewBuffer(newSize)
	newBuffer.Melt()
	newBuffer.Copy(oldBuffer.Bytes())

	f.enclave = newBuffer.Seal()
	return nil
}

func (f *secureFile) ensureOpen() error {
	if f.closed {
		return afero.ErrFileClosed
	}
	return nil
}

package storage

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

type Validator struct {
	file afero.File
	info os.FileInfo
}

func NewValidator(file afero.File) *Validator {
	return &Validator{
		file: file,
		// info is set on every run of Validate()
	}
}

func (v *Validator) Validate() error {
	stat, err := v.file.Stat()
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	v.info = stat

	if err := v.validateIsFile(); err != nil {
		return fmt.Errorf("is file: %w", err)
	}
	if err := v.validateSize(); err != nil {
		return fmt.Errorf("size: %w", err)
	}

	return nil
}

func (v Validator) validateIsFile() error {
	if v.info.IsDir() {
		return fmt.Errorf("file is directory")
	}
	if !v.info.Mode().Perm().IsRegular() {
		return fmt.Errorf("file is not a regular file")
	}
	return nil
}

func (v Validator) validateSize() error {
	size := v.info.Size()
	if size%page.Size != 0 {
		return fmt.Errorf("invalid file size, must be multiple of page size (=%v), but was %v", page.Size, size)
	}
	return nil
}

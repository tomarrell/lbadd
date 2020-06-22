package storage

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// Validator can be used to validate a database file prior to opening it. If
// validation fails, a speaking error is returned. If validation does not fail,
// the file is a valid database file and can be used. Valid means, that the file
// is not structurally corrupted and usable.
type Validator struct {
	file afero.File
	info os.FileInfo
}

// NewValidator creates a new validator over the given file.
func NewValidator(file afero.File) *Validator {
	return &Validator{
		file: file,
		// info is set on every run of Validate()
	}
}

// Validate runs the file validation and returns a speaking error on why the
// validation failed, if it failed.
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

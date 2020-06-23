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

	validations := []struct {
		name      string
		validator func() error
	}{
		{"is file", v.validateIsFile},
		{"size", v.validateSize},
		{"page count", v.validatePageCount},
	}

	for _, validation := range validations {
		if err := validation.validator(); err != nil {
			return fmt.Errorf("%v: %w", validation.name, err)
		}
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

func (v Validator) validatePageCount() error {
	mgr, err := NewPageManager(v.file)
	if err != nil {
		return fmt.Errorf("new page manager: %w", err)
	}

	headerPage, err := mgr.ReadPage(HeaderPageID)
	if err != nil {
		return fmt.Errorf("read header page: %w", err)
	}

	val, ok := headerPage.Cell([]byte(HeaderPageCount))
	if !ok {
		return fmt.Errorf("no page count header field in header page")
	}
	pageCountField, ok := val.(page.RecordCell)
	if !ok {
		return fmt.Errorf("page count cell is not a record cell (%v)", val.Type())
	}
	pageCount := byteOrder.Uint64(pageCountField.Record)
	if int64(pageCount) != v.info.Size()/page.Size {
		return fmt.Errorf("page count does not match file size (pageCount=%v,size=%v,expected count=%v)", pageCount, v.info.Size(), v.info.Size()/page.Size)
	}
	return nil
}

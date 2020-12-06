package raft

import "fmt"

type multiError []error

var _ error = (*multiError)(nil)

func (e multiError) Error() string {
	return fmt.Sprintf("multiple errors: %v", []error(e))
}

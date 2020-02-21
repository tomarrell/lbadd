package btree

import "fmt"

type (
	Key   int
	Value interface{}
)

// entry is a key/value pair that is stored in the b-tree
type Entry struct {
	key   Key
	value Value
}

func (e *Entry) Key() Key {
	return e.key
}

func (e *Entry) Value() Value {
	return e.value
}

func (e *Entry) String() string {
	if e.value != nil {
		return fmt.Sprintf("\033[1m%v\033[0m", e.key)
	}
	return fmt.Sprintf("%v", e.key)
}

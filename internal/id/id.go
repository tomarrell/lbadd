package id

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

// ID describes a general identifier. An ID has to be unique application-wide.
// IDs must not be re-used.
type ID interface {
	fmt.Stringer
	Bytes() []byte
}

var _ ID = (*id)(nil)

type id ulid.ULID

var (
	lock       sync.Mutex
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	entropy    = ulid.Monotonic(randSource, 0)
)

// Create creates a globally unique ID. This function is safe for concurrent
// use.
func Create() ID {
	lock.Lock()
	defer lock.Unlock()

	genID, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		// For this to happen, the random module would have to fail. Since we
		// use Go's pseudo RNG, which just jumps around a few numbers, instead
		// of using crypto/rand, and we also made this function safe for
		// concurrent use, this is nearly impossible to happen. However, with
		// the current version of oklog/ulid v1.3.1, this will also break after
		// 2121-04-11 11:53:25.01172576 UTC.
		log.Fatal(fmt.Errorf("new ulid: %w", err))
	}
	return id(genID)
}

// Parse parses an ID from a byte slice.
func Parse(idBytes []byte) (ID, error) {
	parsed, err := ulid.Parse(string(idBytes))
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return id(parsed), nil
}

func (id id) String() string {
	return ulid.ULID(id).String()
}

func (id id) Bytes() []byte {
	return []byte(id.String())
}

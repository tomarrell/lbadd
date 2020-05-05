package network

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

var _ ID = (*id)(nil)

type id ulid.ULID

var (
	lock       sync.Mutex
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	entropy    = ulid.Monotonic(randSource, 0)
)

func createID() ID {
	lock.Lock()
	defer lock.Unlock()

	id, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		// For this to happen, the random module would have to fail. Since we
		// use Go's pseudo RNG, which just jumps around a few numbers, instead
		// of using crypto/rand, and we also made this function safe for
		// concurrent use, this is nearly impossible to happen. However, with
		// the current version of oklog/ulid v1.3.1, this will also break after
		// 2121-04-11 11:53:25.01172576 UTC.
		log.Fatal(fmt.Errorf("new ulid: %w", err))
	}
	return ID(id)
}

func (id id) String() string {
	return ulid.ULID(id).String()
}

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

	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		log.Fatal(fmt.Errorf("new ulid: %w", err))
	}
	return ID(id)
}

func (id id) String() string {
	return ulid.ULID(id).String()
}

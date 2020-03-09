package test

import (
	"os"
	"testing"
)

const (
	TestAddress = "localhost:57263"
)

func TestMain(m *testing.M) {

	exitCode := m.Run()

	os.Exit(exitCode)
}

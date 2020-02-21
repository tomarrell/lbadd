package scanner

import (
	"fmt"
	"testing"
	"time"
)

func Test_generateScannerInputAndExpectedOutput(t *testing.T) {
	start := time.Now()
	scIn, _ := generateScannerInputAndExpectedOutput()
	fmt.Printf("took %v\n", time.Since(start).Round(time.Millisecond))

	t.Log(scIn)
	// t.Fail()
}

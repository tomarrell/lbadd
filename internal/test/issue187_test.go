package test

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/engine"
	"github.com/tomarrell/lbadd/internal/engine/profile"
)

func TestIssue187(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "issue187",
		Statement: `VALUES (1,"2",3), (4,"5",6)`,
	})
}

func TestIssue187WithProfile(t *testing.T) {
	prof := profile.NewProfiler()
	RunAndCompare(t, Test{
		Name:      "issue187",
		Statement: `VALUES (1,"2",3), (4,"5",6)`,
		EngineOptions: []engine.Option{
			engine.WithProfiler(prof),
		},
	})
	t.Logf("engine profile:\n%v", prof.Profile().String())
}

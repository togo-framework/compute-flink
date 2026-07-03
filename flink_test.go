package computeflink

import (
	"testing"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/togo"
)

func TestArgv(t *testing.T) {
	t.Setenv("FLINK_OPTS", "-d")
	got := newFlink(togo.New()).argv(compute.Job{Cmd: []string{"job.jar", "--in=x"}})
	want := []string{"run", "-d", "job.jar", "--in=x"}
	for i := range want {
		if i >= len(got) || got[i] != want[i] {
			t.Fatalf("argv = %v, want %v", got, want)
		}
	}
}

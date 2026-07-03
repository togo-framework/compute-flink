// Package computeflink is a togo compute backend that submits jobs to Apache
// Flink via `flink run`. Select with `togo provider:use compute flink`.
//
//   FLINK_BIN   flink binary   (default "flink")
//   FLINK_OPTS  extra `flink run` opts (space-separated)
package computeflink

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"os/exec"
	"strings"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("compute-flink", togo.PriorityService+1, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCompute, "flink", newFlink(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "compute-flink")
		}
		return nil
	})
}

type flink struct {
	bin  string
	opts []string
}

func newFlink(k *togo.Kernel) *flink {
	return &flink{
		bin:  providers.Value(k, providers.CapCompute, "flink", "bin", "flink", false),
		opts: strings.Fields(providers.Value(k, providers.CapCompute, "flink", "opts", "", false)),
	}
}

func (f *flink) argv(job compute.Job) []string {
	args := append([]string{"run"}, f.opts...)
	return append(args, job.Cmd...)
}

func (f *flink) Submit(ctx context.Context, job compute.Job) (compute.Run, error) {
	cmd := exec.CommandContext(ctx, f.bin, f.argv(job)...)
	var buf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &buf, &buf
	run := compute.Run{ID: id(), Status: "succeeded"}
	if err := cmd.Run(); err != nil {
		run.Status = "failed"
	}
	run.Output = buf.String()
	return run, nil
}

func id() string { b := make([]byte, 8); _, _ = rand.Read(b); return hex.EncodeToString(b) }

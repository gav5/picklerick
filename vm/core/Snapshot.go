package core

import (
	"../../kernel/process"
	"../ivm"
)

// Snapshot describes a momentary state of a CPU core for record-keeping.
// The core has the ability to save these snapshots periodically over time.
type Snapshot struct {
	Process process.Process
	Next    ivm.State
}

// TakeSnapshot takes a snapshot of the current state of the CPU core.
func (c *Core) TakeSnapshot() {
	(*c).snapshots = append(c.snapshots, Snapshot{
		Process: c.Process.Copy(),
		Next:    c.Next.Copy(),
	})
}

// Snapshots returns all the snapshots taken.
func (c Core) Snapshots() []Snapshot {
	return c.snapshots
}

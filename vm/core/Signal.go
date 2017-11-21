package core

// Signal is the signal sent from the CPU Core to the VM to indicate completion.
// The VM should wait until this signal is received from all to continue.
type Signal struct {
  Error error
  Halted bool
}

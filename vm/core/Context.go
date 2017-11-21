package core

import (
  "../ivm"
)

// Context describes the current enviornment and task of a CPU Core
type Context struct {
  VM ivm.IVM
  StartPC ivm.Address
}

// NoContextError describes the event in which no context is provided.
type NoContextError struct {
}

func (err NoContextError) Error() string {
  return "no context was provided to the CPU core"
}
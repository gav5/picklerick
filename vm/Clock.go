package vm

// Clock defines a clock module for the virtual machine
type Clock uint32

// Tick executes a single clock
func (c *Clock) Tick() {
	(*c)++
}

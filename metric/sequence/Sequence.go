package sequence

// Sequence records a sequence of values as a metric.
type Sequence struct {
	values []uint32
}

// Value returns the most recent value in the sequence.
func (m Sequence) Value() uint32 {
	return m.values[len(m.values)-1]
}

// Set sets the most recent value by appending to the sequence.
func (m *Sequence) Set(value uint32) {
	m.values = append(m.values, value)
}

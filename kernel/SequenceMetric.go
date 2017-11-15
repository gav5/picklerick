package kernel

// SequenceMetricUint32 records a sequence of values as a metric.
type SequenceMetricUint32 struct {
	values []uint32
}

// Value returns the most recent value in the sequence.
func (m SequenceMetricUint32) Value() uint32 {
	return m.values[len(m.values)-1]
}

// Set sets the most recent value by appending to the sequence.
func (m *SequenceMetricUint32) Set(value uint32) {
	m.values = append(m.values, value)
}

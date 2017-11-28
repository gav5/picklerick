package counter

// Counter records a counting value.
type Counter struct {
	count uint32
}

// Make makes a counted metric.
func Make() Counter {
	return Counter{
		count: 0,
	}
}

// Increment adds to the count of the metric.
func (m *Counter) Increment(value uint32) {
	m.count += value
}

// Value returns the count of the metric.
func (m Counter) Value() uint32 {
	return m.count
}

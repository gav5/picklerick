package kernel

// CountMetricUint32 records a counting value.
type CountMetricUint32 struct {
	count uint32
}

// MakeCountMetricUint32 makes a counted metric.
func MakeCountMetricUint32() CountMetricUint32 {
	return CountMetricUint32{
		count: 0,
	}
}

// Increment adds to the count of the metric.
func (m *CountMetricUint32) Increment(value uint32) {
	m.count += value
}

// Value returns the count of the metric.
func (m CountMetricUint32) Value() uint32 {
	return m.count
}

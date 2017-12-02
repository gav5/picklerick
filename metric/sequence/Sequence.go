package sequence

// Sequence records a sequence of values as a metric.
type Sequence struct {
	values []uint32
}

// Make builds a new sequence metric.
func Make() Sequence {
	return Sequence{
		values: []uint32{},
	}
}

// Value returns the most recent value in the sequence.
func (m Sequence) Value() uint32 {
	return m.values[len(m.values)-1]
}

// Set sets the most recent value by appending to the sequence.
func (m *Sequence) Set(value uint32) {
	m.values = append(m.values, value)
}

// Max returns the largest value in the sequence.
func (m Sequence) Max() uint32 {
	var maxval uint32 = 0x00000000
	for _, v := range m.values {
		if v > maxval {
			maxval = v
		}
	}
	return maxval
}

// Min returns the smallest value in the sequence.
func (m Sequence) Min() uint32 {
	var minval uint32 = 0xffffffff
	for _, v := range m.values {
		if v < minval {
			minval = v
		}
	}
	return minval
}

package metric

// FractionalMetricUint32 records the value of something out of a total.
type FractionalMetricUint32 struct {
	numerators  SequenceMetricUint32
	denominator uint32
}

// MakeFractionalMetricUint32 makes a fractional metric with a given denominator.
func MakeFractionalMetricUint32(denominator uint32) FractionalMetricUint32 {
	return FractionalMetricUint32{
		numerators:  SequenceMetricUint32{},
		denominator: denominator,
	}
}

// SetDenominator sets the numerator of the fractional metric.
func (m *FractionalMetricUint32) SetDenominator(value uint32) {
	m.numerators.Set(value)
}

// Numerator returns the most recently-recorded numerator.
func (m FractionalMetricUint32) Numerator() uint32 {
	return m.numerators.Value()
}

// Value returns the numerator and denominator of the fractional metric.
func (m FractionalMetricUint32) Value() float64 {
	return float64(m.Numerator()) / float64(m.denominator)
}

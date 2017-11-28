package fractional

import "../sequence"

// Fractional records the value of something out of a total.
type Fractional struct {
	numerators  sequence.Sequence
	denominator uint32
}

// Make makes a fractional metric with a given denominator.
func Make(denominator uint32) Fractional {
	return Fractional{
		numerators:  sequence.Sequence{},
		denominator: denominator,
	}
}

// SetDenominator sets the numerator of the fractional metric.
func (m *Fractional) SetDenominator(value uint32) {
	m.numerators.Set(value)
}

// Numerator returns the most recently-recorded numerator.
func (m Fractional) Numerator() uint32 {
	return m.numerators.Value()
}

// Value returns the numerator and denominator of the fractional metric.
func (m Fractional) Value() float64 {
	return float64(m.Numerator()) / float64(m.denominator)
}

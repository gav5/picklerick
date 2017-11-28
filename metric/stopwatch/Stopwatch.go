package stopwatch

import "time"

// Stopwatch records metrics that could be recorded "by a stopwatch".
// (they can be stopped and started multiple times)
type Stopwatch struct {
	durations   []timeframe
	currentTime *time.Time
}

// Make makes a new stopwatch instance.
func Make() Stopwatch {
	return Stopwatch{
		durations: []timeframe{},
		currentTime: nil,
	}
}

// Value displays the current duration of the metric.
func (m Stopwatch) Value() time.Duration {
	var total = time.Duration(0)
	for _, d := range m.durations {
		total += d.value()
	}
	total += time.Now().Sub(*m.currentTime)
	return total
}

// Start marks the start of the duration metric at the current time.
func (m *Stopwatch) Start() {
	m.StartAt(time.Now())
}

// Stop marks the end of the duration metric at the current time.
func (m *Stopwatch) Stop() {
	m.StopAt(time.Now())
}

// StartAt marks the start of the duration metric.
func (m *Stopwatch) StartAt(t time.Time) {
	m.currentTime = new(time.Time)
	*m.currentTime = t
}

// StopAt marks the end of the duration metric.
func (m *Stopwatch) StopAt(t time.Time) {
	ct := *m.currentTime
	m.durations = append(m.durations, timeframe{ct, t})
	m.currentTime = nil
}

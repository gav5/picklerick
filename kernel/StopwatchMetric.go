package kernel

import "time"

// StopwatchMetric records metrics that could be recorded "by a stopwatch".
// (they can be stopped and started multiple times)
type StopwatchMetric struct {
	durations   []DurationMetric
	currentTime *time.Time
}

// Value displays the current duration of the metric.
func (m StopwatchMetric) Value() time.Duration {
	var total = time.Duration(0)
	for _, d := range m.durations {
		total += d.Value()
	}
	total += time.Now().Sub(*m.currentTime)
	return total
}

// Start marks the start of the duration metric.
func (m *StopwatchMetric) Start(t time.Time) {
	m.currentTime = new(time.Time)
	*m.currentTime = t
}

// Stop marks the end of the duration metric.
func (m *StopwatchMetric) Stop(t time.Time) {
	m.durations = append(m.durations, DurationMetric{*m.currentTime, t})
	m.currentTime = nil
}

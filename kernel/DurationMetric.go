package kernel

import "time"

// DurationMetric stores the information on a process.
type DurationMetric struct {
	Start time.Time
	End   time.Time
}

// Value displays the current duration of the metric.
func (m DurationMetric) Value() time.Duration {
	return m.End.Sub(m.Start)
}

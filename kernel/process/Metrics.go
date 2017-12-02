package process

import (
	"../../metric/stopwatch"
)

// Metrics is used to collect information about the process.
type Metrics struct {
	WaitTime       stopwatch.Stopwatch
	CompletionTime stopwatch.Stopwatch
	// IOOperationCount  counter.Counter
	// RAMUse            fractional.Fractional
	// CacheUse          fractional.Fractional
}

// MakeMetrics makes the metrics structure.
func MakeMetrics() Metrics {
	return Metrics{
		WaitTime:       stopwatch.Make(),
		CompletionTime: stopwatch.Make(),
	}
}

type statusTransition struct {
	from Status
	to   Status
}

// ReactToStatus reacts to a change in state.
func (m *Metrics) ReactToStatus(o, n Status) {
	switch n {
	case Ready:
		// we are now waiting and tracking execution
		(*m).CompletionTime.Start()
		(*m).WaitTime.Start()
	case Run:
		// we are not waiting anymore
		(*m).WaitTime.Stop()
	case Wait:
		// we are now waiting again
		(*m).WaitTime.Start()
	case Terminated:
		// we have completed the job!
		(*m).CompletionTime.Stop()
	}
}

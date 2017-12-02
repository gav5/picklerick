package process

import (
	"../../metric/counter"
	"../../metric/sequence"
	"../../metric/stopwatch"
	"../../vm/ivm"
)

// Metrics is used to collect information about the process.
type Metrics struct {
	WaitTime       stopwatch.Stopwatch
	CompletionTime stopwatch.Stopwatch
	Cycles         counter.Counter
	CacheSize      sequence.Sequence
}

// MakeMetrics makes the metrics structure.
func MakeMetrics() Metrics {
	return Metrics{
		WaitTime:       stopwatch.Make(),
		CompletionTime: stopwatch.Make(),
		Cycles:         counter.Make(),
		CacheSize:      sequence.Make(),
	}
}

type statusTransition struct {
	from Status
	to   Status
}

// ReactToState reacts to a change in state.
func (m *Metrics) ReactToState(s ivm.State) {
	// report the cache size
	(*m).CacheSize.Set(uint32(len(s.Caches)))
}

// ReactToStatus reacts to a change in state.
func (m *Metrics) ReactToStatus(s Status) {
	switch s {
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

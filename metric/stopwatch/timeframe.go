package stopwatch

import "time"

type timeframe struct {
  start time.Time
  end time.Time
}

func (m timeframe) value() time.Duration {
  return m.end.Sub(m.start)
}

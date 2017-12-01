package stopwatch

import "testing"

func TestMake(t *testing.T) {
	s := Make()
	_ = s.Value()

	for _ = range []int{1, 2, 3, 4} {
		s.Start()
		s.Stop()
	}
	_ = s.Value()
}

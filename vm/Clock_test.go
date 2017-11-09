package vm

import "testing"

func TestClockTick(t *testing.T) {
	c := Clock(0x0)
	for i := range []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		if uint32(c) != uint32(i) {
			t.Errorf("expected clock to be %v (got %v)", i, c)
		}
		c.Tick()
	}
}

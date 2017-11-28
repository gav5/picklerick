package vm

import "testing"

func TestClockTick(t *testing.T) {
	c := Clock(0x0)
	for i := 0; i < 1000; i++ {
		if uint32(c) != uint32(i*4) {
			t.Errorf("expected clock to be %v (got %v)", i, c)
		}
		c.Tick()
	}
}

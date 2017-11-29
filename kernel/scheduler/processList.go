package scheduler

import (
	"../process"
	"fmt"
	"io"
	"log"
)

type processList struct {
	base       []process.Process
	sortMethod Method
	logger     *log.Logger
}

func (pl processList) Len() int {
	return len(pl.base)
}

func (pl processList) Less(i, j int) bool {
	return pl.sortMethod(pl.base[i], pl.base[j])
}

// Swap is for the heap interface
func (pl processList) Swap(i, j int) {
	pl.base[i], pl.base[j] = pl.base[j], pl.base[i]
}

// Push is for the heap interface
func (pl *processList) Push(x interface{}) {
	p := x.(process.Process)

	if p.IsSleep() {
		pl.logger.Panicf(
			"ERROR pushing: sleep processes not allowed",
		)
	}
	for _, px := range pl.base {
		if p.ProcessNumber == px.ProcessNumber {
			pl.logger.Panicf(
				"ERROR pushing: duplicate process number: %d",
				p.ProcessNumber,
			)
		}
	}
	(*pl).base = append(pl.base, p)
	pl.logger.Printf("Pushed process %d", p.ProcessNumber)
}

// Pop is for the heap interface
func (pl *processList) Pop() interface{} {
	old := pl.base
	n := len(old)
	x := old[n-1]
	(*pl).base = old[0 : n-1]

	pl.logger.Printf("Popped process %d", x.ProcessNumber)
	return x
}

func (pl processList) Peek() process.Process {
	return pl.base[len(pl.base)-1]
}

func (pl processList) fprint(w io.Writer) error {
	for i := pl.Len() - 1; i >= 0; i-- {
		p := pl.base[i]
		out := fmt.Sprintf(
			"[%02d] %-10s p%02d (%d instructions) {RAM: %2d pages} {Disk: %2d pages}\n",
			p.ProcessNumber, p.Status(), p.Priority,
			p.CodeSize, len(p.RAMPageTable), len(p.DiskPageTable),
		)
		if _, err := w.Write([]byte(out)); err != nil {
			return err
		}
	}
	return nil
}

package scheduler

import "../process"
import "log"
import "io"
import "fmt"

type processList struct {
  base []process.Process
  sortMethod Method
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
    log.Panicf("[processList.Push] Error: cannot push in a sleep process!")
  }
  for _, px := range pl.base {
    if p.ProcessNumber == px.ProcessNumber {
      log.Panicf(
        "[processList.Push] Error: duplicate process number: %d",
        p.ProcessNumber,
      )
    }
  }
  (*pl).base = append(pl.base, p)
}

// Pop is for the heap interface
func (pl *processList) Pop() interface{} {
  old := pl.base
  n := len(old)
  x := old[n-1]
  (*pl).base = old[0:n-1]
  return x
}

func (pl processList) fprint(w io.Writer) error {
  for i := pl.Len()-1; i >= 0; i-- {
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

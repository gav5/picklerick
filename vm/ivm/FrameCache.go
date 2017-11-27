package ivm

import "fmt"
import "sort"

// FrameCache is a holding space for addressible frames.
type FrameCache map[FrameNumber]Frame

// AddressFetchWord gets a word at the given address.
func (fc FrameCache) AddressFetchWord(addr Address) Word {
  fn := FrameNumber(addr / FrameSize)
  return fc[fn][addr % FrameSize]
}

// AddressWriteWord writes the given word to the given address.
func (fc *FrameCache) AddressWriteWord(addr Address, word Word) {
  fn := FrameNumber(addr / FrameSize)
  frame := (*fc)[fn]
  frame[addr % FrameSize] = word
  (*fc)[fn] = frame
}

func (fc FrameCache) String() string {
  return fmt.Sprintf("%s", fc.Slice())
}

// Slice returns the frame numbers as a slice.
func (fc FrameCache) Slice() []FrameNumber {
  s := make([]FrameNumber, len(fc))
  i := 0
  for fn := range fc {
    s[i] = fn
    i++
  }
  sort.Sort(frameNumberList(s))
  return s
}

// Copy makes a duplicate FrameCache
func (fc FrameCache) Copy() FrameCache {
  c := make(FrameCache)
  for fn, fr := range fc {
    c[fn] = fr.Copy()
  }
  return c
}

type frameNumberList []FrameNumber

func (fnl frameNumberList) Len() int {
  return len(fnl)
}

func (fnl frameNumberList) Less(i, j int) bool {
  return fnl[i] < fnl[j]
}

func (fnl frameNumberList) Swap(i, j int) {
  fnl[i], fnl[j] = fnl[j], fnl[i]
}

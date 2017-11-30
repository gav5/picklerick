package ivm

import "fmt"
import "sort"
import "io"
import "os"

// FrameCache is a holding space for addressible frames.
type FrameCache map[FrameNumber]Frame

func addressTranslate(addr Address) (FrameNumber, int) {
	byteAddr := addr / 4
	return FrameNumber(byteAddr / FrameSize), (int(byteAddr) % FrameSize)
}

// AddressFetchWord gets a word at the given address.
func (fc FrameCache) AddressFetchWord(addr Address) Word {
	fn, fi := addressTranslate(addr)
	frame, ok := fc[fn]
	if !ok {
		return 0x00000000
	}
	return frame[fi]
}

// AddressWriteWord writes the given word to the given address.
func (fc *FrameCache) AddressWriteWord(addr Address, word Word) {
	fn, fi := addressTranslate(addr)
	frame, ok := (*fc)[fn]
	if !ok {
		frame = MakeFrame()
	}
	frame[fi] = word
	(*fc)[fn] = frame
}

func (fc FrameCache) String() string {
	return fmt.Sprintf("%s", fc.Slice())
}

// Print prints the contents of the FrameCache to Stdout
func (fc FrameCache) Print() error {
	slice := fc.Slice()
	for i, fn := range slice {
		fr := fc[fn]
		fmt.Printf("[%02X: ", int(fn))
		fr.Fprint(os.Stdout)
		fmt.Print("]")
		if i%2 == 1 {
			fmt.Print("\n")
		} else {
			fmt.Print("  ")
		}
	}
	return nil
}

// Fprint prints the contents of the FrameCache to the given writer.
func (fc FrameCache) Fprint(w io.Writer) error {
	slice := fc.Slice()
	var err error
	for _, fn := range slice {
		_, err = fmt.Fprintf(w, "\n[%02X: %v]", int(fn), fc[fn])
		if err != nil {
			return err
		}
	}
	return nil
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

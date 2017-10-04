package reg

import (
	"bytes"
	"fmt"
)

// List is a list of general-purpose registers used by the CPU
type List [16]Storage

func (rl List) String() string {
	var buffer bytes.Buffer

	for index, reg := range rl {
		buffer.WriteString(fmt.Sprintf("r%d:\t%v\n", index, reg))
	}
	return buffer.String()
}

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
		if index > 0 {
			buffer.WriteString(" | ")
		}
		buffer.WriteString(fmt.Sprintf("r%d:%v", index, reg))
	}
	return buffer.String()
}

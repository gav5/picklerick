package instrType

import (
	"regexp"
)

// Components encapsulates the structure of a binary command.
type Components struct {
	Opcode Opcode
	Args   Args
}

// MakeComponents makes a Components object for a hex string
func MakeComponents(cmd string) (Components, error) {
	re := regexp.MustCompile("^(?:0x)([0-9a-fA-F]{2})([0-9a-fA-F]{6})$")
	matches := re.FindStringSubmatch(cmd)

	opcode, opErr := makeOpcode(matches[1])
	if opErr != nil {
		return Components{}, opErr
	}

	args, argErr := makeArgs(matches[2])
	if argErr != nil {
		return Components{}, argErr
	}

	return Components{Opcode: opcode, Args: args}, nil
}

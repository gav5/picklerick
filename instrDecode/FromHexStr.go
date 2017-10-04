package instrDecode

import "../instrType"

// FromHexStr takes a hex string and turns it into an instruction
func FromHexStr(hex string) (instrType.Base, error) {
	c, cErr := instrType.MakeComponents(hex)
	if cErr != nil {
		return nil, cErr
	}
	factory, fErr := decodeOpcode(c.Opcode)
	if fErr != nil {
		return nil, fErr
	}
	return factory(c.Args), nil
}

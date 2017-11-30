package program

import (
	"fmt"
	"io"

	"../../vm/decoder"
	"../../vm/ivm"
)

// DecodedInstructionsList holds a list of decoded instructions
type DecodedInstructionsList []DecodedInstruction

func makeDecodedInstructionsList(rawList []uint32) (DecodedInstructionsList, error) {
	retval := DecodedInstructionsList{}
	for i, rawInstruction := range rawList {
		decodedInstruction, err := decoder.DecodeInstruction(rawInstruction)
		if err != nil {
			return DecodedInstructionsList{}, err
		}
		retval = append(retval, DecodedInstruction{
			ADDR: ivm.AddressForIndex(i),
			RAW:  rawInstruction,
			ASM:  decodedInstruction.Assembly(),
		})
	}
	return retval, nil
}

// FprintInstructions prints instructions in a human-readable format.
func (p Program) FprintInstructions(w io.Writer) error {
	for _, di := range p.decodedInstructionsList {
		_, err := fmt.Fprintf(w, "\n%v", di)
		if err != nil {
			return err
		}
	}
	return nil
}

// FprintInstructionsPC prints instructions with given program counter.
func (p Program) FprintInstructionsPC(w io.Writer, pc ivm.Address) error {
	for _, di := range p.decodedInstructionsList {
		var indicatorChar string
		if pc == di.ADDR {
			indicatorChar = ">"
		} else {
			indicatorChar = " "
		}

		_, err := fmt.Fprintf(w, "\n%s %v", indicatorChar, di)
		if err != nil {
			return err
		}
	}
	return nil
}

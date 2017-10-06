package util

import (
	"fmt"
	"math/big"
	"regexp"
)

// HexExtractionFormatError describes an error in transcribing hex caused by incorrect formatting
type HexExtractionFormatError struct {
	in string
}

func (err HexExtractionFormatError) Error() string {
	return fmt.Sprintf("could not extract because the data is in the incorrect format (given \"%v\")", err.in)
}

// HexExtract8 extracts an unsigned 8-bit value from a hex string of valid format
func HexExtract8(h string) (uint8, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,2}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint8(i.Uint64()), nil
}

// HexExtractArray8 extracts an array of unsigned 8-bit values from hex strings of valid format
func HexExtractArray8(h []string) ([]uint8, error) {
	retval := make([]uint8, len(h))
	for i, strVal := range h {
		x, err := HexExtract8(strVal)
		if err != nil {
			return []uint8{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtractFixed8 extracts an unsigned 8-bit value from a valid 2-character hex string
func HexExtractFixed8(h string) (uint8, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{2}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint8(i.Uint64()), nil
}

// HexExtractArrayFixed8 extracts an array of unsigned 8-bit values from valid 2-character hex strings
func HexExtractArrayFixed8(h []string) ([]uint8, error) {
	retval := make([]uint8, len(h))
	for i, strVal := range h {
		x, err := HexExtractFixed8(strVal)
		if err != nil {
			return []uint8{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtract16 extracts an unsigned 16-bit value from a hex string of valid format
func HexExtract16(h string) (uint16, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,4}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint16(i.Uint64()), nil
}

// HexExtractArray16 extracts an array of unsigned 16-bit values from hex strings of valid format
func HexExtractArray16(h []string) ([]uint16, error) {
	retval := make([]uint16, len(h))
	for i, strVal := range h {
		x, err := HexExtract16(strVal)
		if err != nil {
			return []uint16{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtractFixed16 extracts an unsigned 16-bit value from a valid 4-character hex string
func HexExtractFixed16(h string) (uint16, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{4}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint16(i.Uint64()), nil
}

// HexExtractArrayFixed16 extracts an array of unsigned 16-bit values from valid 4-character hex strings
func HexExtractArrayFixed16(h []string) ([]uint16, error) {
	retval := make([]uint16, len(h))
	for i, strVal := range h {
		x, err := HexExtractFixed16(strVal)
		if err != nil {
			return []uint16{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtract32 extracts an unsigned 32-bit value from a hex string of valid format
func HexExtract32(h string) (uint32, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,8}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint32(i.Uint64()), nil
}

// HexExtractArray32 extracts an array of unsigned 32-bit values from hex strings of valid format
func HexExtractArray32(h []string) ([]uint32, error) {
	retval := make([]uint32, len(h))
	for i, strVal := range h {
		x, err := HexExtract32(strVal)
		if err != nil {
			return []uint32{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtractFixed32 extracts an unsigned 32-bit value from a valid 8-character hex string
func HexExtractFixed32(h string) (uint32, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{8}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint32(i.Uint64()), nil
}

// HexExtractArrayFixed32 extracts an array of unsigned 32-bit values from valid 8-character hex strings
func HexExtractArrayFixed32(h []string) ([]uint32, error) {
	retval := make([]uint32, len(h))
	for i, strVal := range h {
		x, err := HexExtractFixed32(strVal)
		if err != nil {
			return []uint32{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtract64 extracts an unsigned 64-bit value from a hex string of valid format
func HexExtract64(h string) (uint64, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,16}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return i.Uint64(), nil
}

// HexExtractArray64 extracts an array of unsigned 64-bit values from hex strings of valid format
func HexExtractArray64(h []string) ([]uint64, error) {
	retval := make([]uint64, len(h))
	for i, strVal := range h {
		x, err := HexExtract64(strVal)
		if err != nil {
			return []uint64{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

// HexExtractFixed64 extracts an unsigned 64-bit value from a valid 16-character hex string
func HexExtractFixed64(h string) (uint64, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{16}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return i.Uint64(), nil
}

// HexExtractArrayFixed64 extracts an array of unsigned 64-bit values from valid 16-character hex strings
func HexExtractArrayFixed64(h []string) ([]uint64, error) {
	retval := make([]uint64, len(h))
	for i, strVal := range h {
		x, err := HexExtractFixed64(strVal)
		if err != nil {
			return []uint64{}, err
		}
		retval[i] = x
	}
	return retval, nil
}

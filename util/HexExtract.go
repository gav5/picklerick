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

// HexExtract8 extracts an unsigned 8-bit value from a hex string
func HexExtract8(h string) (uint8, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,2}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint8(i.Uint64()), nil
}

// HexExtract16 extracts an unsigned 16-bit value from a hex string
func HexExtract16(h string) (uint16, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,4}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint16(i.Uint64()), nil
}

// HexExtract32 extracts an unsigned 32-bit value from a hex string
func HexExtract32(h string) (uint32, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,8}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return uint32(i.Uint64()), nil
}

// HexExtract64 extracts an unsigned 64-bit value from a hex string
func HexExtract64(h string) (uint64, error) {
	r := regexp.MustCompile("^(?:0x)?[0-9a-fA-F]{1,16}$")
	if !r.MatchString(h) {
		return 0, HexExtractionFormatError{in: h}
	}
	i := new(big.Int)
	i.SetString(h, 16)
	return i.Uint64(), nil
}

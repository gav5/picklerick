package main

import (
	"fmt"

	"./instrDecode"
)

func main() {
	prog := []string{
		"0xC050005C",
		"0x4B060000",
		"0x4B010000",
		"0x4B000000",
		"0x4F0A005C",
		"0x4F0D00DC",
		"0x4C0A0004",
		"0xC0BA0000",
		"0x42BD0000",
		"0x4C0D0004",
		"0x4C060001",
		"0x10658000",
		"0x56810018",
		"0x4B060000",
		"0x4F0900DC",
		"0x43970000",
		"0x05070000",
		"0x4C060001",
		"0x4C090004",
		"0x10658000",
		"0x5681003C",
		"0xC10000AC",
		"0x92000000",
	}
	for index, str := range prog {
		instr, err := instrDecode.FromHexStr(str)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("%s  | %04X |  %s\n", str, (index * 4), instr.ASM())
	}
}

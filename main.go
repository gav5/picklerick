package main

import (
	"fmt"
	"os"

	"./instrDecode"
	"./prog"
)

func main() {
	filename := os.Args[1]
	fmt.Println("picklerick OS")
	fmt.Printf("filename: %s\n", filename)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	pAry, pErr := prog.ParseFile(filename)
	if pErr != nil {
		fmt.Printf("error: %v\n", pErr)
		return
	}
	fmt.Printf("~> Got %d programs!\n", len(pAry))
	for _, p := range pAry {
		fmt.Printf("Job ID: %d\n", p.Job.ID)
		for index, iraw := range p.Job.Instructions {
			instr, iErr := instrDecode.FromUint32(iraw)
			if iErr != nil {
				fmt.Printf("error: %v\n", iErr)
				return
			}
			fmt.Printf("%08X  | %04X |  %s\n", iraw, (index * 4), instr.ASM())
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
	// prog := []uint32{
	// 	0xC050005C,
	// 	0x4B060000,
	// 	0x4B010000,
	// 	0x4B000000,
	// 	0x4F0A005C,
	// 	0x4F0D00DC,
	// 	0x4C0A0004,
	// 	0xC0BA0000,
	// 	0x42BD0000,
	// 	0x4C0D0004,
	// 	0x4C060001,
	// 	0x10658000,
	// 	0x56810018,
	// 	0x4B060000,
	// 	0x4F0900DC,
	// 	0x43970000,
	// 	0x05070000,
	// 	0x4C060001,
	// 	0x4C090004,
	// 	0x10658000,
	// 	0x5681003C,
	// 	0xC10000AC,
	// 	0x92000000,
	// }
	// for index, val := range prog {
	// 	instr, err := instrDecode.FromUint32(val)
	// 	if err != nil {
	// 		fmt.Printf("error: %v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("%08X  | %04X |  %s\n", val, (index * 4), instr.ASM())
	// }
}

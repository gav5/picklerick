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
}

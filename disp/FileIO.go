package disp

//
// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path"
//
// 	"../cpuType"
// 	"../decoder"
// 	"../prog"
// 	"../ram"
// )
//
// var dirs = []string{"asm", "ram", "cpu", "prog"}
//
// // MakeAll makes all the output files for display purposes
// func MakeAll(outdir string, programs []prog.Program, cpus []cpuType.CPU) error {
// 	if err := CleanOutDir(outdir); err != nil {
// 		return err
// 	}
// 	for _, p := range programs {
// 		if err := AssemblyFile(outdir, p); err != nil {
// 			return err
// 		}
// 	}
// 	if err := PhysicalMemoryFile(outdir); err != nil {
// 		return err
// 	}
// 	for _, c := range cpus {
// 		if err := CPUFile(outdir, c); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//
// // CleanOutDir cleans all the files in the outdir (except .keep)
// // this is so no old files from previous executions stay around
// // (otherwise this could communicate something incorrect)
// func CleanOutDir(outdir string) error {
// 	os.RemoveAll(outdir)
// 	for _, d := range dirs {
// 		p := path.Join(outdir, d)
// 		if err := os.MkdirAll(p, 0777); err != nil {
// 			return nil
// 		}
// 	}
// 	return nil
// }
//
// // AssemblyFile prints the program assembly data to the correct file
// func AssemblyFile(outdir string, program prog.Program) error {
// 	filename := fmt.Sprintf("%02d.txt", program.Job.ID)
// 	filepath := path.Join(outdir, "asm", filename)
// 	log.Printf("Generating assembly file: %s", filepath)
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	if err := decoder.ProgramWriteASM(file, program); err != nil {
// 		return err
// 	}
// 	return file.Close()
// }
//
// // PhysicalMemoryFile prints the physical memory to the appropriate file
// func PhysicalMemoryFile(outdir string) error {
// 	filepath := path.Join(outdir, "ram", "dump.txt")
// 	log.Printf("Generating physical memory dump: %s", filepath)
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	if err := ram.FprintPhysicalMemory(file); err != nil {
// 		return err
// 	}
// 	return file.Close()
// }
//
// // CPUFile prints the cpu state to the appropriate file
// func CPUFile(outdir string, c cpuType.CPU) error {
// 	filename := fmt.Sprintf("%d.txt", c.ID)
// 	filepath := path.Join(outdir, "cpu", filename)
// 	log.Printf("Generating cpu state dump: %s", filepath)
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.State.Write(file); err != nil {
// 		return err
// 	}
// 	return file.Close()
// }
//
// // ProgramOutputFile prints the output of a given program
// func ProgramOutputFile(outdir string, c cpuType.CPU) error {
// 	filename := fmt.Sprintf("%d.txt", c.State.Program.Job.ID)
// 	filepath := path.Join(outdir, "prog", filename)
// 	log.Printf("Generating prog state dump: %s", filepath)
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	if _, err := fmt.Fprintln(file, "-- Program --"); err != nil {
// 		return err
// 	}
// 	if err := decoder.ProgramWriteASM(file, c.State.Program); err != nil {
// 		return err
// 	}
// 	if _, err := fmt.Fprintln(file, "\n-- CPU State --"); err != nil {
// 		return err
// 	}
// 	if err := c.State.Write(file); err != nil {
// 		return err
// 	}
// 	if _, err := fmt.Fprintln(file, "\n-- Memory Dump --"); err != nil {
// 		return err
// 	}
// 	if err := ram.FprintPhysicalMemory(file); err != nil {
// 		return err
// 	}
// 	return file.Close()
// }

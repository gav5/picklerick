package disp

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const (
	// tiocgwinsz is something I do not understand
	tiocgwinsz = 1074295912
)

// Task defines a progress-reportable execution.
type Task interface {
	Run(chan Progress)
}

// Progress allows for the reporting of progress to the system.
type Progress struct {
	Title string
	Value float32
}

// RunTask runs the given task and displays the reported progress to the system.
func RunTask(t Task) {
	var p Progress
	var bar string
	cols := terminalWidth()
	progress := make(chan Progress)
	t.Run(progress)
	os.Stdout.Write([]byte("\n\n"))

	for true {
		p = <-progress
		if int(p.Value*100000)%100 == 0 {
			titleStr := fmt.Sprintf("\u001b[1A\u001b[0J• %s\n", p.Title)
			os.Stdout.Write([]byte(titleStr))
			bar = showProgress(p.Value, cols)
			os.Stdout.Write([]byte(bar + "\r"))
			os.Stdout.Sync()
		}
		if p.Value == 1.0 {
			os.Stdout.Write([]byte("\n"))
			os.Stdout.Sync()
			break
		}
	}
	os.Stdout.Write([]byte("\n"))
}

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func terminalWidth() int {
	sizeobj, _ := getWinsize()
	return int(sizeobj.Col)
}

func showProgress(v float32, cols int) string {
	total := 100
	current := int(v * float32(total))
	prefix := fmt.Sprintf("%3d%%", current)
	barStart := " ["
	barEnd := "] "

	barSize := cols - len(prefix+barStart+barEnd)
	amount := int(float32(current) / (float32(total) / float32(barSize)))
	remain := barSize - amount

	bar := strings.Repeat("#", amount) + strings.Repeat(" ", remain)
	return "\u001b[1000D" + bold(prefix) + barStart + bar + barEnd
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWinsize() (*winsize, error) {
	ws := new(winsize)

	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(tiocgwinsz),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(r1) == -1 {
		return nil, os.NewSyscallError("GetWinsize", errno)
	}
	return ws, nil
}

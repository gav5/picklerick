package disp

import (
	"fmt"
	"os"
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
	progress := make(chan Progress)
	t.Run(progress)

	for true {
		p = <-progress
		if int(p.Value*10000)%100 == 0 {
			titleStr := fmt.Sprintf(
				"\r\u001b[0J%s %3d%%",
				bold(p.Title),
				int(p.Value * 100.00),
			)
			os.Stdout.WriteString(titleStr)
			os.Stdout.Sync()
		}
		if p.Value == 1.0 {
			os.Stdout.Write([]byte("\n"))
			os.Stdout.Sync()
			break
		}
	}
	// os.Stdout.Write([]byte("\n"))
}

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

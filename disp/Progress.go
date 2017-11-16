package disp

import (
	"fmt"
)

// ProgressChannel is a channel used to send progress
type ProgressChannel (chan Progress)

// ProgressReporter is a channel used to report progress
type ProgressReporter (chan<- Progress)

// ProgressReceiver is a channel used to receive progress
type ProgressReceiver (<-chan Progress)

// Task defines a progress-reportable execution.
type Task interface {
	Run() ProgressReceiver
}

// Progress allows for the reporting of progress to the system.
type Progress struct {
	Title string
	Value float32
}

// RunTask runs the given task and displays the reported progress to the system.
func RunTask(t Task) {
	var p Progress
	progress := t.Run()

	for true {
		p = <-progress
		fmt.Printf("%s %3d%%\n", bold(p.Title), int(p.Value * 100.0))

		// if int(p.Value*10000)%100 == 0 {
		// 	titleStr := fmt.Sprintf(
		// 		"\r\u001b[0J%s %3d%%",
		// 		bold(p.Title),
		// 		int(p.Value * 100.00),
		// 	)
		// 	os.Stdout.WriteString(titleStr)
		// 	os.Stdout.Sync()
		// }
		// if p.Value == 1.0 {
		// 	os.Stdout.Write([]byte("\n"))
		// 	os.Stdout.Sync()
		// 	break
		// }
	}
}

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

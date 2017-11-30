package logger

import "log"
import "io"
import "os"
import "path"
import "fmt"
import "../../config"

const (

	// LogDir indicates the directory where all the log files will go
	LogDir = "log"

	// GlobalModule indicates what the global logfile will be called
	// (.log is implicitly added to this file, of course)
	GlobalModule = "global"
)

var cfg = config.Default

// New makes a new logger with the given module name.
func New(modulename string) *log.Logger {

	// get shared configuration
	cfg, err := config.Shared()
	if err != nil {
		log.Panicf("[NewLogger] no configuration provided: %v\n", err)
	}

	// get the global log file writer
	globalLogWriter, err := buildGlobalLogWriter()
	if err != nil {
		log.Panicf("[NewLogger] cannot get global logfile: %v\n", err)
	}

	// get the module log file writer
	moduleLogWriter, err := buildModuleLogWriter(modulename)
	if err != nil {
		log.Panicf(
			"[NewLogger] cannot get logfile for module \"%s\": %v",
			modulename, err,
		)
	}

	// determine what to log to (so put into an array first)
	logWriters := []io.Writer{}

	// --quiet suppresses output to Stdout
	if !cfg.Quiet {
		logWriters = append(logWriters, os.Stdout)
	}

	// we should always log to the global log file
	logWriters = append(logWriters, globalLogWriter)

	// we should always log to the module log file
	logWriters = append(logWriters, moduleLogWriter)

	// we want to pipe all of these outputs together!
	combinedWriter := io.MultiWriter(logWriters...)

	prefix := fmt.Sprintf("%s:", modulename)

	return log.New(
		combinedWriter, prefix,
		log.Lshortfile|log.Lmicroseconds,
	)
}

// Dummy returns a fake logger for things like testing.
func Dummy() *log.Logger {
	return log.New(fakeWriter{}, "", log.Lshortfile)
}

type fakeWriter struct{}

func (fakeWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Init sets up all the utilities.
func Init(c config.Config) error {
	var err error

	// assign to the global config variable
	cfg = c

	// ensure there is a log directory to push into.
	err = os.MkdirAll(LogDir, os.ModeDir|os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func buildGlobalLogWriter() (io.Writer, error) {
	return openLogfile(GlobalModule)
}

func buildModuleLogWriter(modulename string) (io.Writer, error) {
	return openLogfile(modulename)
}

func openLogfile(modulename string) (io.Writer, error) {
	path := path.Join(LogDir, fmt.Sprintf("%s.log", modulename))
	return os.Create(path)
}

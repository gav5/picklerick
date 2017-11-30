package report

import (
	"fmt"
	"io"
	"os"
	"path"

	"../vm"
)

// reportBase describes what a qualifying report file should look like
// this makes it easier to go through each report type and generate
type reportBase interface {
	Namespace() string
	Name() string
}

// reportMaker makes a reportBase for a VM
type reportMaker func(*vm.VM) ([]reportBase, reportBuilder, error)

// reportBuilder describes how to make a report of a given type
type reportBuilder interface {
	Extension() string
	Fprint(io.Writer, reportBase) error
}

func directoryPathForReportFile(r reportBase) string {
	return path.Join(
		sharedConfig.Outdir,
		r.Namespace(),
	)
}

// fullPathForReportFile makes a path for a given report and builder
func fullPathForReportFile(r reportBase, b reportBuilder) string {
	return path.Join(
		directoryPathForReportFile(r),
		fmt.Sprintf("%s.%s", r.Name(), b.Extension()),
	)
}

func fileForReport(r reportBase, b reportBuilder) (*os.File, string, error) {
	p := fullPathForReportFile(r, b)
	f, err := os.Create(p)
	return f, p, err
}

func ensureDirectoryExistence(r reportBase) error {
	return os.MkdirAll(
		directoryPathForReportFile(r),
		os.ModePerm|os.ModeDir,
	)
}

// saveReportFile takes a report, builds it, and saves the file
func saveReportFile(r reportBase, b reportBuilder) (string, error) {
	var err error

	// make sure the directory exists
	err = ensureDirectoryExistence(r)
	if err != nil {
		return fullPathForReportFile(r, b), err
	}

	// get the file to write to
	f, p, err := fileForReport(r, b)
	if err != nil {
		return p, err
	}

	// have the builder write to the file
	err = b.Fprint(io.Writer(f), r)
	if err != nil {
		return p, err
	}

	return p, nil
}
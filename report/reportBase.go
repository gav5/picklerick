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

	// get the file to write to (and make sure to close it before this exits)
	f, p, err := fileForReport(r, b)
	defer f.Close()
	if err != nil {
		return p, err
	}

	// have the builder write to the file
	err = b.Fprint(io.Writer(f), r)
	if err != nil {
		return p, err
	}

	// since this succeeded, add to the ignore list
	// (items not on this list will be removed)
	ignoreList = append(ignoreList, p)

	return p, nil
}

var ignoreList = []string{}

func cleanOutdir() error {
	return cleanDir(sharedConfig.Outdir)
}

func cleanDir(p string) error {

	// open the directory
	dirf, err := os.Open(p)
	if err != nil {
		return err
	}
	defer dirf.Close()

	// read the contents of the directory
	dirinfo, err := dirf.Readdir(-1)
	if err != nil {
		return err
	}

	// go through the contents of the directory
	for _, inf := range dirinfo {
		fname := path.Join(p, inf.Name())
		if inf.IsDir() {
			// recursively call on this name joined to the original
			err := cleanDir(fname)
			if err != nil {
				return err
			}
		} else if shouldRemove(fname) {
			// this should be removed from the filesystem
			// (becuase it's not on the ignore list)
			err := os.Remove(fname)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func shouldRemove(p string) bool {
	for _, ip := range ignoreList {
		if ip == p {
			return false
		}
	}
	return true
}

package report

import (
	"encoding/csv"
	"io"
)

// csvReport describes what a report should look like as a csv table
// (based on the base reportBase interface, which is more generic)
// csv reports are made using tables with a first row of headers
type csvReport interface {
	reportBase
	TableHeaders() []string
	TableRows() [][]string
}

// csvReportBuilder describes how to build a metics report file
type csvReportBuilder struct{}

// csv reports are saved as type *.csv for use as a spreadsheet
func (csvReportBuilder) Extension() string {
	return "csv"
}

// csv reports are built using csv format to describe a table
func (csvReportBuilder) Fprint(w io.Writer, rRAW reportBase) error {
	var err error

	// this is assumed to be of type metricsReport (of course)
	r := rRAW.(csvReport)

	// get a CSV writer to make this easier (because why not?)
	csvWriter := csv.NewWriter(w)

	// write the header row to the file
	err = csvWriter.Write(r.TableHeaders())
	if err != nil {
		return err
	}

	// write the data rows to the file
	err = csvWriter.WriteAll(r.TableRows())
	if err != nil {
		return err
	}

	return nil
}

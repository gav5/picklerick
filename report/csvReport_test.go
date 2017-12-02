package report

import (
	"testing"
)

func TestCSVReportBuilderImplementsReportBuilder(t *testing.T) {
	_ = reportBuilder(csvReportBuilder{})
}

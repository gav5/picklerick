package report

import (
	"testing"
)

func TestTXTReportBuilderImplementsReportBuilder(t *testing.T) {
	_ = reportBuilder(txtReportBuilder{})
}

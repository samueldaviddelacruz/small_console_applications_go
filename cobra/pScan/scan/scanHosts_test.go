package scan_test

import (
	"small_console_applications_go/cobra/pScan/scan"
	"testing"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}
	if ps.Open.String() != "closed" {
		t.Errorf("expected %q, got %q instead\n", "closed", ps.Open.String())
	}
	ps.Open = true
	if ps.Open.String() != "open" {
		t.Errorf("expected %q, got %q instead\n", "open", ps.Open.String())
	}
}

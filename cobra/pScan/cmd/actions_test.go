package cmd

import (
	"os"
	"small_console_applications_go/cobra/pScan/scan"
	"testing"
)

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	// Create temp file
	tf, err := os.CreateTemp("", "pScan")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	// Initialize list if requested
	if initList {
		hl := &scan.HostsList{}
		for _, host := range hosts {
			hl.Add(host)
		}
		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}
	// return temp file name and cleanup function
	return tf.Name(), func() {
		os.Remove(tf.Name())
	}
}

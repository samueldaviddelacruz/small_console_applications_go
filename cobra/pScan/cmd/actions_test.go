package cmd

import (
	"bytes"
	"io"
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

func TestHostActions(t *testing.T) {
	// Define hosts for actions test
	hosts := []string{"host1", "host2", "host3"}

	// Test cases for Action test
	testCases := []struct {
		name           string
		args           []string
		expectedOut    string
		initList       bool
		actionFunction func(io.Writer, string, []string) error
	}{
		{
			name:           "AddAction",
			args:           hosts,
			expectedOut:    "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			initList:       false,
			actionFunction: addAction,
		},
		{
			name:           "ListAction",
			expectedOut:    "host1\nhost2\nhost3\n",
			initList:       true,
			actionFunction: listAction,
		},
		{
			name:           "DeleteAction",
			args:           []string{"host1", "host2"},
			expectedOut:    "Deleted host: host1\nDeleted host: host2\n",
			initList:       true,
			actionFunction: deleteAction,
		},
	}
	for _, tc := range testCases {
		tf, cleanup := setup(t, hosts, tc.initList)
		defer cleanup()
		var out bytes.Buffer
		if err := tc.actionFunction(&out, tf, tc.args); err != nil {
			t.Fatalf("Expected no error, but got %q\n", err)
		}
		if out.String() != tc.expectedOut {
			t.Errorf("Expected output %q, but got %q\n", tc.expectedOut, out.String())
		}
	}
}

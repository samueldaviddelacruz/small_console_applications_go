package scan_test

import (
	"errors"
	"os"
	"small_console_applications_go/cobra/pScan/scan"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"AddNew", "host2", 2, nil},
		{"AddExisting", "host1", 1, scan.ErrExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Initialize the list
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}
			err := hl.Add(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatal("expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("expected error %q, got %q instead\n", tc.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %q instead\n", err)
			}
			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}
			if hl.Hosts[1] != tc.host {
				t.Errorf("expected host name %q at index 1, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"RemoveExisting", "host1", 1, nil},
		{"RemoveNotFound", "host3", 1, scan.ErrNotExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Initialize the list
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}
			err := hl.Remove(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatal("expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("expected error %q, got %q instead\n", tc.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %q instead\n", err)
			}
			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}
			if hl.Hosts[0] == tc.host {
				t.Errorf("Host name %q should not be in the list\n", tc.host)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hl1 := &scan.HostsList{}
	hl2 := &scan.HostsList{}
	hostName := "host1"
	hl1.Add(hostName)
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}
	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("error loading list from file: %s", err)
	}
	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("hosts %q should match %q host.", hl1.Hosts[0], hl2.Hosts[0])
	}
}

package main

import (
	"os"
	"strings"
	"testing"

	"bytes"
)

const (
	inputFile = "./testdata/test1.md"

	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(result, expected) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result does not match golden file")
	}
}

func TestRun(t *testing.T) {

	var mockStdout bytes.Buffer
	err := run(inputFile, "", &mockStdout, true)
	if err != nil {
		t.Fatal(err)
	}
	resultFile := strings.TrimSpace(mockStdout.String())
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(result, expected) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result does not match golden filessss")
	}
	//t.Logf("result file: %s", resultFile)
	os.Remove(resultFile)
}

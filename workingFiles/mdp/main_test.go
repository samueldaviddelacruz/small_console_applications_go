package main

import (
	"os"
	"testing"

	"bytes"
)

const (
	inputFile  = "./testdata/test1.md"
	outputFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseContent(input)
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
	err := run(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := os.ReadFile(outputFile)
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
	os.Remove(outputFile)
}

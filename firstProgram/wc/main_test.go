package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	// create a buffer to hold the input
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	// call count with the buffer as input
	res := count(b, false, false)
	// check the result
	if res != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	// create a buffer to hold the input
	b := bytes.NewBufferString("line1\nline2\nline3\n")
	expected := 3
	// call count with the buffer as input
	res := count(b, true, false)
	// check the result
	if res != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, res)
	}
}

// TestCountBytes tests the count function set to count bytes
func TestCountBytes(t *testing.T) {
	// create a buffer to hold the input
	b := bytes.NewBufferString("line1\nline2\nline3\n")
	expected := 18
	// call count with the buffer as input
	res := count(b, false, true)
	// check the result
	if res != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, res)
	}
}

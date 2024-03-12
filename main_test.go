package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 word4")

	expectedOutput := 4
	gotOutput := Count(input, false)

	if expectedOutput != gotOutput {
		t.Errorf("Expected %d, got %d instead.\n", expectedOutput, gotOutput)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("line1 word1\nline2 word2\nline3 word3\n")

	expectedOutput := 3
	gotOutput := Count(input, true)

	if expectedOutput != gotOutput {
		t.Errorf("Expected %d, got %d instead.", expectedOutput, gotOutput)
	}
}

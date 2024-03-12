package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 word4")

	expectedOutputType := "words:"
	expectedOutputCount := 4
	gotOutputType, gotOutputCount := Count(input, false, false)

	if (expectedOutputType != gotOutputType) || (expectedOutputCount != gotOutputCount) {
		t.Errorf("Expected type %s, got type %s instead, expected count %d, got %d", expectedOutputType, gotOutputType, expectedOutputCount, gotOutputCount)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("line1 word1\nline2 word2\nline3 word3\n")

	expectedOutputType := "lines:"
	expectedOutputCount := 3
	gotOutputType, gotOutputCount := Count(input, true, false)

	if (expectedOutputType != gotOutputType) || (expectedOutputCount != gotOutputCount) {
		t.Errorf("Expected type %s, got type %s instead, expected count %d, got %d", expectedOutputType, gotOutputType, expectedOutputCount, gotOutputCount)
	}
}

func TestCountBytes1(t *testing.T) {
	input := bytes.NewBufferString("some text for counting bytes")

	expectedOutputType := "bytes:"
	expectedOutputCount := 28
	gotOutputType, gotOutputCount := Count(input, false, true)

	if (expectedOutputType != gotOutputType) || (expectedOutputCount != gotOutputCount) {
		t.Errorf("Expected type %s, got type %s instead, expected count %d, got %d", expectedOutputType, gotOutputType, expectedOutputCount, gotOutputCount)
	}
}

func TestCountBytes2(t *testing.T) {
	input := bytes.NewBufferString("some second text for counting bytes")

	expectedOutputType := "bytes:"
	expectedOutputCount := 35
	gotOutputType, gotOutputCount := Count(input, true, true)

	if (expectedOutputType != gotOutputType) || (expectedOutputCount != gotOutputCount) {
		t.Errorf("Expected type %s, got type %s instead, expected count %d, got %d", expectedOutputType, gotOutputType, expectedOutputCount, gotOutputCount)
	}
}

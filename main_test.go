package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 word4")

	expectedOutput := 4
	gotOutput := Count(input)

	if expectedOutput != gotOutput {
		t.Errorf("Expected %d, got %d instead.\n", expectedOutput, gotOutput)
	}
}

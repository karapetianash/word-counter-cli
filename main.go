package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()

	fmt.Println(Count(os.Stdin, *lines, *bytes))
}

func Count(r io.Reader, countLines, countBytes bool) (string, int) {
	scanner := bufio.NewScanner(r)

	peace := "lines:"
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
		peace = "words:"
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
		peace = "bytes:"
	}

	count := 0
	for scanner.Scan() {
		count++
	}

	return peace, count
}

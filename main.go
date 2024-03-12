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
	flag.Parse()

	fmt.Println(Count(os.Stdin, *lines))
}

func Count(r io.Reader, countLInes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLInes {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

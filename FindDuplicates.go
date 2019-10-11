package dup

import (
	"bufio"
	"fmt"
	"io"
)

// FindDuplicates reads input lines and prints out which are duplicated with the
// nubmer of times.
func FindDuplicates(input io.Reader, output io.Writer) {
	counts := make(map[string]int)

	countLines(input, counts)
	printLines(output, counts)
}

func countLines(f io.Reader, counts map[string]int) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		counts[scanner.Text()]++
	}
}

func printLines(output io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(output, "%d\t%s\n", n, line)
		}
	}
}

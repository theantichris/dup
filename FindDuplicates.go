package dup

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ErrOpenFile contains the message shown when a file cannot be open.
const ErrOpenFile = "an error occurred opening file"

// FindDuplicates reads input lines and prints out which are duplicated with the
// nubmer of times.
func FindDuplicates(input io.Reader, output io.Writer, args []string) {
	counts := make(map[string]int)
	files := args[1:]

	if len(files) == 0 {
		countLines(input, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(output, "%s %q: %v\n", ErrOpenFile, arg, err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	printResults(output, counts)
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

func printResults(output io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(output, "%d\t%s\n", n, line)
		}
	}
}

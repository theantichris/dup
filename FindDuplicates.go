package dup

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// ErrOpenFile contains the message shown when a file cannot be open.
const ErrOpenFile = "an error occurred opening file"

// FindDuplicates reads input lines and prints out which are duplicated with the
// nubmer of times.
func FindDuplicates(input io.Reader, output io.Writer, args []string) {
	counts := make(map[string]int)
	files := args[1:]

	if len(files) == 0 {
		countFromInput(input, counts)
	} else {
		for _, arg := range files {
			err := countFromFile(arg, output, counts)

			if err != nil {
				fmt.Fprintf(output, "%s %q: %v\n", ErrOpenFile, arg, err)
			}
		}
	}

	printResults(output, counts)
}

func countFromInput(input io.Reader, counts map[string]int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		counts[scanner.Text()]++
	}
}

func countFromFile(fileName string, output io.Writer, counts map[string]int) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	return nil
}

func printResults(output io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(output, "%d\t%s\n", n, line)
		}
	}
}

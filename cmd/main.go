package main

import (
	"os"

	"github.com/theantichris/dup"
)

func main() {
	dup.FindDuplicates(os.Stdin, os.Stdout)
}

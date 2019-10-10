package dup_test

import (
	"bytes"
	"testing"

	"github.com/theantichris/dup"
)

func TestFindDuplicates(t *testing.T) {
	input := &bytes.Buffer{}
	input.Write([]byte("one line\ntwo line\none line\n\n"))
	output := &bytes.Buffer{}

	dup.FindDuplicates(input, output)

	got := output.String()
	want := "2\tone line\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

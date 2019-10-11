package dup_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/theantichris/dup"
)

func TestFindDuplicates(t *testing.T) {
	t.Run("it finds duplicate lines from user input", func(t *testing.T) {
		input := createInitialInput(t, "one line\ntwo line\none line\n\n")
		output := &bytes.Buffer{}

		args := []string{"dup"}

		dup.FindDuplicates(input, output, args)

		got := output.String()
		want := "2\tone line\n"

		assertResults(t, got, want)
	})

	t.Run("it finds duplicate lines from files", func(t *testing.T) {
		input := createInitialInput(t, "")
		output := &bytes.Buffer{}

		initialData := `one line
two line
one line`

		tempFile, removeFile := createTempFile(t, initialData)
		defer removeFile()

		args := []string{"dup", "./" + tempFile.Name()}

		dup.FindDuplicates(input, output, args)

		got := output.String()
		want := "2\tone line\n"

		assertResults(t, got, want)
	})

	t.Run("it reports error if file cannot be open", func(t *testing.T) {
		input := createInitialInput(t, "")
		output := &bytes.Buffer{}

		args := []string{"dup", "./non_existant_file"}

		dup.FindDuplicates(input, output, args)

		got := output.String()

		if !strings.Contains(got, dup.ErrOpenFile) {
			t.Errorf("error not reported: %q", got)
		}
	})
}

func assertResults(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func createInitialInput(t *testing.T, content string) io.Reader {
	t.Helper()

	input := &bytes.Buffer{}
	input.Write([]byte(content))

	return input
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("./", "test_file")

	if err != nil {
		t.Fatalf("could not create tmp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

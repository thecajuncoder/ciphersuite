// Package util contains any functions that may helpful in multiple places throughout the project that don't belong in any other sub-package
package util

import (
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/thecajuncoder/ciphersuite/cipher"
	"github.com/thecajuncoder/ciphersuite/cipher/caesar"
)

var validCipherTypes = map[string]reflect.Type{
	"caesar": reflect.TypeOf(caesar.Cipher{}),
}

// GetCipherFromName returns the matching cipher using the provided cipher name
// Returns nil if no matching cipher was found
func GetCipherFromName(name string) cipher.Cipher {

	// clean the provided string to make name matching case and whitespace in-sensitive
	name = strings.ToLower(strings.TrimSpace(name))

	// uses reflection to instantiate a new instance of the specified Cipher type
	if t, ok := validCipherTypes[name]; ok {
		return reflect.New(t).Interface().(cipher.Cipher)
	}
	return nil
}

// GetInputReader checks the input file parameter
// If an input file was not provided, returns STDIN
// Otherwise attempts to open the file and return a reader for it and any errors that occurred
func GetInputReader(inputFile string) (io.ReadCloser, error) {

	if inputFile == "" {
		return os.Stdin, nil
	}

	return os.Open(inputFile)
}

// GetOutputWriter checks the output file parameter
// If an output file was not provided, returns STDOUT
// Otherwise attempts to open the file and return a writer for it and any errors that occurred
func GetOutputWriter(outputFile string) (io.WriteCloser, error) {

	if outputFile == "" {
		return os.Stdout, nil
	}

	return os.Create(outputFile)
}

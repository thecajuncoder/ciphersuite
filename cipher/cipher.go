package cipher

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Key a key used by the cipher to determine how it encodes or decodes messages.
// A key's value and structure is determined by the type of Cipher using it
type Key string

// Cipher interface for all ciphers used in the Cipher Suite application.
// Defines the basic encoding and decoding functions needed to use the Cipher
type Cipher interface {
	SetKey(Key) error
	GetKey() Key
	Encode(io.Reader, io.Writer) (int, error)
	Decode(io.Reader, io.Writer) (int, error)
}

// EncodeString encodes a string message using the provided Cipher
// Returns the encoded message or any errors that occurred during encoding
func EncodeString(c Cipher, message string) (string, error) {

	r := strings.NewReader(message)

	buf := &bytes.Buffer{}

	_, err := c.Encode(r, buf)

	return buf.String(), err
}

// DecodeString decodes a string message using the provided Cipher
// Returns the decoded message or any errors that occurred during decoding
func DecodeString(c Cipher, message string) (string, error) {

	r := strings.NewReader(message)

	buf := &bytes.Buffer{}

	_, err := c.Decode(r, buf)

	return buf.String(), err
}

// ReadCipherKeyFile reads the contents of a file and turns it into a Cipher Key
// Returns any IO errors that may have occurred during reading
func ReadCipherKeyFile(fileName string) (Key, error) {

	file, err := os.Open(fileName)

	if err != nil {
		return Key(""), fmt.Errorf("Unable to open Cipher key file '%s': %s", fileName, err.Error())
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return Key(""), fmt.Errorf("Unable to open Cipher key file '%s': %s", fileName, err.Error())
	}

	return Key(string(b)), nil
}

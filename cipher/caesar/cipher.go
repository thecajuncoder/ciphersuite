package caesar

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/thecajuncoder/ciphersuite/cipher"
)

// Cipher a struct that the defines the popular Caesar Cipher.
// For more information on the Caesar Cipher, visit: https://en.wikipedia.org/wiki/Caesar_cipher
type Cipher struct {
	offset int
}

// GetKey get this cipher's key
func (me *Cipher) GetKey() cipher.Key {
	return cipher.Key(fmt.Sprintf("%d", me.offset))
}

// SetKey set this cipher's key.
// For the Caesar Cipher, the key is a number representing how many letters to shift each letter in the encoded message by.
// Returns an error if the provided key cannot be parsed into a valid integer
func (me *Cipher) SetKey(key cipher.Key) error {
	offset, err := strconv.Atoi(string(key))
	me.offset = offset
	return err
}

// Encode encode the contents of the incoming reader stream
// and write the encoded message to the provided writer stream using this Cipher
// Returns the number of bytes read (and written) and any IO errors that may have occurred
func (me *Cipher) Encode(r io.Reader, w io.Writer) (int, error) {

	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	defer bw.Flush()

	bytesRead := 0
	var c rune
	var sz int
	var err error

	for c, sz, err = br.ReadRune(); err == nil; c, sz, err = br.ReadRune() {

		bytesRead += sz

		// Shift the letters to encode the message
		bw.WriteRune(shiftRune(c, me.offset))
	}

	if err == io.EOF {
		return bytesRead, nil
	}
	return bytesRead, err
}

// Decode decode the contents of the incoming reader stream
// and write the decoded message to the provided writer stream using this Cipher
// Returns the number of bytes read (and written) and any IO errors that may have occurred
func (me *Cipher) Decode(r io.Reader, w io.Writer) (int, error) {

	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	defer bw.Flush()

	bytesRead := 0
	var c rune
	var sz int
	var err error

	for c, sz, err = br.ReadRune(); err == nil; c, sz, err = br.ReadRune() {

		bytesRead += sz

		// Apply a negative offset to decode the encoded message
		bw.WriteRune(shiftRune(c, -me.offset))
	}

	if err == io.EOF {
		return bytesRead, nil
	}
	return bytesRead, err
}

// shiftRune performs the "shift" for any English letters based on the provided offset
// Letters that "overflow" when shifted are wrapped around to the other side of the alphabet
func shiftRune(r rune, offset int) rune {

	offset %= 26
	ascii := int(r)

	// Check only for capital or lowercase English letters (A - Z)
	if (ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122) {

		if ascii <= 90 && ascii+offset > 90 || ascii+offset > 122 {
			ascii += offset - 26
		} else if (ascii <= 90 && ascii+offset < 65) || (ascii >= 97 && ascii+offset < 97) {
			ascii += offset + 26
		} else {
			ascii += offset
		}
	}

	return rune(ascii)
}

package caesar

import (
	"fmt"
	"io"
	"strconv"

	"github.com/thecajuncoder/ciphersuite/cipher"
	"github.com/thecajuncoder/ciphersuite/runeops"
)

// messageOperation defines either an encode or decode message operation to perform on a single letter in a message
type messageOperation func(r rune) rune

// Cipher a struct that the defines the popular Caesar Cipher.
// This cipher scrambles an input message by "shifting" all of the letters of the alphabet by a certain amount.
// The message can then be decoded by shifting the letters back, or shifting by a negative number.
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

	return runeops.ReadAllRunes(r, w, func(r rune) (rune, error) {
		return runeops.ShiftRune(r, me.offset), nil
	})
}

// Decode decode the contents of the incoming reader stream
// and write the decoded message to the provided writer stream using this Cipher
// Returns the number of bytes read (and written) and any IO errors that may have occurred
func (me *Cipher) Decode(r io.Reader, w io.Writer) (int, error) {

	return runeops.ReadAllRunes(r, w, func(r rune) (rune, error) {
		return runeops.ShiftRune(r, -me.offset), nil
	})
}

package vigenere

import (
	"errors"
	"io"
	"regexp"
	"strings"

	"github.com/thecajuncoder/ciphersuite/runeops"

	"github.com/thecajuncoder/ciphersuite/cipher"
)

// runeOperation defines a function for either an encode or decode operation on a single letter in a message
type runeOperation func(r rune, offset int) rune

var keyCleaningRegex *regexp.Regexp = regexp.MustCompile("[^A-Z]")

// Cipher a struct that the defines the Vigenère Cipher.
// This cipher works similarly to the Caesar cipher. It encodes messages by shifting letters.
// However, instead of shifting all letters in the message by a single amount, this cipher shifts each letter
// by a different amount based on a provided code word or "key". Each letter in the key determines how much the
// next letter in the message gets shifted by. A single letter key is the same as a Caesar Cipher.
// For more information on the Vigenère Cipher, visit: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Cipher struct {
	key string
}

// GetKey get this cipher's key
func (me *Cipher) GetKey() cipher.Key {
	return cipher.Key(me.key)
}

// SetKey set this cipher's key.
// For the Caesar Cipher, the key is a number representing how many letters to shift each letter in the encoded message by.
// Returns an error if the provided key cannot be parsed into a valid integer
func (me *Cipher) SetKey(key cipher.Key) error {

	str := strings.ToUpper(string(key))

	str = keyCleaningRegex.ReplaceAllString(str, "")

	if str == "" {
		return errors.New("Invalid key. Must contain at least 1 letter")
	}

	return nil
}

// Encode encode the contents of the incoming reader stream
// and write the encoded message to the provided writer stream using this Cipher
// Returns the number of bytes read (and written) and any IO errors that may have occurred
func (me *Cipher) Encode(r io.Reader, w io.Writer) (int, error) {

	return processMessage(r, w, me.key, func(r rune, offset int) rune {
		return runeops.ShiftRune(r, offset)
	})
}

// Decode decode the contents of the incoming reader stream
// and write the decoded message to the provided writer stream using this Cipher
// Returns the number of bytes read (and written) and any IO errors that may have occurred
func (me *Cipher) Decode(r io.Reader, w io.Writer) (int, error) {

	return processMessage(r, w, me.key, func(r rune, offset int) rune {
		return runeops.ShiftRune(r, -offset)
	})
}

// processMessage wraps the key offset setup code into a single function for both encoding and decoding
// Returns the number of bytes read from the reader and any errors that occurred
func processMessage(r io.Reader, w io.Writer, key string, operation runeOperation) (int, error) {

	keyOffsets := getKeyOffsets(key)
	runeIndex := 0

	return runeops.ReadAllRunes(r, w, func(r rune) (rune, error) {

		result := operation(r, keyOffsets[runeIndex%len(keyOffsets)])
		runeIndex++
		return result, nil
	})
}

// getKeyOffsets convert a provided key word into a set of letter offsets for the cipher
func getKeyOffsets(key string) []int {

	offsets := make([]int, 0)

	for _, c := range key {
		offsets = append(offsets, getLetterOffset(c))
	}
	return offsets
}

// getLetterOffset gets the specified offset for a given letter
// since the keys for Vigenère ciphers are "cleaned" before this function gets called
// it is assumed the provided rune will be an uppercase English letter
func getLetterOffset(r rune) int {

	return (int(r) - 65) % 26
}

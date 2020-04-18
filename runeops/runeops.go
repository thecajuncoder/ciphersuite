package runeops

import (
	"bufio"
	"fmt"
	"io"
)

// Operation an operation performed on a single letter to transform it
type Operation func(r rune) (rune, error)

// ReadAllRunes reads the provided reader rune by rune,
// performs the specified operation on each rune,
// and writes the operation's results to the specified writer
// returns number of bytes read and any errors that occurred (besides io.EOF)
func ReadAllRunes(r io.Reader, w io.Writer, op Operation) (int, error) {

	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	defer bw.Flush()

	totalRead := 0

	var c, l rune
	var sz int
	var err error

	for c, sz, err = br.ReadRune(); err == nil; c, sz, err = br.ReadRune() {

		totalRead += sz

		if l, err = op(c); err != nil {
			return totalRead, fmt.Errorf("Operation error: %s", err.Error())
		}

		if _, err = bw.WriteRune(l); err != nil {
			return totalRead, fmt.Errorf("Write error: %s", err.Error())
		}
	}

	if err == io.EOF {
		return totalRead, nil
	}
	return totalRead, err
}

// ShiftRune shifts the value of an English letter rune by the specified offset
// Letters that overflow will wrap around to the opposite bounds of either capital or lowercase letters.
func ShiftRune(r rune, offset int) rune {

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

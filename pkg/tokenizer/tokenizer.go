package tokenizer

import (
	"bufio"
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
	"io"
	"unicode/utf8"
)

var _ iface.Tokenizer = &BufferedTokenizer{}

// BufferedTokenizer extracts tokens from an io.Reader using a bufferred scanner
type BufferedTokenizer struct {
	s *bufio.Scanner
}

// NewInclusiveTokenizer creates a new BufferedTokenizer that extracts tokens that are contiguous
// strings of runes for which incl returns true. Assumes tokens are readable as UTF-8.
func NewInclusiveTokenizer(r io.Reader, incl RuneFunc) *BufferedTokenizer {
	s := bufio.NewScanner(r)
	s.Split(split(IsLatinAlpha))
	return &BufferedTokenizer{
		s: s,
	}
}

// RuneFunc defines a function that performs a test on a given rune
type RuneFunc func(rune) bool

// Next obtains the next available token, returns nil if there is no remaining tokens or returns an error.
func (r *BufferedTokenizer) Next() ([]byte, error) {
	if r.s.Scan() {
		return copyBytes(r.s.Bytes()), nil
	}
	return nil, r.s.Err()
}

// IsLatinAlpha returns true for any alphabetic character in the latin alphabet
func IsLatinAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func split(incl RuneFunc) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Skip non-alphabetics
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if incl(r) {
				break
			}
		}

		// Scan until non-alphabetic, marking end of word.
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if !incl(r) {
				return i + width, data[start:i], nil
			}
		}

		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}

		// Request more data.
		return start, nil, nil
	}
}

func copyBytes(b []byte) []byte {
	bnew := make([]byte, len(b))
	copy(bnew, b)
	return bnew
}

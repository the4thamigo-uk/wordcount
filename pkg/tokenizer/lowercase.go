package tokenizer

import (
	"bytes"
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
)

type lowercase struct {
	src iface.Tokenizer
}

// ToLowerCase is a Tokenizer adapter to force tokens to lowercase
func ToLowerCase(src iface.Tokenizer) iface.Tokenizer {
	return &lowercase{
		src: src,
	}
}

// Next extracts a token from the source Tokenizer and convers it to lowercase
func (l *lowercase) Next() ([]byte, error) {
	b, err := l.src.Next()
	if b == nil || err != nil {
		return b, err
	}
	return bytes.ToLower(b), err
}

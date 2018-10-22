package counter

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/wordcount/pkg/hash"
	"github.com/the4thamigo-uk/wordcount/pkg/heap"
	"github.com/the4thamigo-uk/wordcount/pkg/tokenizer"
	"testing"
)

func TestCounter(t *testing.T) {
	data := bytes.NewBuffer([]byte("third first second third second third"))
	tok := tokenizer.NewInclusiveTokenizer(data, tokenizer.IsLatinAlpha)
	hs := hash.New(1024)
	hp := heap.New(20)

	items, err := CountTokens(tok, hs, hp)
	require.Nil(t, err)
	assert.Len(t, items, 3)
	assert.Equal(t, 3, items[0].Count())
	assert.Equal(t, 2, items[1].Count())
	assert.Equal(t, 1, items[2].Count())
}

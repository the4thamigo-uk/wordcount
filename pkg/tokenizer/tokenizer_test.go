package tokenizer

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBufferedTokenizer(t *testing.T) {
	data := bytes.NewBuffer([]byte("123first second1123third\n123 *fourth"))
	tk := NewInclusiveTokenizer(data, IsLatinAlpha)
	require.NotNil(t, tk)
	assert.Equal(t, "first", nextString(t, tk))
	assert.Equal(t, "second", nextString(t, tk))
	assert.Equal(t, "third", nextString(t, tk))
	assert.Equal(t, "fourth", nextString(t, tk))
}

func nextString(t *testing.T, tk *BufferedTokenizer) string {
	b, err := tk.Next()
	require.Nil(t, err)
	return string(b)
}

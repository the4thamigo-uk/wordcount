package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// testHash is a highly skewed hash function that is useful for testing
func testLenHash(b []byte) uint {
	return uint(len(b))
}

func TestHash_Add(t *testing.T) {
	h := New(2).WithHashFunc(testLenHash)
	c := h.AddString("a")
	assert.Equal(t, 1, c)
	assert.Equal(t, 1, h.CountString("a"))
	c = h.AddString("aa")
	assert.Equal(t, 1, c)
	assert.Equal(t, 1, h.CountString("aa"))
	c = h.AddString("a")
	assert.Equal(t, 2, c)
	assert.Equal(t, 2, h.CountString("a"))
}

func TestHash_CountNilIsZero(t *testing.T) {
	h := New(2).WithHashFunc(testLenHash)
	assert.Equal(t, 0, h.CountString("a"))
}

func TestHash_AddNil(t *testing.T) {
	h := New(2).WithHashFunc(testLenHash)
	c := h.Add(nil)
	assert.Equal(t, 0, c)
	assert.Equal(t, 0, h.Count(nil))
}

func TestHash_Counts(t *testing.T) {
	h := New(4).WithHashFunc(testLenHash)
	c := h.AddString("apple")
	assert.Equal(t, 1, c)
	c = h.AddString("banana")
	assert.Equal(t, 1, c)
	c = h.AddString("peach")
	assert.Equal(t, 1, c)
	c = h.AddString("grapefruit")
	assert.Equal(t, 1, c)
	c = h.AddString("apple")
	assert.Equal(t, 2, c)
	c = h.AddString("pear")
	assert.Equal(t, 1, c)
	c = h.AddString("satsuma")
	assert.Equal(t, 1, c)
	c = h.AddString("tangerine")
	assert.Equal(t, 1, c)

	counts := h.Counts()
	assert.Len(t, counts, 7)
}

package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRSHash_Nil(t *testing.T) {
	h := RSHash(nil)
	assert.Equal(t, uint(0), h)
}

func TestRSHash_Different(t *testing.T) {
	h1 := RSHash([]byte("apple"))
	h2 := RSHash([]byte("orange"))
	assert.NotEqual(t, h1, h2)
}

func TestRSHash_Same(t *testing.T) {
	h1 := RSHash([]byte("apple"))
	h2 := RSHash([]byte("orange"))
	assert.NotEqual(t, h1, h2)
}

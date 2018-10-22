package hash

import (
	"testing"
)

func TestRSHash_Nil(t *testing.T) {
	h := RSHash(nil)
	t.Log(h)
}

func TestRSHash_Values(t *testing.T) {
	h1 := RSHash([]byte("apple"))
	t.Log(h1)
	h2 := RSHash([]byte("orange"))
	t.Log(h2)
}

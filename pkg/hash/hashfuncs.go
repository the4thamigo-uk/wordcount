package hash

// Func defines the function prototype for a hash function
type Func func([]byte) uint

const (
	rsHashA uint = 63689
	rsHashB uint = 378551
)

// RSHash is an implementation of
// http://www.partow.net/programming/hashfunctions/#RSHashFunction
func RSHash(data []byte) uint {
	a := rsHashA
	b := rsHashB
	h := uint(0)
	for _, d := range data {
		h = h*a + uint(d)
		a *= b
	}
	return h
}

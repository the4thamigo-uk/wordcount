package iface

// Hash is a hash to count tokens
type Hash interface {
	Count(token []byte) int
	CountString(token string) int
	AddString(token string) int
	Add(token []byte) int
	Counts() []Item
}

// Heap is a token heap used to order tokens by their count
type Heap interface {
	Add(token []byte, count int)
	AddItem(item Item)
	AddString(token string, count int)
	Counts() []Item
}

// Item is a (token,count) pair
type Item interface {
	Token() []byte
	Count() int
	String() string
	Incr() int
}

// Tokenizer is used to extract tokens from a data source
type Tokenizer interface {
	Next() ([]byte, error)
}

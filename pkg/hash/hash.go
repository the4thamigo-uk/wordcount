package hash

import (
	"bytes"
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
	"github.com/the4thamigo-uk/wordcount/pkg/tokenitem"
)

// Hash is a hash table that counts occurrences of tokens.
type Hash struct {
	buckets []bucket
	hf      Func
}

type bucket []iface.Item

// New creates a new hash table with the given number of buckets
func New(numBuckets int) *Hash {
	return &Hash{
		buckets: make([]bucket, numBuckets),
		hf:      RSHash,
	}
}

// WithHashFunc allows specification of custom hash function. Default is RSHash.
func (h *Hash) WithHashFunc(hf Func) *Hash {
	h.hf = hf
	return h
}

func (h *Hash) index(token []byte) int {
	return int(h.hf(token) % uint(len(h.buckets)))
}

// CountString returns the count for the given token string
func (h *Hash) CountString(token string) int {
	return h.Count([]byte(token))
}

// Count returns the count for the given token
func (h *Hash) Count(token []byte) int {
	i := h.index(token)
	bucket := h.buckets[i]
	item := bucket.find(token)
	if item == nil {
		return 0
	}
	return item.Count()
}

// Counts returns the counts for all tokens in the hash
func (h *Hash) Counts() []iface.Item {
	var items []iface.Item
	for _, b := range h.buckets {
		items = append(items, b...)
	}
	return items
}

// AddString adds an occurence of the token string to the hash
func (h *Hash) AddString(token string) int {
	return h.Add([]byte(token))
}

// Add adds an occurence of the token to the hash
func (h *Hash) Add(token []byte) int {
	if len(token) == 0 {
		return 0
	}
	i := h.index(token)
	pbucket := &h.buckets[i]
	bucket := *pbucket
	item := bucket.find(token)
	if item != nil {
		item.Incr()
	} else {
		item = tokenitem.New(token, 1)
		*pbucket = append(bucket, item)
	}
	return item.Count()
}

func (b bucket) find(token []byte) iface.Item {
	for _, item := range b {
		if bytes.Equal(token, item.Token()) {
			return item
		}
	}
	return nil
}

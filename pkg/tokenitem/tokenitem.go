package tokenitem

import (
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
)

// Item is a (token,count) pair
type Item struct {
	token []byte
	count int
}

var _ iface.Item = &Item{}

// New creates a new Item
func New(token []byte, count int) *Item {
	return &Item{
		token: token,
		count: count,
	}
}

// Count retrieves the count for the token
func (h *Item) Count() int {
	return h.count
}

// Token retrieves the token
func (h *Item) Token() []byte {
	return h.token
}

// Token retrieves the token as a string
func (h *Item) String() string {
	return string(h.token)
}

// Incr increments the count by 1
func (h *Item) Incr() int {
	h.count++
	return h.count
}

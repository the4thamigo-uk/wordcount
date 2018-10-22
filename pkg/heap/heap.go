package heap

import (
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
	"github.com/the4thamigo-uk/wordcount/pkg/tokenitem"
)

var _ iface.Heap = &Heap{}

// Heap is a fixed-size heap to sort tokens using an insertion sort on the token count.
type Heap struct {
	items []iface.Item
}

// New creates a new Heap
func New(len int) *Heap {
	return &Heap{
		items: make([]iface.Item, 0, len),
	}
}

// AddItem adds a token and count to the heap.
// Items with tokens that already exist in the heap will be added.
func (h *Heap) AddItem(item iface.Item) {
	if item.Count() <= h.min() {
		h.append(item)
		return
	}
	if item.Count() >= h.max() {
		h.insert(item, 0)
		return
	}

	for i, hItem := range h.items {
		if item.Count() >= hItem.Count() {
			h.insert(item, i)
			break
		}
	}
}

// AddString adds a string token and count to the heap.
func (h *Heap) AddString(token string, count int) {
	h.Add([]byte(token), count)
}

// Add adds a token and count to the heap.
func (h *Heap) Add(token []byte, count int) {
	h.AddItem(tokenitem.New(token, count))
}

// Counts gets all tokens and their associated counts in order (high -> low) from the heap.
func (h *Heap) Counts() []iface.Item {
	var items []iface.Item
	return append(items, h.items...)
}

func (h *Heap) min() int {
	if len(h.items) > 0 {
		return h.items[len(h.items)-1].Count()
	}
	return 0
}

func (h *Heap) max() int {
	if len(h.items) > 0 {
		return h.items[0].Count()
	}
	return 0
}

func (h *Heap) hasSpace() bool {
	return len(h.items) < cap(h.items)
}

func (h *Heap) grow() bool {
	if h.hasSpace() {
		h.items = append(h.items, nil)
		return true
	}
	return false
}

func (h *Heap) insert(item iface.Item, pos int) {
	h.grow()
	h.shift(pos)
	h.items[pos] = item
}

func (h *Heap) append(item iface.Item) {
	if h.grow() {
		h.items[len(h.items)-1] = item
	}
}

func (h *Heap) shift(pos int) {
	if pos+1 < len(h.items) {
		copy(h.items[pos+1:], h.items[pos:len(h.items)-1])
	}
}

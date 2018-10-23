package heap

import (
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	"testing"
)

func TestHeap_Empty_Add(t *testing.T) {
	h := New(1)
	h.AddString("apple", 1)
	items := h.Counts()
	assert.Len(t, items, 1)
}

func TestHeap_SingleItem_AddGreaterThanMax(t *testing.T) {
	h := New(2)
	h.AddString("banana", 1)
	h.AddString("apple", 2)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_SingleItem_AddEqualToMinAndMax(t *testing.T) {
	h := New(2)
	h.AddString("apple", 1)
	h.AddString("banana", 1)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_SingleItem_AddLessThanMin(t *testing.T) {
	h := New(2)
	h.AddString("apple", 2)
	h.AddString("banana", 1)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_TwoItems_AddGreaterThanMax(t *testing.T) {
	h := New(3)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	h.AddString("apple", 3)
	items := h.Counts()
	assert.Len(t, items, 3)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
	assert.Equal(t, "orange", items[2].String())
}

func TestHeap_TwoItems_AddEqualToMax(t *testing.T) {
	h := New(3)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	h.AddString("apple", 2)
	items := h.Counts()
	assert.Len(t, items, 3)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
	assert.Equal(t, "orange", items[2].String())
}

func TestHeap_TwoItems_LessThanMax(t *testing.T) {
	h := New(3)
	h.AddString("apple", 3)
	h.AddString("orange", 1)
	h.AddString("banana", 2)
	items := h.Counts()
	assert.Len(t, items, 3)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
	assert.Equal(t, "orange", items[2].String())
}

func TestHeap_TwoItems_EqualToMin(t *testing.T) {
	h := New(3)
	h.AddString("apple", 2)
	h.AddString("banana", 1)
	h.AddString("orange", 1)
	items := h.Counts()
	assert.Len(t, items, 3)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
	assert.Equal(t, "orange", items[2].String())
}

func TestHeap_TwoItems_LessThanMin(t *testing.T) {
	h := New(3)
	h.AddString("apple", 3)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	items := h.Counts()
	assert.Len(t, items, 3)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
	assert.Equal(t, "orange", items[2].String())
}

func TestHeap_ExceedCapacity_AddGreaterThanMax(t *testing.T) {
	h := New(2)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	h.AddString("apple", 3)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_ExceedCapacity_AddEqualToMax(t *testing.T) {
	h := New(2)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	h.AddString("apple", 2)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_ExceedCapacity_LessThanMax(t *testing.T) {
	h := New(2)
	h.AddString("apple", 3)
	h.AddString("orange", 1)
	h.AddString("banana", 2)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_ExceedCapacity_EqualToMin(t *testing.T) {
	h := New(2)
	h.AddString("apple", 2)
	h.AddString("banana", 1)
	h.AddString("orange", 1)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

func TestHeap_ExceedCapacity_LessThanMin(t *testing.T) {
	h := New(2)
	h.AddString("apple", 3)
	h.AddString("banana", 2)
	h.AddString("orange", 1)
	items := h.Counts()
	assert.Len(t, items, 2)
	assert.Equal(t, "apple", items[0].String())
	assert.Equal(t, "banana", items[1].String())
}

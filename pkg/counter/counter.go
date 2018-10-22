package counter

import (
	"github.com/the4thamigo-uk/wordcount/pkg/iface"
)

// CountTokens reads the tokens from the tokenizer and inserts them in the hash for the purpose
// of counting, The tokens and their counts are then sorted using the heap.
func CountTokens(tok iface.Tokenizer, hash iface.Hash, heap iface.Heap) ([]iface.Item, error) {
	for {
		token, err := tok.Next()
		if err != nil {
			return nil, err
		}
		if token == nil {
			break
		}
		hash.Add(token)
	}
	for _, item := range hash.Counts() {
		heap.AddItem(item)
	}
	return heap.Counts(), nil
}

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/the4thamigo-uk/wordcount/pkg/counter"
	"github.com/the4thamigo-uk/wordcount/pkg/hash"
	"github.com/the4thamigo-uk/wordcount/pkg/heap"
	"github.com/the4thamigo-uk/wordcount/pkg/tokenizer"
	"os"
)

func main() {
	fn := pflag.StringP("input", "i", "", "Path to input data file")
	top := pflag.IntP("top", "t", 20, "Display top N words in file")
	buckets := pflag.IntP("heap", "h", 1024, "Number of buckets in hash")

	pflag.Parse()
	err := countWords(*fn, *top, *buckets)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func countWords(fn string, top int, buckets int) error {
	f, err := os.Open(fn)
	if err != nil {
		return errors.Wrapf(err, "Could not open file %s", fn)
	}
	tk := tokenizer.NewInclusiveTokenizer(f, tokenizer.IsLatinAlpha)
	ltk := tokenizer.ToLowerCase(tk)
	hs := hash.New(buckets)
	hp := heap.New(top)

	items, err := counter.CountTokens(ltk, hs, hp)
	if err != nil {
		return errors.Wrap(err, "Failed to count words in file")
	}

	for _, item := range items {
		fmt.Printf("%d %s\n", item.Count(), item.String())
	}
	return nil
}

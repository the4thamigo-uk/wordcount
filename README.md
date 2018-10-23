# wordcount

[![Build Status](https://travis-ci.org/the4thamigo-uk/wordcount.svg?branch=master)](https://travis-ci.org/the4thamigo-uk/wordcount?branch=master)
[![Coverage Status](https://coveralls.io/repos/the4thamigo-uk/wordcount/badge.svg?branch=master&service=github)](https://coveralls.io/github/the4thamigo-uk/wordcount?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/the4thamigo-uk/wordcount)](https://goreportcard.com/report/github.com/the4thamigo-uk/wordcount)


## Description

Counts tokens in a file without use of golang's map.

## Getting started

To build and run the example :

    go build ./cmd/wordcount && ./cmd/wordcount/wordcount -i ./data/mobydick.txt

The options are given by :

    ./cmd/wordcount/wordcount --help
    Usage of ./cmd/wordcount/wordcount:
      -h, --heap int       Number of buckets in hash (default 1024)
      -i, --input string   Path to input data file
      -t, --top int        Display top N words in file (default 20)

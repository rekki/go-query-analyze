# tokenize [![GitHub Actions Status](https://github.com/rekki/go-query-tokenize/workflows/test/badge.svg?branch=master)](https://github.com/rekki/go-query-tokenize/actions) [![codecov](https://codecov.io/gh/rekki/go-query-tokenize/branch/master/graph/badge.svg)](https://codecov.io/gh/rekki/go-query-tokenize) [![GoDoc](https://godoc.org/github.com/rekki/go-query-tokenize?status.svg)](https://godoc.org/github.com/rekki/go-query-tokenize)

> Simlpe tokenize chain

Example:

```go
package main

import t "github.com/rekki/go-query-analyze/tokenize"

func main() {
  tokenizer := []t.Tokenizer{
    t.NewWhitespace(),
    t.NewLeftEdge(1),
    t.NewUnique(),
  }
  tokens := t.Tokenize("hello world", tokenizer...)

  fmt.Printf("%v", tokens) // [h he hel hell hello w wo wor worl world]
}
```

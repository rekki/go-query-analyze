# normalize [![GitHub Actions Status](https://github.com/rekki/go-query-normalize/workflows/test/badge.svg?branch=master)](https://github.com/rekki/go-query-normalize/actions) [![codecov](https://codecov.io/gh/rekki/go-query-normalize/branch/master/graph/badge.svg)](https://codecov.io/gh/rekki/go-query-normalize) [![GoDoc](https://godoc.org/github.com/rekki/go-query-normalize?status.svg)](https://godoc.org/github.com/rekki/go-query-normalize)

> Simlpe normalize chain

Example:

```go
package main
import n "github.com/rekki/go-query-analyze/normalize"

func main() {
  normalize := []n.Normalizer{
    n.NewUnaccent(),
    n.NewLowerCase(),
    n.NewSpaceBetweenDigits(),
    n.NewCleanup(n.BASIC_NON_ALPHANUMERIC),
    n.NewTrim(" "),
  }
  normal := n.Normalize("Hęllö wÖrld. べぺ Ł2ł  ", normalize...)

  fmt.Printf("%s", normal) // hello world へへ l 2 l
}
```

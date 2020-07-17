package normalize

import (
	"strings"

	ps "github.com/blevesearch/go-porterstemmer"
)

type PorterStemmer struct {
	min int
}

// Create new porter stemmer that stems words with minimum length of minimumCharacters, others are left as is
func NewPorterStemmer(minimumCharacters int) *PorterStemmer {
	return &PorterStemmer{
		min: minimumCharacters,
	}
}

func (p *PorterStemmer) Apply(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	from := 0
	to := 0
	runned := []rune(s)

	for i, c := range runned {
		if c == ' ' {
			if to > from {
				chunk := runned[from:to]
				if len(chunk) >= p.min {
					for _, r := range ps.StemWithoutLowerCasing(chunk) {
						sb.WriteRune(r)
					}
					sb.WriteRune(' ')
				} else {
					sb.WriteString(string(chunk))
					sb.WriteRune(' ')
				}
				from = i + 1
				to = i + 1
			}
		} else {
			to++
		}
	}

	if from != to {
		chunk := runned[from:to]
		if len(chunk) >= p.min {
			for _, r := range ps.StemWithoutLowerCasing(runned[from:to]) {
				sb.WriteRune(r)
			}
			sb.WriteRune(' ')
		} else {
			sb.WriteString(string(chunk))
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}

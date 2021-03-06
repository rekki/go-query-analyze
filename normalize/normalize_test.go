package normalize

import "testing"

type TestCase struct {
	in  string
	out string
	n   []Normalizer
}

func testMany(t *testing.T, cases []TestCase) {
	for _, c := range cases {
		normalized := Normalize(c.in, c.n...)
		if normalized != c.out {
			t.Fatalf("in: <%s>, out: <%s>, expected: <%s>", c.in, normalized, c.out)
		}
	}

}

func TestSpaceBetweenDigits(t *testing.T) {
	cases := []TestCase{
		{
			in:  "c 111b 2c 22",
			out: "c 111 b 2 c 22",
			n:   []Normalizer{NewSpaceBetweenDigits()},
		},

		{
			in:  "1",
			out: "1",
			n:   []Normalizer{NewSpaceBetweenDigits()},
		},
		{
			in:  "a1b",
			out: "a 1 b",
			n:   []Normalizer{NewSpaceBetweenDigits()},
		},
		{
			in:  "",
			out: "",
			n:   []Normalizer{NewSpaceBetweenDigits()},
		},
		{
			in:  "9 abc 1b",
			out: "9 abc 1 b",
			n:   []Normalizer{NewSpaceBetweenDigits()},
		},
	}
	testMany(t, cases)
}

var dontOptimizeMe = 0

func BenchmarkRegexp(b *testing.B) {
	n := []Normalizer{NewCleanup(BASIC_NON_ALPHANUMERIC)}
	for i := 0; i < b.N; i++ {
		dontOptimizeMe += len(Normalize("c 1b!2&& にっぽん。。ぽ", n...))
	}
}
func BenchmarkRegexpEasy(b *testing.B) {
	n := []Normalizer{NewCleanup(BASIC_NON_ALPHANUMERIC)}
	for i := 0; i < b.N; i++ {
		dontOptimizeMe += len(Normalize("a b c", n...))
	}
}

func BenchmarkRemoveNonAlphanumeric(b *testing.B) {
	n := []Normalizer{NewRemoveNonAlphanumeric()}
	for i := 0; i < b.N; i++ {
		dontOptimizeMe += len(Normalize("c 1b!2&& にっぽん。。ぽ", n...))
	}
}

func BenchmarkRemoveNonAlphanumericEasy(b *testing.B) {
	n := []Normalizer{NewRemoveNonAlphanumeric()}
	for i := 0; i < b.N; i++ {
		dontOptimizeMe += len(Normalize("a b c", n...))
	}
}

func TestRegexp(t *testing.T) {
	cases := []TestCase{
		{
			in:  "c 1b!2&& にっぽん。。ぽ",
			out: "c 1b 2 にっぽん ぽ",
			n:   []Normalizer{NewCleanup(BASIC_NON_ALPHANUMERIC)},
		},
	}
	testMany(t, cases)
}

func TestNonAlphanumeric(t *testing.T) {
	cases := []TestCase{
		{
			in:  "c 1b!2&& にっぽん。。ぽ",
			out: "c 1b 2 にっぽん ぽ",
			n:   []Normalizer{NewRemoveNonAlphanumeric()},
		},
	}
	testMany(t, cases)
}

func TestNoop(t *testing.T) {
	cases := []TestCase{
		{
			in:  "c 1b!2&& にっぽん。。ぽ",
			out: "c 1b!2&& にっぽん。。ぽ",
			n:   []Normalizer{NewNoop()},
		},
	}
	testMany(t, cases)
}

func TestPorter(t *testing.T) {
	cases := []TestCase{
		{
			in:  "dogs hello cats",
			out: "dog hello cat ",
			n:   []Normalizer{NewPorterStemmer(0)},
		},
		{
			in:  "dogs",
			out: "dog ",
			n:   []Normalizer{NewPorterStemmer(0)},
		},
		{
			in:  "dogs   ",
			out: "dog ",
			n:   []Normalizer{NewPorterStemmer(0)},
		},

		{
			in:  "cats dogs onions",
			out: "cats dogs onion ",
			n:   []Normalizer{NewPorterStemmer(5)},
		},

		{
			in:  "cats dogs scallion onions       ",
			out: "cats dogs scallion onion ",
			n:   []Normalizer{NewPorterStemmer(5)},
		},

		{
			in:  "",
			out: "",
			n:   []Normalizer{NewPorterStemmer(0)},
		},
	}
	testMany(t, cases)
}

func TestUnaccent(t *testing.T) {
	cases := []TestCase{
		{
			in:  "ğüşöçİĞÜŞÖÇ にっ ぽん べぺぜじがぎゃぽhelloęĘŁłŚśŹźŃńä, ö or ü",
			out: "gusocIGUSOC にっ ほん へへせしかきゃほhelloeELlSsZzNna, o or u",
			n:   []Normalizer{NewUnaccent()},
		},
	}
	testMany(t, cases)
}

func TestLowerCase(t *testing.T) {
	cases := []TestCase{
		{
			in:  "AAğüşöçİĞÜŞÖÇ にっ ぽん べぺぜじがぎゃぽhelloęĘŁłŚśŹźŃńä, ö or ü",
			out: "aağüşöçiğüşöç にっ ぽん べぺぜじがぎゃぽhelloęęłłśśźźńńä, ö or ü",
			n:   []Normalizer{NewLowerCase()},
		},
	}
	testMany(t, cases)
}

func TestCustom(t *testing.T) {
	cases := []TestCase{
		{
			in:  "AAğüşöçİĞÜŞÖÇ にっ ぽん べぺぜじがぎゃぽhelloęĘŁłŚśŹźŃńä, ö or ü",
			out: "NOO AAgusocIGUSOC にっ ほん へへせしかきゃほhelloeELlSsZzNna, o or u",
			n: []Normalizer{NewUnaccent(), NewCustom(func(s string) string {
				return "NOO " + s
			})},
		},
	}
	testMany(t, cases)
}

func TestTrim(t *testing.T) {
	cases := []TestCase{
		{
			in:  "  AAğüşöçİĞÜŞÖÇ にっ ぽん べぺぜじがぎゃぽhelloęĘŁłŚśŹźŃńä, ö or ü !!!",
			out: "AAgusocIGUSOC にっ ほん へへせしかきゃほhelloeELlSsZzNna, o or u",
			n:   []Normalizer{NewTrim(" !"), NewUnaccent()},
		},
	}
	testMany(t, cases)
}

func TestCompose(t *testing.T) {
	cases := []TestCase{
		{
			in:  "AAğüşöç251İĞÜŞÖÇ にっ ぽん べぺぜ12じがぎゃぽhell2oęĘŁ2łŚśŹźŃńä, ö or ü",
			out: "aagusoc 251 igusoc にっ ほん へへせ 12 しかきゃほhell 2 oeel 2 lsszznna o or u",
			n:   []Normalizer{NewUnaccent(), NewLowerCase(), NewSpaceBetweenDigits(), NewCleanup(BASIC_NON_ALPHANUMERIC), NewTrim(" ")},
		},
	}
	testMany(t, cases)
}

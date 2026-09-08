package textstat

import (
	"strings"
	"unicode"
)

// WordCount считает частоты слов (буквы/цифры), регистр игнорируется.
func WordCount(s string) map[string]int {
	toKey := func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }
	tokens := strings.FieldsFunc(s, toKey)
	freq := make(map[string]int, len(tokens))
	for _, t := range tokens {
		k := strings.ToLower(t)
		freq[k]++
	}
	return freq
}

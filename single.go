package single

import "errors"

const (
	quote  = '\''
	escape = '\\'
)

// Quote returns a single-quoted Go string literal representing s. But, nothing else escapes.
func Quote(s string) string {
	out := []rune{quote}
	for _, r := range s {
		switch r {
		case quote:
			out = append(out, escape, r)
		default:
			out = append(out, r)
		}
	}
	out = append(out, quote)
	return string(out)
}

// Unquote interprets s as a single-quoted Go string literal, returning the string value that s quotes.
func Unquote(s string) (string, error) {
	if len(s) < 2 {
		return "", errors.New("invalid syntax")
	}
	if s[0] != quote || s[len(s)-1] != quote {
		return "", errors.New("invalid syntax")
	}
	out := []rune{}
	escaped := false
	for _, r := range s[1 : len(s)-1] {
		switch r {
		case escape:
			escaped = !escaped
			if !escaped {
				out = append(out, escape, escape)
			}
		case quote:
			if !escaped {
				return "", errors.New("invalid syntax")
			}
			out = append(out, r)
			escaped = false
		default:
			out = append(out, r)
			escaped = false
		}
	}
	return string(out), nil
}

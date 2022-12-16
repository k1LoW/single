package single

import "errors"

const (
	quote  = '\''
	escape = '\\'
)

func Quote(in string) string {
	out := []rune{quote}
	for _, r := range in {
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

func Unquote(in string) (string, error) {
	if len(in) < 2 {
		return "", errors.New("invalid syntax")
	}
	if in[0] != quote || in[len(in)-1] != quote {
		return "", errors.New("invalid syntax")
	}
	out := []rune{}
	escaped := false
	for _, r := range in[1 : len(in)-1] {
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

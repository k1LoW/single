package single_test

import (
	"testing"

	"github.com/k1LoW/single"
)

var tests = []struct {
	raw    string
	quoted string
}{
	{"", "''"},
	{"hello", "'hello'"},
	{"hello world", "'hello world'"},
	{`rock'n'roll`, `'rock\'n\'roll'`},
	{`"rock'n'roll"`, `'"rock\'n\'roll"'`},
	{`rock\\'n\\'roll`, `'rock\\\'n\\\'roll'`},
}

func TestQuote(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			got := single.Quote(tt.raw)
			if got != tt.quoted {
				t.Errorf("got %v\nwant %v", got, tt.quoted)
			}
		})
	}
}

func TestUnquote(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.quoted, func(t *testing.T) {
			got, err := single.Unquote(tt.quoted)
			if err != nil {
				t.Error(err)
			}
			if got != tt.raw {
				t.Errorf("got %v\nwant %v", got, tt.raw)
			}
		})
	}
}

func TestUnquoteError(t *testing.T) {
	tests := []struct {
		quoted  string
		wantErr bool
	}{
		{"hello world", true},
		{"'hello world", true},
		{"hello world'", true},
		{`''hello'`, true},
		{`'\'hello'`, false},
		{`'\\'hello'`, true},
		{`''`, false},
		{`h`, true},
	}
	for _, tt := range tests {
		t.Run(tt.quoted, func(t *testing.T) {
			_, err := single.Unquote(tt.quoted)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Error(err)
			}
			if tt.wantErr {
				t.Error("want error")
			}
		})
	}
}

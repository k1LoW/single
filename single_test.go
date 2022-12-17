package single_test

import (
	"fmt"
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

func ExampleQuote() {
	s := single.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)
	s = single.Quote(`rock'n'roll`)
	fmt.Println(s)
	// Output:
	// '"Fran & Freddie\'s Diner	☺"'
	// 'rock\'n\'roll'
}

func ExampleUnquote() {
	s, err := single.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)
	s, err = single.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)
	s, err = single.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)
	s, err = single.Unquote("'\u263a'")
	fmt.Printf("%q, %v\n", s, err)
	s, err = single.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err)
	s, err = single.Unquote("'\\'The string must be either single-quoted\\''")
	fmt.Printf("%q, %v\n", s, err)
	// Output:
	// "", invalid syntax
	// "", invalid syntax
	// "", invalid syntax
	// "☺", <nil>
	// "☹☹", <nil>
	// "'The string must be either single-quoted'", <nil>
}

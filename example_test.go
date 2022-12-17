package single_test

import (
	"fmt"

	"github.com/k1LoW/single"
)

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

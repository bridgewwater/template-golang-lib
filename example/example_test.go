package example

import (
	"fmt"
)

// To use this lib package, just
// the SetOutput function when your application starts.
func Example() {
	// Basic usage of this lib
	fooOut := Foo(1, 2)

	fmt.Printf("fooOut %v\n", fooOut)
	// Output:
	// fooOut 3
}

// other different examples such as bar naming can be written like this
// When parsing, the first letter of the Example_ content will be automatically capitalized.
func Example_bar() {
	// For some usage of example.Foo
	barOut := Foo(3, 4)

	fmt.Printf("barOut %v", barOut)
	// Output:
	// barOut 7
}

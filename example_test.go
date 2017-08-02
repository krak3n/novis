package novis_test

import (
	"fmt"

	"github.com/thisissoon/novis"
)

func ExampleRev() {
	foo := novis.Add("foo", "/foo")
	foo.Add("bar", "/bar/{id}", "{id}")
	fmt.Println(novis.Rev("foo.bar", "baz"))
	// Output: /foo/bar/baz
}

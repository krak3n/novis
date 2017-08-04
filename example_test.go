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

func ExampleGetBranch() {
	novis.Add("foo", "/foo")
	foo := novis.GetBranch("foo")
	fmt.Println(foo.Rel())
	// Output: /foo
}

func ExampleBranch_Rel() {
	foo := novis.Add("foo", "/foo")
	bar := foo.Add("bar", "/bar")
	fmt.Println(bar.Rel())
	// Output: /bar
}

func ExampleBranch_Add() {
	novis.Add("foo", "/foo")
	novis.Add("foo.bar", "/bar")
	novis.Add("foo.bar.baz", "/baz")
	fmt.Println(novis.Rev("foo.bar.baz"))
	// Output: /foo/bar/baz
}

func ExampleBranch_Path() {
	foo := novis.Add("foo", "/foo")
	bar := foo.Add("bar", "/bar")
	baz := bar.Add("baz", "/baz")
	fmt.Println(baz.Path())
	// Output: /foo/bar/baz
}

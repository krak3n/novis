/*
	The novis library is a named URL path building and reversing tool designed to
	compliment the HTTP router of your choice.

	You can use novis as a singleton:
		foo := novis.Add("foo", "/foo")
		foo.Add("bar", "/bar/:id", ":id")
		fmt.Println(novis.Rev("foo.bar", "baz"))
		// Output: "/foo/bar/baz"

	Or you can instantiate your own instance:
		nvs := novis.New()
		foo := nvs.Add("foo", "/foo")
		foo.Add("bar", "/bar/:id", ":id")
		fmt.Println(novis.Rev("foo.bar", "baz"))
		// Output: "/foo/bar/baz"

	Here is a simple, if contrived, example using github.com/go-chi/chi
		package main

		import (
			"fmt"
			"net/http"

			"github.com/go-chi/chi"
			"github.com/thisissoon/novis"
		)

		func main() {
			r := chi.NewMux()
			r.Get(novis.Add("foo", "/foo").Rel(), func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(fmt.Sprintf("Goto Bar: %s", novis.Rev("bar"))))
			})
			r.Get(novis.Add("bar", "/bar").Rel(), func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(fmt.Sprintf("Goto Foo: %s", novis.Rev("foo"))))
			})
			http.ListenAndServe(":5000", r)
		}

	The examples direcctory contains many more implementation examples.
*/
package novis

package main

import (
	"fmt"
	"net/http"

	"github.com/thisissoon/novis"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(novis.Add("foo", "/foo").Rel(), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Goto Bar: %s", novis.Rev("bar"))))
	})
	mux.HandleFunc(novis.Add("bar", "/bar").Rel(), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Goto Foo: %s", novis.Rev("foo"))))
	})
	http.ListenAndServe(":5000", mux)
}

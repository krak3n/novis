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

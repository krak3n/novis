# novis

[![GoDoc](https://godoc.org/github.com/tarm/serial?status.svg)](http://godoc.org/github.com/thisissoon/novis)
[![Build Status](https://travis-ci.org/thisissoon/novis.svg?branch=master)](https://travis-ci.org/thisissoon/novis)
[![Coverage Status](https://coveralls.io/repos/github/thisissoon/novis/badge.svg?branch=master)](https://coveralls.io/github/thisissoon/novis?branch=master)
[![Code Climate](https://codeclimate.com/github/thisissoon/novis/badges/gpa.svg)](https://codeclimate.com/github/thisissoon/novis)

A URL Reversing Library for dynamically building URL paths.

``` go
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
```

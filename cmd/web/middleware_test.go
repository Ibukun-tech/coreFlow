package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	h := NoSurf(testHandler)
	switch h.(type) {
	case http.Handler:

	default:
		t.Errorf("This is an error bro")
	}
}
func TestSessionLoad(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	h := SessionLoad(testHandler)
	switch h.(type) {
	case http.Handler:

	default:
		t.Errorf(fmt.Sprintf("It should return an http.Handler not %v", h))
	}
}

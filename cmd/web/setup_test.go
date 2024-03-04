package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

func (mH *myHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {

}

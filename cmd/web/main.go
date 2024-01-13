package main

import (
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
)

// its the Home Page handler
func main() {

	http.HandleFunc("/", Handler.Home)
	http.HandleFunc("/about", Handler.About())
	http.ListenAndServe(":2000", nil)
}

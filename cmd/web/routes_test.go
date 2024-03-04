package main

import (
	"net/http"
	"testing"

	"github.com/Ibukun-tech/coreFlow/pkg/config"
)

var testApp config.AppConfig

func TestRoutes(t *testing.T) {
	st := routes(&testApp)
	switch st.(type) {
	case http.Handler:
	default:
		t.Errorf("Its meant to be an handler")
	}
}

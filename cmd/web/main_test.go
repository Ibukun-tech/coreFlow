package main

import "testing"

func TestRun(t *testing.T) {
	err := Run()
	if err != nil {
		t.Errorf("Its not expected to bring an error")
	}
}

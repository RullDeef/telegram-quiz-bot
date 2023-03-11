package main

import "testing"

func TestSample(t *testing.T) {
	if someFunc(2) != 5 {
		t.Error("2 + 3 != 5")
	}
}

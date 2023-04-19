package main

import (
	"testing"
)

func TestCat(t *testing.T) {
	got := Cat([]string{"testdata/file1.txt"})
	want := "abcd\nefgh\nhijk\nlmnop"
	if got != want {
		t.Fatalf("expected %s but got %s", want, got)
	}
}

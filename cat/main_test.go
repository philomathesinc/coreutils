package main

import (
	"testing"
)

func TestCat(t *testing.T) {
	t.Run("one file", func(t *testing.T) {
		got := Cat([]string{"testdata/file1.txt"})
		want := "abcd\nefgh\nhijk\nlmnop"
		if got != want {
			t.Fatalf("expected \n%s but got \n%s", want, got)
		}
	})

	t.Run("multiple files", func(t *testing.T) {
		got := Cat([]string{"testdata/file1.txt", "testdata/file2.txt", "testdata/file3.txt"})
		want := "abcd\nefgh\nhijk\nlmnopqrst\nuvw\nxyz123\n4567\n890"
		if got != want {
			t.Fatalf("expected \n%s but got \n%s", want, got)
		}
	})

}

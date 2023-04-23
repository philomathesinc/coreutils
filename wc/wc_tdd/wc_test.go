package wctdd_test

import (
	"testing"

	wc "github.com/philomathesinc/coreutils/wc/wc_tdd"
)

func TestCountLines(t *testing.T) {
	t.Run("no new lines", func(t *testing.T) {
		got := wc.CountLines("hello world")
		want := 0

		if got != want {
			t.Errorf("CountLines() got: %+v, want: %+v", got, want)
		}
	})

	t.Run("one new line", func(t *testing.T) {
		got := wc.CountLines("hello world\n")
		want := 1

		if got != want {
			t.Errorf("CountLines() got: %+v, want: %+v", got, want)
		}
	})

	t.Run("only new lines", func(t *testing.T) {
		got := wc.CountLines("\n\n")
		want := 2

		if got != want {
			t.Errorf("CountLines() got: %+v, want: %+v", got, want)
		}
	})
}

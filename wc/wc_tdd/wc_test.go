package wctdd_test

import (
	"testing"

	wc "github.com/philomathesinc/coreutils/wc/wc_tdd"
)

func TestCountLines(t *testing.T) {
	got := wc.CountLines("hello world")
	want := 0

	if got != want {
		t.Errorf("CountLines() got: %+v, want: %+v", got, want)
	}
}

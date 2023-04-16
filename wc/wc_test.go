package wc_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/philomathesinc/coreutils/wc"
)

func TestCount(t *testing.T) {
	in := strings.NewReader(`abc
	abc
	pqr`)
	want := wc.Result{2, 3, 11, ""}
	wcObj := wc.New(true, true, true, in)
	got := wcObj.Count()

	if !cmp.Equal(got, want) {
		t.Errorf("expected %v but got %v\n%v", want, got, cmp.Diff(got, want))
	}
}

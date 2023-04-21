package cut_test

import (
	"testing"

	"github.com/philomathesinc/coreutils/cut"
)

func TestFields(t *testing.T) {
	t.Run("No delimiter", func(t *testing.T) {
		got := cut.Fields(`a:b
aa:bb:cc
aaa:bbb:ccc:ddd`, "2")
		want := `b
bb
bbb`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

package cut_test

import (
	"testing"

	"github.com/philomathesinc/coreutils/cut"
)

func TestFields(t *testing.T) {
	t.Run("No delimiter", func(t *testing.T) {
		fieldsTests := []struct {
			input  string
			fields string
			want   string
		}{
			{
				`a:b
aa:bb:cc
aaa:bbb:ccc:ddd`,
				"2",
				`b
bb
bbb`,
			},
		}

		for _, tt := range fieldsTests {
			got := cut.Fields(tt.input, tt.fields)
			want := tt.want

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})
}

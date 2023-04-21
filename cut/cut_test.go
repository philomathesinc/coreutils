package cut_test

import (
	"testing"

	"github.com/philomathesinc/coreutils/cut"
)

func TestFields(t *testing.T) {
	t.Run("No delimiter", noDelimiters(t))
}

func noDelimiters(t *testing.T) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		type args struct {
			input  string
			fields string
		}

		fieldsTests := []struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{
			{
				"letters",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"2",
				},
				"b\nbb\nbbb",
				false,
			},
			{
				"numbers",
				args{
					"1	2\n11	22	33\n111	222	333	444",
					"2",
				},
				"2\n22\n222",
				false,
			},
			{
				"unacceptable fields",
				args{
					"1	2\n11	22	33\n111	222	333	444",
					"a",
				},
				"",
				true,
			},
			{
				"field not present in all lines",
				args{
					"1	2\n11	22	33\n111	222	333	444",
					"3",
				},
				"33\n333",
				false,
			},
		}

		for _, tt := range fieldsTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := cut.Fields(tt.args.input, tt.args.fields)
				if (err != nil) != tt.wantErr {
					t.Errorf("Fields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Fields() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

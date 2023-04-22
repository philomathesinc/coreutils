package cut_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/philomathesinc/coreutils/cut"
)

func TestFields(t *testing.T) {
	t.Run("No delimiter", noDelimiters(t))
	t.Run("Various ranges", variousRanges(t))
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
				"happy path",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"2",
				},
				"b\nbb\nbbb",
				false,
			},
			{
				"unacceptable fields",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"a",
				},
				"",
				true,
			},
			{
				"field not present",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"3-3",
				},
				"cc\nccc",
				false,
			},
			{
				"field not present in all lines",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"3",
				},
				"\ncc\nccc",
				false,
			},
			{
				"field less than 1",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"0",
				},
				"",
				true,
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

func variousRanges(t *testing.T) func(t *testing.T) {
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
				"happy path",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"1-2",
				},
				"a	b\naa	bb\naaa	bbb",
				false,
			},
			{
				"start incorrect for range",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"a-3",
				},
				"",
				true,
			},
			{
				"end incorrect for range",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"1-a",
				},
				"",
				true,
			},
			{
				"start only range",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"3-",
				},
				"\ncc\nccc	ddd",
				false,
			},
			{
				"end only range",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"-2",
				},
				"a	b\naa	bb\naaa	bbb",
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
					t.Errorf("Fields() mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
				}
			})
		}
	}
}

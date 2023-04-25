package cut_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/philomathesinc/coreutils/cut"
)

func TestFields(t *testing.T) {
	t.Run("No delimiter", noDelimiters(t))
	t.Run("Types of ranges", typesOfRanges(t))
	t.Run("Delimiter specified", delimiterSpecified(t))
	t.Run("Multiple ranges", multipleRanges(t))
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
					"5",
				},
				"\n\n",
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
				got, err := cut.Fields(tt.args.input, tt.args.fields, "")
				if (err != nil) != tt.wantErr {
					t.Errorf("Fields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				diff := cmp.Diff(tt.want, got)
				if diff != "" {
					t.Errorf("Fields() mismatch (-want +got):\n%+v", diff)
				}
			})
		}
	}
}

func typesOfRanges(t *testing.T) func(t *testing.T) {
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
				"invalid range with no endpoint",
				args{
					"a	b\naa	bb	cc\naaa	bbb	ccc	ddd",
					"-",
				},
				"",
				true,
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
				got, err := cut.Fields(tt.args.input, tt.args.fields, "")
				if (err != nil) != tt.wantErr {
					t.Errorf("Fields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				diff := cmp.Diff(tt.want, got)
				if diff != "" {
					t.Errorf("Fields() mismatch (-want +got):\n%+v", diff)
				}
			})
		}
	}
}

func delimiterSpecified(t *testing.T) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		type args struct {
			input     string
			fields    string
			delimiter string
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
					"a:b\naa:bb:cc\naaa:bbb:ccc:ddd",
					"1-2",
					":",
				},
				"a:b\naa:bb\naaa:bbb",
				false,
			},
			{
				"mixed delimiters in input",
				args{
					"a	b\naa	bb	cc\naaa:bbb:ccc:ddd",
					"1-2",
					":",
				},
				"a	b\naa	bb	cc\naaa:bbb",
				false,
			},
		}

		for _, tt := range fieldsTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := cut.Fields(tt.args.input, tt.args.fields, tt.args.delimiter)
				if (err != nil) != tt.wantErr {
					t.Errorf("Fields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				diff := cmp.Diff(tt.want, got)
				if diff != "" {
					t.Errorf("Fields() mismatch (-want +got):\n%+v", diff)
				}
			})
		}
	}
}

func multipleRanges(t *testing.T) func(t *testing.T) {
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
					"1,3",
				},
				"a\naa	cc\naaa	ccc",
				false,
			},
		}

		for _, tt := range fieldsTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := cut.Fields(tt.args.input, tt.args.fields, "")
				if (err != nil) != tt.wantErr {
					t.Errorf("Fields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				diff := cmp.Diff(tt.want, got)
				if diff != "" {
					t.Errorf("Fields() mismatch (-want +got):\n%+v", diff)
				}
			})
		}
	}
}

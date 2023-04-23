package wctdd_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	wc "github.com/philomathesinc/coreutils/wc/wc_tdd"
)

func TestCountLines(t *testing.T) {
	type args struct {
		lines string
	}
	countLinesTests := []struct {
		name string
		args args
		want int
	}{
		{
			"no new lines",
			args{
				"hello world",
			},
			0,
		},
		{
			"one new line",
			args{
				"hello world\n",
			},
			1,
		},
		{
			"only new lines",
			args{
				"\n\n",
			},
			2,
		},
	}

	for _, tt := range countLinesTests {
		t.Run(tt.name, func(t *testing.T) {
			got := wc.CountLines(tt.args.lines)
			want := tt.want
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf("CountLines() mismatch (-want +got):\n%+v", diff)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	type args struct {
		words string
	}
	countWordsTests := []struct {
		name string
		args args
		want int
	}{
		{
			"two new words - happy path",
			args{
				"hello world",
			},
			2,
		},
	}

	for _, tt := range countWordsTests {
		t.Run(tt.name, func(t *testing.T) {
			got := wc.CountWords(tt.args.words)
			want := tt.want
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf("CountWords() mismatch (-want +got):\n%+v", diff)
			}
		})
	}
}

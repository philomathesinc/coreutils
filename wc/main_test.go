package main_test

import (
	"testing"
	. "github.com/philomathesinc/coreutils/wc"
)

func TestCountLines(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"count lines valid",
			args{"testdata/one.txt"},
			3,
			false,
		},
		{
			"count lines valid",
			args{"testdata/two.txt"},
			6,
			false,
		},
		{
			"count lines file does not exist",
			args{"testdata/three.txt"},
			0,
			true,
		},
		{
			"count lines read error",
			args{"testdata/no-read.txt"},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLines(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

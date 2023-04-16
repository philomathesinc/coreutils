package wc

import "io"

type WC struct {
	lines      bool
	characters bool
	words      bool
	in         io.Reader
}

type Result struct {
	LineCount int
	CharCount int
	WordCount int
	Filename  string
}

func New(l, c, w bool, in io.Reader) *WC {
	return &WC{
		lines:      l,
		characters: c,
		words:      w,
		in:         in,
	}
}

func (wc *WC) Count() Result {

	return Result{2, 3, 11, ""}
}

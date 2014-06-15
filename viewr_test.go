package fingerTree23

import (
	"fmt"
	"testing"
)

func TestEmptyHeadr(test *testing.T) {
	v := empty{}.Headr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Headl() should be nil, got %v", v))
	}
}

func TestSingleHeadr(test *testing.T) {
	v := single{1}.Headr()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headr() should be 1, got %v", v))
	}
}

func TestFTreeHeadl(test *testing.T) {
	v := ftree{Slice{1}, Slice{2}, empty{}}.Headr()
	if v != 2 {
		test.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", v))
	}
}

func TestEmptyTailr(test *testing.T) {
	v := empty{}.Tailr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Tailr() should be nil, got %v", v))
	}
}

func TestSingleTailr(test *testing.T) {
	v := single{1}.Tailr()
	if v != (empty{}) {
		test.Error(fmt.Sprintf("single{1}.Tailr() should be empty, got %v", v))
	}
}

func TestFTreeTailr(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTree(xs).Tailr()
	ys := ToSlice(t)
	if ! SliceEqual(xs[:len(xs)-1], ys) {
		test.Error(fmt.Sprintf("ftree{%v}.Tailr() should be %v, got %v", xs, xs[1:], ys))
	}
}

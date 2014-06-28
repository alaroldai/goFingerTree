package fingerTree

import (
	"testing"
	"fmt"
)

func ToSliceL(s Iterable) Slice {
	var res = Slice{}
	s.Iterl(func(item Any) {
		res = res.Pushr(item)
	})
	return res
}

func ToSliceR(s Iterable) Slice {
	var res = Slice{}
	s.Iterr(func(item Any) {
		res = res.Pushl(item)
	})
	return res
}

func Iter_TestL(test *testing.T, it Iterable, x Slice) {
	r := ToSliceL(it)
	if ! SliceEqual(r, x) {
		test.Error(fmt.Sprintf("Iter_TestL: ToSliceL(%v) should be %v, got %v", it, x, r))
	}
}

func Iter_TestR(test *testing.T, it Iterable, x Slice) {
	r := ToSliceR(it)
	if ! SliceEqual(r, x) {
		test.Error(fmt.Sprintf("Iter_TestR: ToSliceR(%v) should be %v, got %v", it, x, r))
	}
}

func ToSliceLF(s Foldable) Slice {
	return s.Foldl(func(acc Any, item Any) Any {
		return acc.(Slice).Pushr(item)
	}, Slice{}).(Slice)
}

func ToSliceRF(s Foldable) Slice {
	return s.Foldr(func(acc Any, item Any) Any {
		return acc.(Slice).Pushl(item)
	}, Slice{}).(Slice)
}

func Fold_TestL(test *testing.T, it Foldable, x Slice) {
	r := ToSliceLF(it)
	if ! SliceEqual(r, x) {
		test.Error(fmt.Sprintf("Fold_TestL: ToSliceLF(%v) should be %v, got %v", it, x, r))
	}
}

func Fold_TestR(test *testing.T, it Foldable, x Slice) {
	r := ToSliceRF(it)
	if ! SliceEqual(r, x) {
		test.Error(fmt.Sprintf("Fold_TestR: ToSliceRF(%v) should be %v, got %v", it, x, r))
	}
}

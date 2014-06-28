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

func Iter_Test(test *testing.T, it Iterable, x Slice) {
	Iter_TestL(test, it, x)
	Iter_TestR(test, it, x)
}

func Iter_TestL(test *testing.T, it Iterable, x Slice) {
	r := ToSliceL(it)
	Slice_TestM(test, r, x, fmt.Sprintf("ToSliceL(%v)", it))
}

func Iter_TestR(test *testing.T, it Iterable, x Slice) {
	r := ToSliceR(it)
	Slice_TestM(test, r, x, fmt.Sprintf("ToSliceR(%v)", it))
}

func Fold_Test(test *testing.T, it Foldable, x Slice) {
	Fold_TestL(test, it, x)
	Fold_TestR(test, it, x)
}

func Fold_TestL(test *testing.T, it Foldable, x Slice) {
	r := ToSliceLF(it)
	Slice_TestM(test, r, x, fmt.Sprintf("ToSliceLF(%v)", it))
}

func Fold_TestR(test *testing.T, it Foldable, x Slice) {
	r := ToSliceRF(it)
	Slice_TestM(test, r, x, fmt.Sprintf("ToSliceRF(%v)", it))
}

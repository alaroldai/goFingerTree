package fingerTree

import (
	"reflect"
	"testing"
	"fmt"
)

func TestSliceImplementsFoldable(test *testing.T) {
	stype := reflect.TypeOf(Slice{})
	itype := reflect.TypeOf((*Foldable)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSliceImplementsIterable(test *testing.T) {
	stype := reflect.TypeOf(Slice{})
	itype := reflect.TypeOf((*Iterable)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSliceImplementsMeasured(test *testing.T) {
	stype := reflect.TypeOf(Slice{})
	itype := reflect.TypeOf((*Measured)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSliceImplementsMonoid(test *testing.T) {
	stype := reflect.TypeOf(Slice{})
	itype := reflect.TypeOf((*Monoid)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSliceIter(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Iter_Test(test, x, x)
}

func TestSliceFold(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Fold_Test(test, x, x)
}

func Slice_TestM(test *testing.T, x Any, y Slice, msg string) {
	s, success := x.(Slice)
	if ! (success && SliceEqual(s, y)) {
		test.Error(fmt.Sprintf("%s should be %v, got %#v", msg, y, x))
	}
}

func TestSlicePlus(test *testing.T) {
	x := Slice{1, 2}
	y := Slice{3, 4}

	Slice_TestM(test, x.Plus(y), Slice{1, 2, 3, 4}, "[1 2] + [3 4]")
}

func TestSliceMeasure(test *testing.T) {
	x := Slice{mfree{1}, mfree{2}, mfree{3}, mfree{4}}

	Slice_TestM(test, x.Measure(), Slice{1, 2, 3, 4}, "Measure(Slice{[1] [2] [3] [4]})")
}

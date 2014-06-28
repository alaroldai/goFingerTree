package fingerTree

import (
	"reflect"
	"testing"
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

func TestSliceIterL(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Iter_TestL(test, x, x)
}

func TestSliceIterR(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Iter_TestR(test, x, x)
}

func TestSliceFoldL(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Fold_TestL(test, x, x)
}

func TestSliceFoldR(test *testing.T) {
	x := Slice{1, 2, 3, 4, 5}
	Fold_TestR(test, x, x)
}

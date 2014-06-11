package fingerTree23

import (
	"fmt"
	"testing"
)


func ToSliceR(t Foldable) []Any {
	app := func(a Any, acc Any) Any {
		slice := acc.([]Any)
		return append([]Any{a}, slice...)
	}
	return t.Foldr(app, make([]Any, 0)).([]Any)
}

func SliceEqual(a, b []Any) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func test_foldr(test *testing.T, name string, box Foldable, values []Any) {
	unfolded := ToSliceR(box)
	if ! SliceEqual(unfolded, values) {
		test.Error(fmt.Sprintf("ToSliceR(%s): expected %v, got %v", name, values, unfolded))
	}
}

func TestNode2Foldr(test *testing.T) {
	test_foldr(test, "node2{1,2}", node2{[2]Any{1, 2}}, []Any{1,2})
}

func TestNode3Foldr(test *testing.T) {
	test_foldr(test, "node2{1,2,3}", node3{[3]Any{1,2,3}}, []Any{1,2,3})
}

func TestEmptyFoldr(test *testing.T) {
	test_foldr(test, "empty{}", empty{}, []Any{})
}

func TestSingleFoldr(test *testing.T) {
	test_foldr(test, "single{1}", single{1}, []Any{1})
}

func TestFTreeFoldr(test *testing.T) {
	tree := ftree{
		[]Any{1,2},
		[]Any{6,7},
		single{node3{[3]Any{3,4,5}}},
	}

	test_foldr(test, "ftree{[1,2] [3 4 5] [6 7]}", tree, []Any{1,2,3,4,5,6,7})
}

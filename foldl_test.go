package fingerTree23

import (
	"fmt"
	"testing"
)

func ToSliceL(t Foldable) []Any {
	app := func(a Any, b Any) Any {
		return append(a.([]Any), b)
	}
	return t.Foldl(app, []Any{}).([]Any)
}

func test_foldl(test *testing.T, name string, box Foldable, values []Any) {
	unfolded := ToSliceL(box)
	if ! SliceEqual(unfolded, values) {
		test.Error(fmt.Sprintf("ToSliceL(%s): expected %v, got %v", name, values, unfolded))
	}
}

func TestNode2Foldl(test *testing.T) {
	test_foldl(test, "node2{1,2}", node2{[2]Any{1, 2}}, []Any{1,2})
}

func TestNode3Foldl(test *testing.T) {
	test_foldl(test, "node2{1,2,3}", node3{[3]Any{1,2,3}}, []Any{1,2,3})
}

func TestEmptyFoldl(test *testing.T) {
	test_foldl(test, "empty{}", empty{}, []Any{})
}

func TestSingleFoldl(test *testing.T) {
	test_foldl(test, "single{1}", single{1}, []Any{1})
}

func TestFTreeFoldl(test *testing.T) {
	tree := ftree{
		[]Any{1,2},
		[]Any{6,7},
		single{node3{[3]Any{3,4,5}}},
	}

	test_foldl(test, "ftree{[1,2] [3 4 5] [6 7]}", tree, []Any{1,2,3,4,5,6,7})
}

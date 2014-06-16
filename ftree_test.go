package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFTreeImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(ftree{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestFTreeFoldl(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 45 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func TestFTreeFoldr(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushr(i)
	}

	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 45 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func TestFTreePushl(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushl(i)
	}

	if cmpslices(ToSlice(n), []Any{7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [7 6 5 4 3 2 1 0], got %v", ToSlice(n)))
	}
}

func TestFTreePopl(test *testing.T) {
	var n FingerTree = &empty{}

	for i := 0; i < 20; i++ {
		n = n.Pushl(i)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.Popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushr(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreePushr(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 8; i++ {
		n = n.Pushr(i)
	}

	if cmpslices(ToSlice(n), []Any{0, 1, 2, 3, 4, 5, 6, 7}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [0 1 2 3 4 5 6 7], got %v", ToSlice(n)))
	}
}

func TestFTreeIterl(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	add := func(d Any) {
		sum += d.(int)
	}
	n.Iterl(add)
	if sum != 45 {
		test.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}

func TestFTreeIterr(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	add := func(d Any) {
		sum += d.(int)
	}
	n.Iterr(add)
	if sum != 45 {
		test.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}

func TestFTreeHeadr(test *testing.T) {
	v := ftree{Slice{1}, Slice{2}, empty{}}.Headr()
	if v != 2 {
		test.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", v))
	}
}

func TestFTreeTailr(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTree(xs).Tailr()
	ys := ToSlice(t)
	expected := xs[:len(xs)-1]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftree{%v}.Tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeHeadl(test *testing.T) {
	v := ftree{Slice{1}, Slice{2}, empty{}}.Headl()
	if v != 1 {
		test.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 1, got %v", v))
	}
}

func TestFTreeTaill(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTree(xs).Taill()
	ys := ToSlice(t)
	expected := xs[1:]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftree{%v}.Tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeIsEmpty(test *testing.T) {
	v := ToFingerTree(Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if v.IsEmpty() {
		test.Error("Expected isEmpty to be false")
	}
}

package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSinglePushl(test *testing.T) {
	n := &single{1}
	r := n.Pushl(2)
	if cmpslices(ToSlice(r), []Any{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushl(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestSinglePopl(test *testing.T) {
	n := &single{1}
	r, e := n.Popl()
	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error("Expected n.Popl() result to be an empty node")
	}
	if e != 1 {
		test.Error("Expected n.Popl() result to be 1")
	}
}

func TestSinglePushr(test *testing.T) {
	n := &single{1}
	r := n.Pushr(2)
	if cmpslices(ToSlice(r), []Any{1, 2}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushr(2) to result in sequence [1 2], got %v", ToSlice(r)))
	}
}

func TestSingleImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(single{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSingleFoldl(test *testing.T) {
	n := &single{1}
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 1 {
		test.Error("Expected n.Foldl to return 1, got " + string(r.(int)))
	}
}

func TestSingleFoldr(test *testing.T) {
	n := &single{1}
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 1 {
		test.Error("Expected n.Foldr to return 1, got " + string(r.(int)))
	}
}

func TestSingleIterr(test *testing.T) {
	n := &single{1}
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 1 {
		test.Error("Expected n.Iterr to return 1, got " + string(sum))
	}
}

func TestSingleIterl(test *testing.T) {
	n := &single{1}
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 1 {
		test.Error("Expected n.Iterl to return 1, got " + string(sum))
	}
}

func TestSingleHeadr(test *testing.T) {
	v := single{1}.Headr()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headr() should be 1, got %v", v))
	}
}

func TestSingleTailr(test *testing.T) {
	v := single{1}.Tailr()
	if v != (empty{}) {
		test.Error(fmt.Sprintf("single{1}.Tailr() should be empty, got %v", v))
	}
}

func TestSingleHeadl(test *testing.T) {
	v := single{1}.Headl()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headl() should be 1, got %v", v))
	}
}

func TestSingleTaill(test *testing.T) {
	v := single{1}.Taill()
	if v != (empty{}) {
		test.Error(fmt.Sprintf("single{1}.Taill() should be empty, got %v", v))
	}
}

func TestSingleIsEmpty(test *testing.T) {
	v := &single{1}
	if v.IsEmpty() {
		test.Error("Expected &single{1}.IsEmpty() to be false")
	}
}

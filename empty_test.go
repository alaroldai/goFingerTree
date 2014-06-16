package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(empty{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestEmptyFoldl(test *testing.T) {
	n := &empty{}
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldl(add, 0)
	if r != 0 {
		test.Error("Expected n.Foldl to return 0, got " + string(r.(int)))
	}
}

func TestEmptyFoldr(test *testing.T) {
	n := &empty{}
	add := func(a Any, b Any) Any {
		return Any(a.(int) + b.(int))
	}
	r := n.Foldr(add, 0)
	if r != 0 {
		test.Error("Expected n.Foldr to return 0, got " + string(r.(int)))
	}
}

func TestEmptyIterr(test *testing.T) {
	n := &empty{}
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 0 {
		test.Error("Expected n.Iterr to return 0, got " + string(sum))
	}
}

func TestEmptyIterl(test *testing.T) {
	n := &empty{}
	sum := 0
	add := func(b Any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 0 {
		test.Error("Expected n.Iterl to return 0, got " + string(sum))
	}
}

func TestEmptyPushl(test *testing.T) {
	v := empty{}.Pushl(1)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushl(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyPopl(test *testing.T) {
	n := empty{}
	r, e := n.Popl()

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.Popl() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.Popl() result to be nil")
	}
}

func TestEmptyPushr(test *testing.T) {
	v := empty{}.Pushr(1)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushr(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyHeadr(test *testing.T) {
	v := empty{}.Headr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Headl() should be nil, got %v", v))
	}
}

func TestEmptyTailr(test *testing.T) {
	v := empty{}.Tailr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Tailr() should be nil, got %v", v))
	}
}

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
		return append(a.(Slice), b)
	}
	r := n.Foldl(add, Slice{})
	if !cmpslices(r.(Slice), Slice{}) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", Slice{}, r))
	}
}

func TestEmptyFoldr(test *testing.T) {
	n := &empty{}
	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldr(add, Slice{})
	if !cmpslices(r.(Slice), Slice{}) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", Slice{}, r))
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

func TestEmptyHeadl(test *testing.T) {
	v := empty{}.Headl()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Headl() should be nil, got %v", v))
	}
}

func TestEmptyTaill(test *testing.T) {
	v := empty{}.Taill()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Tailr() should be nil, got %v", v))
	}
}

func TestEmptyIsEmpty(test *testing.T) {
	v := empty{}
	if !v.IsEmpty() {
		test.Error("Expected &empty{}.IsEmpty() to be true")
	}
}

func TestEmptyConcatl(test *testing.T) {
	e := &empty{}
	s := e.Pushl(1)
	t := s.Pushl(2)

	o := &empty{}

	if !cmpslices(ToSlice(t), ToSlice(e.Concatl(t))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(t), ToSlice(e.Concatl(t))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.Concatl(s))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(s), ToSlice(e.Concatl(s))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.Concatl(o))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(o), ToSlice(e.Concatl(o))))
	}
}

func TestEmptyConcatr(test *testing.T) {
	e := &empty{}
	s := e.Pushl(1)
	t := s.Pushl(2)

	o := e

	if !cmpslices(ToSlice(t), ToSlice(e.Concatr(t))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(t), ToSlice(e.Concatr(t))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.Concatr(s))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(s), ToSlice(e.Concatr(s))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.Concatr(o))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(o), ToSlice(e.Concatr(o))))
	}
}

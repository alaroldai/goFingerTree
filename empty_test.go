package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyImplementsFingerTreeComponent(test *testing.T) {
	stype := reflect.TypeOf(makeEmpty())
	itype := reflect.TypeOf((*FingerTreeComponent)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestEmptyFoldl(test *testing.T) {
	n := makeEmpty()
	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldl(add, Slice{})
	if !cmpslices(r.(Slice), Slice{}) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", Slice{}, r))
	}
}

func TestEmptyFoldr(test *testing.T) {
	n := makeEmpty()
	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldr(add, Slice{})
	if !cmpslices(r.(Slice), Slice{}) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", Slice{}, r))
	}
}

func TestEmptyIterr(test *testing.T) {
	n := makeEmpty()
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
	n := makeEmpty()
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
	v := (makeEmpty()).Pushl(1, mdataStandardTypes)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushl(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyPopl(test *testing.T) {
	n := makeEmpty()
	r, e := n.Popl(mdataStandardTypes)

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.Popl() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.Popl() result to be nil")
	}
}

func TestEmptyPopr(test *testing.T) {
	n := makeEmpty()
	r, e := n.Popr(mdataStandardTypes)

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.Popr() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.Popr() result to be nil")
	}
}

func TestEmptyPushr(test *testing.T) {
	v := (makeEmpty()).Pushr(1, mdataStandardTypes)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.Pushr(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyHeadr(test *testing.T) {
	v := (makeEmpty()).Headr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Headl() should be nil, got %v", v))
	}
}

func TestEmptyTailr(test *testing.T) {
	v := (makeEmpty()).Tailr(mdataStandardTypes)
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Tailr() should be nil, got %v", v))
	}
}

func TestEmptyHeadl(test *testing.T) {
	v := (makeEmpty()).Headl()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Headl() should be nil, got %v", v))
	}
}

func TestEmptyTaill(test *testing.T) {
	v := (makeEmpty()).Taill(mdataStandardTypes)
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.Tailr() should be nil, got %v", v))
	}
}

func TestEmptyIsEmpty(test *testing.T) {
	v := makeEmpty()
	if !v.IsEmpty() {
		test.Error("Expected makeEmpty().IsEmpty() to be true")
	}
}

func TestEmptyConcatl(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1, mdataStandardTypes)
	t := s.Pushl(2, mdataStandardTypes)

	o := makeEmpty()

	if !cmpslices(ToSlice(t), ToSlice(e.Concatl(t, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(t), ToSlice(e.Concatl(t, mdataStandardTypes))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.Concatl(s, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(s), ToSlice(e.Concatl(s, mdataStandardTypes))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.Concatl(o, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatl to return %v, got %v", ToSlice(o), ToSlice(e.Concatl(o, mdataStandardTypes))))
	}
}

func TestEmptyConcatr(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1, mdataStandardTypes)
	t := s.Pushl(2, mdataStandardTypes)

	o := e

	if !cmpslices(ToSlice(t), ToSlice(e.Concatr(t, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(t), ToSlice(e.Concatr(t, mdataStandardTypes))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.Concatr(s, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(s), ToSlice(e.Concatr(s, mdataStandardTypes))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.Concatr(o, mdataStandardTypes))) {
		test.Error(fmt.Sprintf("Expected e.Concatr to return %v, got %v", ToSlice(o), ToSlice(e.Concatr(o, mdataStandardTypes))))
	}
}

func TestEmptyFTSize(test *testing.T) {
	e := makeEmpty()
	if e.mdataForKey(ft_size_key, mdataStandardTypes).(int) != 0 {
		test.Error(fmt.Sprintf("Expected e.ft_size to equal 0, got %v", e.mdataForKey(ft_size_key, mdataStandardTypes).(int)))
	}
}

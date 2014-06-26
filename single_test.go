package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSinglePushl(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
	r := n.Pushl(2, mdataStandardTypes)
	if cmpslices(ToSlice(r), []Any{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushl(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestSinglePopl(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
	r, e := n.Popl(mdataStandardTypes)
	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error("Expected n.Popl() result to be an empty node")
	}
	if e != 1 {
		test.Error("Expected n.Popl() result to be 1")
	}
}

func TestSinglePushr(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
	r := n.Pushr(2, mdataStandardTypes)
	if cmpslices(ToSlice(r), []Any{1, 2}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushr(2) to result in sequence [1 2], got %v", ToSlice(r)))
	}
}

func TestSingleImplementsFingerTreeComponent(test *testing.T) {
	stype := reflect.TypeOf(&single{})
	itype := reflect.TypeOf((*FingerTreeComponent)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSingleFoldl(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldl(add, Slice{})
	if !cmpslices(r.(Slice), Slice{1}) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", Slice{1}, r))
	}
}

func TestSingleFoldr(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldr(add, Slice{})
	if !cmpslices(r.(Slice), Slice{1}) {
		test.Error(fmt.Sprintf("Expected n.Foldr to return %v, got %v", Slice{1}, r))
	}
}

func TestSingleIterr(test *testing.T) {
	n := makeSingle(1, mdataStandardTypes)
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
	n := makeSingle(1, mdataStandardTypes)
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
	v := (makeSingle(1, mdataStandardTypes)).Headr()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headr() should be 1, got %v", v))
	}
}

func TestSingleTailr(test *testing.T) {
	v := (makeSingle(1, mdataStandardTypes)).Tailr(mdataStandardTypes)
	if !v.IsEmpty() {
		test.Error(fmt.Sprintf("single{1}.Tailr() should be empty, got %v", v))
	}
}

func TestSingleHeadl(test *testing.T) {
	v := (makeSingle(1, mdataStandardTypes)).Headl()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headl() should be 1, got %v", v))
	}
}

func TestSingleTaill(test *testing.T) {
	v := (makeSingle(1, mdataStandardTypes)).Taill(mdataStandardTypes)
	if !v.IsEmpty() {
		test.Error(fmt.Sprintf("single{1}.Taill() should be empty, got %v", v))
	}
}

func TestSingleIsEmpty(test *testing.T) {
	v := makeSingle(1, mdataStandardTypes)
	if v.IsEmpty() {
		test.Error("Expected makeSingle(1).IsEmpty() to be false")
	}
}

func TestSingleConcatl(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1, mdataStandardTypes)
	t := s.Pushl(2, mdataStandardTypes)
	o := (makeEmpty()).Pushl(3, mdataStandardTypes)

	expected := append(ToSlice(t), ToSlice(s)...)
	r := s.Concatl(t, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(t, mdataStandardTypes))))
	}

	expected = append(ToSlice(o), ToSlice(s)...)
	r = s.Concatl(o, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(o, mdataStandardTypes))))
	}

	expected = append(ToSlice(e), ToSlice(s)...)
	r = s.Concatl(e, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(e, mdataStandardTypes))))
	}
}

func TestSingleConcatr(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1, mdataStandardTypes)
	t := s.Pushl(2, mdataStandardTypes)
	o := (makeEmpty()).Pushl(3, mdataStandardTypes)

	expected := append(ToSlice(s), ToSlice(t)...)
	r := s.Concatr(t, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(t, mdataStandardTypes))))
	}

	expected = append(ToSlice(s), ToSlice(o)...)
	r = s.Concatr(o, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(o, mdataStandardTypes))))
	}

	expected = append(ToSlice(s), ToSlice(e)...)
	r = s.Concatr(e, mdataStandardTypes)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(e, mdataStandardTypes))))
	}
}

func TestSingleFTSize(test *testing.T) {
	s := makeSingle(2, mdataStandardTypes)
	if s.mdataForKey(ft_size_key, mdataStandardTypes).(int) != 1 {
		test.Error(fmt.Sprintf("Expected s.ft_size to equal 1, got %v", s.mdataForKey(ft_size_key, mdataStandardTypes).(int)))
	}
}

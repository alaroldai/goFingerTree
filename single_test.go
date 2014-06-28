package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSinglePushl(test *testing.T) {
	n := makeSingle(1)
	r := n.Pushl(2)
	if cmpslices(ToSlice(r), []Any{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushl(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestSinglePopl(test *testing.T) {
	n := makeSingle(1)
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
	n := makeSingle(1)
	r := n.Pushr(2)
	if cmpslices(ToSlice(r), []Any{1, 2}) == false {
		test.Error(fmt.Sprintf("Expected n.Pushr(2) to result in sequence [1 2], got %v", ToSlice(r)))
	}
}

func TestSingleImplementsFingerTree(test *testing.T) {
	stype := reflect.TypeOf(&single{})
	itype := reflect.TypeOf((*FingerTree)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestSingleFold(test *testing.T) {
	n := makeSingle(1)
	Fold_Test(test, n, Slice{1})
}

func TestSingleIter(test *testing.T) {
	n := makeSingle(1)
	Iter_Test(test, n, Slice{1})
}

func TestSingleHeadr(test *testing.T) {
	v := (makeSingle(1)).Headr()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headr() should be 1, got %v", v))
	}
}

func TestSingleTailr(test *testing.T) {
	v := (makeSingle(1)).Tailr()
	if !v.IsEmpty() {
		test.Error(fmt.Sprintf("single{1}.Tailr() should be empty, got %v", v))
	}
}

func TestSingleHeadl(test *testing.T) {
	v := (makeSingle(1)).Headl()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.Headl() should be 1, got %v", v))
	}
}

func TestSingleTaill(test *testing.T) {
	v := (makeSingle(1)).Taill()
	if !v.IsEmpty() {
		test.Error(fmt.Sprintf("single{1}.Taill() should be empty, got %v", v))
	}
}

func TestSingleIsEmpty(test *testing.T) {
	v := makeSingle(1)
	if v.IsEmpty() {
		test.Error("Expected makeSingle(1).IsEmpty() to be false")
	}
}

func TestSingleConcatl(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1)
	t := s.Pushl(2)
	o := (makeEmpty()).Pushl(3)

	expected := append(ToSlice(t), ToSlice(s)...)
	r := s.Concatl(t)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(t))))
	}

	expected = append(ToSlice(o), ToSlice(s)...)
	r = s.Concatl(o)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(o))))
	}

	expected = append(ToSlice(e), ToSlice(s)...)
	r = s.Concatl(e)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatl to return %v, got %v", expected, ToSlice(s.Concatl(e))))
	}
}

func TestSingleConcatr(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1)
	t := s.Pushl(2)
	o := (makeEmpty()).Pushl(3)

	expected := append(ToSlice(s), ToSlice(t)...)
	r := s.Concatr(t)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(t))))
	}

	expected = append(ToSlice(s), ToSlice(o)...)
	r = s.Concatr(o)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(o))))
	}

	expected = append(ToSlice(s), ToSlice(e)...)
	r = s.Concatr(e)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.Concatr to return %v, got %v", expected, ToSlice(s.Concatr(e))))
	}
}

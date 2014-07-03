package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSinglepushl(test *testing.T) {
	n := makeSingle(1)
	r := n.pushl(2)
	if cmpslices(ToSlice(r), []Any{2, 1}) == false {
		test.Error(fmt.Sprintf("Expected n.pushl(2) to result in sequence [2 1], got %v", ToSlice(r)))
	}
}

func TestSinglepopl(test *testing.T) {
	n := makeSingle(1)
	r, e := n.popl()
	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error("Expected n.popl() result to be an empty node")
	}
	if e != 1 {
		test.Error("Expected n.popl() result to be 1")
	}
}

func TestSinglepushr(test *testing.T) {
	n := makeSingle(1)
	r := n.pushr(2)
	if cmpslices(ToSlice(r), []Any{1, 2}) == false {
		test.Error(fmt.Sprintf("Expected n.pushr(2) to result in sequence [1 2], got %v", ToSlice(r)))
	}
}

func TestSingleImplementsFingerTreeComponent(test *testing.T) {
	stype := reflect.TypeOf(&single{})
	itype := reflect.TypeOf((*FingerTreeComponent)(nil)).Elem()
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

func TestSingleMeasure(test *testing.T) {
	n := makeSingle(mfree{1})
	r := n.Measure()
	Slice_TestM(test, r, Slice{1}, "single{mfree{1}}.Measure()")
}

func TestSingleheadr(test *testing.T) {
	v := (makeSingle(1)).headr()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.headr() should be 1, got %v", v))
	}
}

func TestSingletailr(test *testing.T) {
	v := (makeSingle(1)).tailr()
	if !v.isEmpty() {
		test.Error(fmt.Sprintf("single{1}.tailr() should be empty, got %v", v))
	}
}

func TestSingleheadl(test *testing.T) {
	v := (makeSingle(1)).headl()
	if v != 1 {
		test.Error(fmt.Sprintf("single{1}.headl() should be 1, got %v", v))
	}
}

func TestSingletaill(test *testing.T) {
	v := (makeSingle(1)).taill()
	if !v.isEmpty() {
		test.Error(fmt.Sprintf("single{1}.taill() should be empty, got %v", v))
	}
}

func TestSingleisEmpty(test *testing.T) {
	v := makeSingle(1)
	if v.isEmpty() {
		test.Error("Expected makeSingle(1).isEmpty() to be false")
	}
}

func TestSingleconcatl(test *testing.T) {
	e := makeEmpty()
	s := e.pushl(1)
	t := s.pushl(2)
	o := (makeEmpty()).pushl(3)

	expected := append(ToSlice(t), ToSlice(s)...)
	r := s.concatl(t)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatl to return %v, got %v", expected, ToSlice(s.concatl(t))))
	}

	expected = append(ToSlice(o), ToSlice(s)...)
	r = s.concatl(o)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatl to return %v, got %v", expected, ToSlice(s.concatl(o))))
	}

	expected = append(ToSlice(e), ToSlice(s)...)
	r = s.concatl(e)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatl to return %v, got %v", expected, ToSlice(s.concatl(e))))
	}
}

func TestSingleconcatr(test *testing.T) {
	e := makeEmpty()
	s := e.pushl(1)
	t := s.pushl(2)
	o := (makeEmpty()).pushl(3)

	expected := append(ToSlice(s), ToSlice(t)...)
	r := s.concatr(t)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatr to return %v, got %v", expected, ToSlice(s.concatr(t))))
	}

	expected = append(ToSlice(s), ToSlice(o)...)
	r = s.concatr(o)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatr to return %v, got %v", expected, ToSlice(s.concatr(o))))
	}

	expected = append(ToSlice(s), ToSlice(e)...)
	r = s.concatr(e)
	if !cmpslices(expected, ToSlice(r)) {
		test.Error(fmt.Sprintf("Expected s.concatr to return %v, got %v", expected, ToSlice(s.concatr(e))))
	}
}

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

func TestEmptyMeasure(test *testing.T) {
	n := makeEmpty()
	r := n.Measure()
	if r != Zero {
		test.Error(fmt.Sprintf("empty{}.Measure() should be Zero, got %v", r))
	}
}

func TestEmptyFold(test *testing.T) {
	n := makeEmpty()
	Fold_Test(test, n, Slice{})
}

func TestEmptyIter(test *testing.T) {
	n := makeEmpty()
	Iter_Test(test, n, Slice{})
}

func TestEmptypushl(test *testing.T) {
	v := (makeEmpty()).pushl(1)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.pushl(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptypopl(test *testing.T) {
	n := makeEmpty()
	r, e := n.popl()

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.popl() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.popl() result to be nil")
	}
}

func TestEmptypopr(test *testing.T) {
	n := makeEmpty()
	r, e := n.popr()

	_, isEmpty := r.(*empty)
	if !isEmpty {
		test.Error(fmt.Sprintf("Expected n.popr() result to be an empty node, got %v", r))
	}
	if e != nil {
		test.Error("Expected n.popr() result to be nil")
	}
}

func TestEmptypushr(test *testing.T) {
	v := (makeEmpty()).pushr(1)
	if cmpslices(ToSlice(v), []Any{1}) == false {
		test.Error(fmt.Sprintf("Expected empty{}.pushr(1) to result in single{1}, got %v", ToSlice(v)))
	}
}

func TestEmptyheadr(test *testing.T) {
	v := (makeEmpty()).headr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.headl() should be nil, got %v", v))
	}
}

func TestEmptytailr(test *testing.T) {
	v := (makeEmpty()).tailr()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.tailr() should be nil, got %v", v))
	}
}

func TestEmptyheadl(test *testing.T) {
	v := (makeEmpty()).headl()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.headl() should be nil, got %v", v))
	}
}

func TestEmptytaill(test *testing.T) {
	v := (makeEmpty()).taill()
	if v != nil {
		test.Error(fmt.Sprintf("empty{}.tailr() should be nil, got %v", v))
	}
}

func TestEmptyisEmpty(test *testing.T) {
	v := makeEmpty()
	if !v.isEmpty() {
		test.Error("Expected makeEmpty().isEmpty() to be true")
	}
}

func TestEmptyconcatl(test *testing.T) {
	e := makeEmpty()
	s := e.pushl(1)
	t := s.pushl(2)

	o := makeEmpty()

	if !cmpslices(ToSlice(t), ToSlice(e.concatl(t))) {
		test.Error(fmt.Sprintf("Expected e.concatl to return %v, got %v", ToSlice(t), ToSlice(e.concatl(t))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.concatl(s))) {
		test.Error(fmt.Sprintf("Expected e.concatl to return %v, got %v", ToSlice(s), ToSlice(e.concatl(s))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.concatl(o))) {
		test.Error(fmt.Sprintf("Expected e.concatl to return %v, got %v", ToSlice(o), ToSlice(e.concatl(o))))
	}
}

func TestEmptyconcatr(test *testing.T) {
	e := makeEmpty()
	s := e.pushl(1)
	t := s.pushl(2)

	o := e

	if !cmpslices(ToSlice(t), ToSlice(e.concatr(t))) {
		test.Error(fmt.Sprintf("Expected e.concatr to return %v, got %v", ToSlice(t), ToSlice(e.concatr(t))))
	}

	if !cmpslices(ToSlice(s), ToSlice(e.concatr(s))) {
		test.Error(fmt.Sprintf("Expected e.concatr to return %v, got %v", ToSlice(s), ToSlice(e.concatr(s))))
	}

	if !cmpslices(ToSlice(o), ToSlice(e.concatr(o))) {
		test.Error(fmt.Sprintf("Expected e.concatr to return %v, got %v", ToSlice(o), ToSlice(e.concatr(o))))
	}
}

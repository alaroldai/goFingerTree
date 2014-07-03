package fingerTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFTreeTripleImplementsFingerTreeComponent(test *testing.T) {
	stype := reflect.TypeOf(&ftreeTriple{})
	itype := reflect.TypeOf((*FingerTreeComponent)(nil)).Elem()
	TypeConformityTest(test, stype, itype)
}

func TestFTreeTripleMeasure(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	var s Slice = Slice{}
	for i := 0; i < 20; i++ {
		n = n.pushr(mfree{i})
		s = s.pushr(i)
	}
	Slice_TestM(test, n.Measure(), s, fmt.Sprintf("ftreeTriple{mfree%v}.Measure()", s))
}

func TestFTreeTripleFold(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	var s Slice = Slice{}
	for i := 0; i < 20; i++ {
		n = n.pushr(i)
		s = s.pushr(i)
	}
	Fold_Test(test, n, s)
}

func TestFTreeTripleIter(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	var s Slice = Slice{}
	for i := 0; i < 20; i++ {
		n = n.pushr(i)
		s = s.pushr(i)
	}
	Iter_Test(test, n, s)
}

func TestFTreeTriplepushl(test *testing.T) {
	var n FingerTreeComponent = makeSingle(0)

	for i := 1; i < 20; i++ {
		n = n.pushl(i)
	}

	if cmpslices(ToSlice(n), Slice{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0], got %v", ToSlice(n)))
	}
}

func TestFTreeTriplepopl(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()

	for i := 0; i < 20; i++ {
		n = n.pushl(i)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.pushr(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.popl()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreeTriplepopr(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()

	for i := 0; i < 20; i++ {
		n = n.pushr(i)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.popr()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.pushl(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.popr()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreeTriplepushr(test *testing.T) {
	var n FingerTreeComponent = makeSingle(0)

	for i := 1; i < 20; i++ {
		n = n.pushr(i)
	}

	expected := []Any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	if cmpslices(ToSlice(n), expected) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence %v, got %v", expected, ToSlice(n)))
	}
}

func TestFTreeTripleheadr(test *testing.T) {
	v := (makeEmpty()).pushr(1).pushr(2)
	r := v.headr()
	if r != 2 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.headr() should be 2, got %v", r))
	}

	v = (makeEmpty()).pushl(1).pushl(2)
	r = v.headr()
	if r != 1 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.headr() should be 2, got %v", r))
	}
}

func TestFTreeTripletailr(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTreeComponent(xs).tailr()
	ys := ToSlice(t)
	expected := xs[:len(xs)-1]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftreeTriple{%v}.tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeTripleheadl(test *testing.T) {
	v := makeEmpty().pushl(1).pushr(2).headl()
	if v != 1 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.headr() should be 1, got %v", v))
	}
}

func TestFTreeTripletaill(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTreeComponent(xs).taill()
	ys := ToSlice(t)
	expected := xs[1:]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftreeTriple{%v}.tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeTripleisEmpty(test *testing.T) {
	v := ToFingerTreeComponent(Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if v.isEmpty() {
		test.Error("Expected isEmpty to be false")
	}
}

func TestFTreeTripleconcatl(test *testing.T) {
	e := makeEmpty()
	s := e.pushl(1)
	var t FingerTreeComponent = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.pushl(i)
	}

	var o FingerTreeComponent = makeEmpty()
	for i := 0; i < 5; i++ {
		o = o.pushl(i)
	}

	testCombinations := func() {

		expected := append(ToSlice(o), ToSlice(t)...)
		r := t.concatl(o)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(s), ToSlice(t)...)
		r = t.concatl(s)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(e), ToSlice(t)...)
		r = t.concatl(e)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatl to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 5; i++ {
		t = t.pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.pushl(i)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.pushl(i)
	}
	for i := 0; i < 105; i++ {
		o = o.pushl(i)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 105; i++ {
		t = t.pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.pushl(i)
	}
	testCombinations()
}

func TestFTreeTripleconcatr(test *testing.T) {
	e := makeEmpty()
	s := e.pushr(1)
	var t FingerTreeComponent = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.pushr(i)
	}

	var o FingerTreeComponent = makeEmpty()
	for i := 0; i < 5; i++ {
		o = o.pushr(i)
	}

	testCombinations := func() {

		expected := append(ToSlice(t), ToSlice(o)...)
		r := t.concatr(o)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(s)...)
		r = t.concatr(s)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(e)...)
		r = t.concatr(e)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.concatr to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 5; i++ {
		t = t.pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.pushr(i)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.pushr(i)
	}
	for i := 0; i < 105; i++ {
		o = o.pushr(i)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 105; i++ {
		t = t.pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.pushr(i)
	}
	testCombinations()
}

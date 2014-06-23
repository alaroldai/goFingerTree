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
	for i := 0; i < 20; i++ {
		n = n.Pushr(i)
	}

	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldl(add, Slice{})
	expected := Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	if !cmpslices(r.(Slice), expected) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", expected, r))
	}
}

func TestFTreeFoldr(test *testing.T) {
	var n FingerTree = &empty{}
	for i := 0; i < 20; i++ {
		n = n.Pushl(i)
	}

	add := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	r := n.Foldr(add, Slice{})
	expect := Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	if !cmpslices(r.(Slice), expect) {
		test.Error(fmt.Sprintf("Expected n.Foldl to return %v, got %v", expect, r))
	}
}

func TestFTreePushl(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 20; i++ {
		n = n.Pushl(i)
	}

	if cmpslices(ToSlice(n), Slice{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0], got %v", ToSlice(n)))
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

func TestFTreePopr(test *testing.T) {
	var n FingerTree = &empty{}

	for i := 0; i < 20; i++ {
		n = n.Pushr(i)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.Popr()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushl(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popr()
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreePushr(test *testing.T) {
	var n FingerTree = &single{0}

	for i := 1; i < 20; i++ {
		n = n.Pushr(i)
	}

	expected := []Any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	if cmpslices(ToSlice(n), expected) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence %v, got %v", expected, ToSlice(n)))
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
	v := (&empty{}).Pushr(1).Pushr(2)
	r := v.Headr()
	if r != 2 {
		test.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", r))
	}

	v = (&empty{}).Pushl(1).Pushl(2)
	r = v.Headr()
	if r != 1 {
		test.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", r))
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

func TestFTreeConcatl(test *testing.T) {
	e := &empty{}
	s := e.Pushl(1)
	var t FingerTree = &empty{}
	for i := 0; i < 25; i++ {
		t = t.Pushl(i)
	}

	var o FingerTree = &empty{}
	for i := 0; i < 5; i++ {
		o = o.Pushl(i)
	}

	testCombinations := func() {

		expected := append(ToSlice(o), ToSlice(t)...)
		r := t.Concatl(o)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(s), ToSlice(t)...)
		r = t.Concatl(s)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(e), ToSlice(t)...)
		r = t.Concatl(e)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 5; i++ {
		t = t.Pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i)
	}
	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 25; i++ {
		t = t.Pushl(i)
	}
	for i := 0; i < 105; i++ {
		o = o.Pushl(i)
	}
	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 105; i++ {
		t = t.Pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i)
	}
	testCombinations()
}

func TestFTreeConcatr(test *testing.T) {
	e := &empty{}
	s := e.Pushr(1)
	var t FingerTree = &empty{}
	for i := 0; i < 25; i++ {
		t = t.Pushr(i)
	}

	var o FingerTree = &empty{}
	for i := 0; i < 5; i++ {
		o = o.Pushr(i)
	}

	testCombinations := func() {

		expected := append(ToSlice(t), ToSlice(o)...)
		r := t.Concatr(o)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(s)...)
		r = t.Concatr(s)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(e)...)
		r = t.Concatr(e)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 5; i++ {
		t = t.Pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i)
	}
	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 25; i++ {
		t = t.Pushr(i)
	}
	for i := 0; i < 105; i++ {
		o = o.Pushr(i)
	}
	testCombinations()

	t = &empty{}
	o = &empty{}
	for i := 0; i < 105; i++ {
		t = t.Pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i)
	}
	testCombinations()
}

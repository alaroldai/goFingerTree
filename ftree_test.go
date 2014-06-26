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

func TestFTreeTripleFoldl(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	for i := 0; i < 20; i++ {
		n = n.Pushr(i, mdataStandardTypes)
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

func TestFTreeTripleFoldr(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	for i := 0; i < 20; i++ {
		n = n.Pushl(i, mdataStandardTypes)
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

func TestFTreeTriplePushl(test *testing.T) {
	var n FingerTreeComponent = makeSingle(0, mdataStandardTypes)

	for i := 1; i < 20; i++ {
		n = n.Pushl(i, mdataStandardTypes)
	}

	if cmpslices(ToSlice(n), Slice{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence [19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0], got %v", ToSlice(n)))
	}
}

func TestFTreeTriplePopl(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()

	for i := 0; i < 20; i++ {
		n = n.Pushl(i, mdataStandardTypes)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.Popl(mdataStandardTypes)
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushr(i, mdataStandardTypes)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popl(mdataStandardTypes)
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreeTriplePopr(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()

	for i := 0; i < 20; i++ {
		n = n.Pushr(i, mdataStandardTypes)
	}

	var e Any
	for i := 19; i >= 0; i-- {
		n, e = n.Popr(mdataStandardTypes)
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushl(i, mdataStandardTypes)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popr(mdataStandardTypes)
		if e != i {
			test.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreeTriplePushr(test *testing.T) {
	var n FingerTreeComponent = makeSingle(0, mdataStandardTypes)

	for i := 1; i < 20; i++ {
		n = n.Pushr(i, mdataStandardTypes)
	}

	expected := []Any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	if cmpslices(ToSlice(n), expected) == false {
		test.Error(fmt.Sprintf("Expected push sequence to result in sequence %v, got %v", expected, ToSlice(n)))
	}
}

func TestFTreeTripleIterl(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	for i := 0; i < 10; i++ {
		n = n.Pushl(i, mdataStandardTypes)
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

func TestFTreeTripleIterr(test *testing.T) {
	var n FingerTreeComponent = makeEmpty()
	for i := 0; i < 10; i++ {
		n = n.Pushl(i, mdataStandardTypes)
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

func TestFTreeTripleHeadr(test *testing.T) {
	v := (makeEmpty()).Pushr(1, mdataStandardTypes).Pushr(2, mdataStandardTypes)
	r := v.Headr()
	if r != 2 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.Headr() should be 2, got %v", r))
	}

	v = (makeEmpty()).Pushl(1, mdataStandardTypes).Pushl(2, mdataStandardTypes)
	r = v.Headr()
	if r != 1 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.Headr() should be 2, got %v", r))
	}
}

func TestFTreeTripleTailr(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTreeComponent(xs, mdataStandardTypes).Tailr(mdataStandardTypes)
	ys := ToSlice(t)
	expected := xs[:len(xs)-1]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftreeTriple{%v}.Tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeTripleHeadl(test *testing.T) {
	v := makeEmpty().Pushl(1, mdataStandardTypes).Pushr(2, mdataStandardTypes).Headl()
	if v != 1 {
		test.Error(fmt.Sprintf("ftreeTriple{1 2}.Headr() should be 1, got %v", v))
	}
}

func TestFTreeTripleTaill(test *testing.T) {
	xs := Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := ToFingerTreeComponent(xs, mdataStandardTypes).Taill(mdataStandardTypes)
	ys := ToSlice(t)
	expected := xs[1:]
	if !SliceEqual(expected, ys) {
		test.Error(fmt.Sprintf("ftreeTriple{%v}.Tailr() should be %v, got %v", xs, expected, ys))
	}
}

func TestFTreeTripleIsEmpty(test *testing.T) {
	v := ToFingerTreeComponent(Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, mdataStandardTypes)
	if v.IsEmpty() {
		test.Error("Expected isEmpty to be false")
	}
}

func TestFTreeTripleConcatl(test *testing.T) {
	e := makeEmpty()
	s := e.Pushl(1, mdataStandardTypes)
	var t FingerTreeComponent = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.Pushl(i, mdataStandardTypes)
	}

	var o FingerTreeComponent = makeEmpty()
	for i := 0; i < 5; i++ {
		o = o.Pushl(i, mdataStandardTypes)
	}

	testCombinations := func() {

		expected := append(ToSlice(o), ToSlice(t)...)
		r := t.Concatl(o, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(s), ToSlice(t)...)
		r = t.Concatl(s, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(e), ToSlice(t)...)
		r = t.Concatl(e, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatl to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 5; i++ {
		t = t.Pushl(i, mdataStandardTypes)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i, mdataStandardTypes)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.Pushl(i, mdataStandardTypes)
	}
	for i := 0; i < 105; i++ {
		o = o.Pushl(i, mdataStandardTypes)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 105; i++ {
		t = t.Pushl(i, mdataStandardTypes)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i, mdataStandardTypes)
	}
	testCombinations()
}

func TestFTreeTripleConcatr(test *testing.T) {
	e := makeEmpty()
	s := e.Pushr(1, mdataStandardTypes)
	var t FingerTreeComponent = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.Pushr(i, mdataStandardTypes)
	}

	var o FingerTreeComponent = makeEmpty()
	for i := 0; i < 5; i++ {
		o = o.Pushr(i, mdataStandardTypes)
	}

	testCombinations := func() {

		expected := append(ToSlice(t), ToSlice(o)...)
		r := t.Concatr(o, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(s)...)
		r = t.Concatr(s, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}

		expected = append(ToSlice(t), ToSlice(e)...)
		r = t.Concatr(e, mdataStandardTypes)
		if !cmpslices(expected, ToSlice(r)) {
			test.Error(fmt.Sprintf("Expected t.Concatr to return %v, got %v", expected, ToSlice(r)))
		}
	}

	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 5; i++ {
		t = t.Pushr(i, mdataStandardTypes)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i, mdataStandardTypes)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 25; i++ {
		t = t.Pushr(i, mdataStandardTypes)
	}
	for i := 0; i < 105; i++ {
		o = o.Pushr(i, mdataStandardTypes)
	}
	testCombinations()

	t = makeEmpty()
	o = makeEmpty()
	for i := 0; i < 105; i++ {
		t = t.Pushr(i, mdataStandardTypes)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i, mdataStandardTypes)
	}
	testCombinations()
}

func TestFTreeTripleFTSize(test *testing.T) {
	var t FingerTreeComponent = makeEmpty()
	exp := 105
	for i := 0; i < exp; i++ {
		t = t.Pushr(i, mdataStandardTypes)
	}
	if t.mdataForKey(ft_size_key, mdataStandardTypes).(int) != exp {
		test.Error(fmt.Sprintf("Expected t.mdataForKey(ft_size_key).(int) to equal %v, got %v", exp, t.mdataForKey(ft_size_key, mdataStandardTypes).(int)))
	}
}

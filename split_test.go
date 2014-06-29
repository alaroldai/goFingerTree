package fingerTree

import (
	"testing"
	"fmt"
)

// Measure / Monoid that returns the rightmost entry contained
type keyed struct {
	v int
}

func (k keyed) Measure() Monoid {
	return k
}

func (k keyed) Plus(rm Monoid) Monoid {
	if rm == Zero {
		return k
	}
	return rm
}

func TestDigitSplit(test *testing.T) {
	dig := Slice{keyed{1}, keyed{2}, keyed{3}, keyed{4}}
	l, x, r := splitDigit(dig, func(m Monoid) bool { return m.(keyed).v > 1 }, Zero)
	Slice_TestM(test, l, Slice{keyed{1}}, "splitDigit([1 2 3 4], >1).left")
	if x != (keyed{2}) {
		test.Error(fmt.Sprintf("splitDigit([1 2 3 4], >1).x should be 2, got %v", x))
	}
	Slice_TestM(test, r, Slice{keyed{3}, keyed{4}}, "splitDigit([1 2 3 4], >1).right")
}


func TestInternalSplit(test *testing.T) {
	var s = Slice{}
	for i := 0; i < 20; i++ {
		s = s.Pushr(keyed{i})
	}

	tree := ToFingerTree(s)

	for search := 0; search < 20; search++{
		l, x, r := tree.splitTree(func(m Monoid) bool { return m.(keyed).v >= search }, Zero)

		Slice_TestM(test, ToSlice(l), s[:search], fmt.Sprintf("splitTree(%v, >= %v).left", search, s))
		if x != s[search] {
			test.Error(fmt.Sprintf("splitTree(%v, >= %v).x should be 10, got %v", search, s, x))
		}
		Slice_TestM(test, ToSlice(r), s[search + 1:], fmt.Sprintf("splitTree(%v, >= %v).right", search, s))
	}
}

func TestSplit(test *testing.T) {
	var s = Slice{}
	for i := 0; i < 20; i++ {
		s = s.Pushr(keyed{i})
	}

	tree := ToFingerTree(s)

	// Try a split at all locations
	for x := 0; x <= 20; x++ {
		l, r := tree.Split(func(m Monoid) bool { return m.(keyed).v >= x })

		Slice_TestM(test, ToSlice(l), s[:x], fmt.Sprintf("splitTree(%v, >= %v).left", s, x))
		Slice_TestM(test, ToSlice(r), s[x:], fmt.Sprintf("splitTree(%v, >= %v).right", s, x))
	}
}

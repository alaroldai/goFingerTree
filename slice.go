package fingerTree

type Slice []Any

type FoldFunc func(Any, Any) Any
type Foldable interface {
	Iterable
	Foldl(f FoldFunc, initial Any) Any
	Foldr(f FoldFunc, initial Any) Any
}

type IterFunc func(Any)
type Iterable interface {
	Iterl(IterFunc)
	Iterr(IterFunc)
}

type Sliceable interface {
	ToSlice() Slice
}

func (s Slice) Measure() Monoid {
	fold := func(acc Any, item Any) Any {
		return acc.(Monoid).Plus(Measure(item))
	}

	return s.Foldl(fold, Zero).(Monoid)
}

func (s Slice) Foldl(f FoldFunc, init Any) Any {
	var v = init
	for _, x := range s {
		v = f(v, x)
	}
	return v
}

func (s Slice) Foldr(f FoldFunc, init Any) Any {
	var v = init
	for i := range s {
		v = f(v, s[len(s)-1-i])
	}
	return v
}

func (s Slice) Iterl(f IterFunc) {
	for _, x := range s {
		f(x)
	}
}

func (s Slice) Iterr(f IterFunc) {
	for i, _ := range s {
		f(s[len(s)-1-i])
	}
}

func SliceEqual(a, b Slice) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (s Slice) ToSlice() Slice {
	return s
}

func (s Slice) Pushl(item Any) Slice {
	return append(Slice{item}, s...)
}

func (s Slice) Pushr(item Any) Slice {
	return append(s, item)
}

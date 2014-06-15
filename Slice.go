package fingerTree23

type Slice []Any

type Sliceable interface {
	ToSlice() Slice
}

func (n node2) ToSlice() Slice {
	return n.data[:]
}

func (n node3) ToSlice() Slice {
	return n.data[:]
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
		v = f(v, s[len(s)-1 - i])
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

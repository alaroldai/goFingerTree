package fingerTree

type node interface {
	Foldable
	Sliceable
	Measured
}

type node2 struct {
	measure Monoid
	data [2]Any
}

func makeNode2(a, b Any) *node2 {
	mdata := Measure(a).Plus(Measure(b))

	return &node2{mdata, [2]Any{a, b}}
}

func (n node2) Measure() Monoid {
	return n.measure
}

func (n node2) ToSlice() Slice {
	return n.data[:]
}

func (n node2) Foldl(f FoldFunc, initial Any) Any {
	return n.ToSlice().Foldl(f, initial)
}

func (n node2) Foldr(f FoldFunc, initial Any) Any {
	return n.ToSlice().Foldr(f, initial)
}

func (n node2) Iterl(f IterFunc) {
	n.ToSlice().Iterl(f)
}

func (n node2) Iterr(f IterFunc) {
	n.ToSlice().Iterr(f)
}

type node3 struct {
	measure Monoid
	data [3]Any
}

func makeNode3(a, b, c Any) *node3 {
	mdata := Measure(a).Plus(Measure(b)).Plus(Measure(c))

	return &node3{mdata, [3]Any{a, b, c}}
}

func (n node3) Measure() Monoid {
	return n.measure
}

func (n node3) ToSlice() Slice {
	return n.data[:]
}

func (n node3) Foldl(f FoldFunc, initial Any) Any {
	return n.ToSlice().Foldl(f, initial)
}

func (n node3) Foldr(f FoldFunc, initial Any) Any {
	return n.ToSlice().Foldr(f, initial)
}

func (n node3) Iterl(f IterFunc) {
	n.ToSlice().Iterl(f)
}

func (n node3) Iterr(f IterFunc) {
	n.ToSlice().Iterr(f)
}

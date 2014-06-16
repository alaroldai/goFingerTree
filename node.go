package fingerTree

type node interface {
	Foldable
	Sliceable
}

type node2 struct {
	data [2]Any
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
	data [3]Any
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

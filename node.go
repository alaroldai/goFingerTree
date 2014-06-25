package fingerTree

type node interface {
	Foldable
	Sliceable
	mdata
}

type node2 struct {
	size int
	data [2]Any
}

func makeNode2(a, b Any) *node2 {
	var an, bn mdata
	var succ bool
	sz := 2
	an, succ = a.(mdata)
	if succ {
		bn, _ = b.(mdata)
		sz = an.ft_size() + bn.ft_size()
	}
	return &node2{sz, [2]Any{a, b}}
}

func (n node2) ft_size() int {
	return n.size
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
	size int
	data [3]Any
}

func makeNode3(a, b, c Any) *node3 {
	var an, bn, cn mdata
	var succ bool
	sz := 3
	an, succ = a.(mdata)
	if succ {
		bn, _ = b.(mdata)
		cn, _ = c.(mdata)
		sz = an.ft_size() + bn.ft_size() + cn.ft_size()
	}
	return &node3{sz, [3]Any{a, b, c}}
}

func (n node3) ft_size() int {
	return n.size
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

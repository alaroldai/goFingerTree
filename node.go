package fingerTree

type node interface {
	Foldable
	Sliceable
	mdata
}

type node2 struct {
	metadata map[string]Any
	data     [2]Any
}

func makeNode2(a, b Any) *node2 {
	meta := make(map[string]Any)

	var an, bn mdata
	var succ bool

	for k, v := range mdataTypes {
		meta[k] = v.compose(v.unit, v.unit)
	}

	an, succ = a.(mdata)
	if succ {
		bn, _ = b.(mdata)
		for k, v := range mdataTypes {
			meta[k] = v.compose(an.mdataForKey(k), bn.mdataForKey(k))
		}
	}
	return &node2{meta, [2]Any{a, b}}
}

func (n node2) mdataForKey(key string) Any {
	return n.metadata[key]
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
	metadata map[string]Any
	data     [3]Any
}

func makeNode3(a, b, c Any) *node3 {
	meta := make(map[string]Any)

	var an, bn, cn mdata
	var succ bool

	for k, v := range mdataTypes {
		meta[k] = v.compose(v.compose(v.unit, v.unit), v.unit)
	}

	an, succ = a.(mdata)
	if succ {
		bn, _ = b.(mdata)
		cn, _ = c.(mdata)
		for k, v := range mdataTypes {
			meta[k] = v.compose(v.compose(an.mdataForKey(k), bn.mdataForKey(k)), cn.mdataForKey(k))
		}
	}
	return &node3{meta, [3]Any{a, b, c}}
}

func (n node3) mdataForKey(key string) Any {
	return n.metadata[key]
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

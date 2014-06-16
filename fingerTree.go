package fingerTree

type Any interface{}

type FingerTree interface {
	Foldable

	Pushl(d Any) FingerTree
	Pushr(d Any) FingerTree

	Popl() (FingerTree, Any)
	Popr() (FingerTree, Any)

	Headl() Any
	Headr() Any

	Taill() FingerTree
	Tailr() FingerTree

	IsEmpty() bool
}

func ToSlice(t FingerTree) Slice {
	app := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	return t.Foldl(app, make(Slice, 0)).(Slice)
}

func ToFingerTree(f Foldable) FingerTree {
	push := func(tree Any, item Any) Any {
		return tree.(FingerTree).Pushr(item)
	}

	return f.Foldl(push, empty{}).(FingerTree)
}

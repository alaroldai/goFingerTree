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

	/**
	 *	Note = t.Concatl(o) means 'concatenate o to the left of t'
	 *	e.g. (1).Concatl((2)) => (2 1)
	 *
	 *	Similarly Concatr:
	 *	e.g. (1).Concatr((2)) => (1 2)
	 */
	Concatl(other FingerTree) FingerTree
	Concatr(other FingerTree) FingerTree

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

	return f.Foldl(push, &empty{}).(FingerTree)
}

/**
 *	Utility functions
 */

/**
 *	Transform a slice of elements into a slice of nodes
 */
func nodes(xs Slice) Slice {
	if len(xs) == 1 {
		panic("Can't make a node from a single element.")
	}
	if len(xs) == 2 {
		return Slice{&node2{[2]Any{xs[0], xs[1]}}}
	}
	if len(xs) == 3 {
		return Slice{&node3{[3]Any{xs[0], xs[1], xs[2]}}}
	}
	if len(xs) == 4 {
		return Slice{&node2{[2]Any{xs[0], xs[1]}}, &node2{[2]Any{xs[2], xs[3]}}}
	}
	if len(xs) > 4 {
		return append(nodes(xs[:3]), nodes(xs[3:])...)
	}
	return Slice{}
}

/**
 *	Join two finger trees with a 'glue' slice between them
 *	Normally calling Concatr or Concatl will be more useful
 */
func glue(l FingerTree, c Slice, r FingerTree) FingerTree {

	pushl := func(a FingerTree, s Slice) FingerTree {
		m := a
		s.Iterr(func(t Any) {
			m = m.Pushl(t)
		})
		return m
	}

	pushr := func(a FingerTree, s Slice) FingerTree {
		m := a
		s.Iterl(func(t Any) {
			m = m.Pushr(t)
		})
		return m
	}

	// If either branch is empty, it can be ignored
	if l.IsEmpty() {
		return pushl(r, c)
	}
	if r.IsEmpty() {
		return pushr(l, c)
	}

	// If either branch is a single, glue reduces to pushl/pushr
	s, succ := l.(*single)
	if succ {
		return pushl(r, c).Pushl(s.data)
	}
	s, succ = r.(*single)
	if succ {
		return pushr(l, c).Pushr(s.data)
	}

	// Otherwise, both branches are trees. We proceed recursively:
	lt, _ := l.(*ftree)
	rt, _ := r.(*ftree)

	ns := nodes(append(append(lt.right, c...), rt.left...))
	nc := glue(lt.child, ns, rt.child)
	return &ftree{
		lt.left,
		nc,
		rt.right,
	}
}

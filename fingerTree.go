package fingerTree

type Any interface{}

type FingerTreeComponent interface {
	Foldable
	Measured

	pushl(d Any) FingerTreeComponent
	pushr(d Any) FingerTreeComponent

	popl() (FingerTreeComponent, Any)
	popr() (FingerTreeComponent, Any)

	headl() Any
	headr() Any

	taill() FingerTreeComponent
	tailr() FingerTreeComponent

	/**
	 *	Note = t.concatl(o) means 'concatenate o to the left of t'
	 *	e.g. (1).concatl((2)) => (2 1)
	 *
	 *	Similarly concatr:
	 *	e.g. (1).concatr((2)) => (1 2)
	 */
	concatl(other FingerTreeComponent) FingerTreeComponent
	concatr(other FingerTreeComponent) FingerTreeComponent

	split(pred func(Monoid) bool) (FingerTreeComponent, FingerTreeComponent)

	isEmpty() bool

	// internal: split into (left, x, right)
	// where x is the first item whose totaled Measure satisfies 'pred'
	// requires that:
	//    - !pred(init)
	//    - pred(init + tree.Measure())
	//    - !tree.isEmpty()
	splitTree(pred func(Monoid) bool, init Monoid) (FingerTreeComponent, Any, FingerTreeComponent)
}

func ToSlice(t FingerTreeComponent) Slice {
	app := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	return t.Foldl(app, make(Slice, 0)).(Slice)
}

func ToFingerTreeComponent(f Foldable) FingerTreeComponent {
	push := func(tree Any, item Any) Any {
		return tree.(FingerTreeComponent).pushr(item)
	}

	return f.Foldl(push, makeEmpty()).(FingerTreeComponent)
}

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

	Concatl(other FingerTree) FingerTree
	Concatr(other FingerTree) FingerTree

	Split(pred func(Monoid) bool) (FingerTree, FingerTree)
	IsEmpty() bool
}

type MDType struct {
	zero    Monoid
	combine func(Monoid, Monoid) Monoid
	measure func(Any) Monoid
}

type FTree struct {
	root     FingerTreeComponent
	metaType &MDType
}

func New(metaType &MDType) {
	return &FTree{makeEmpty(), metaType}
}

func (t *FTree) Pushl(d Any) FingerTree {
	return New(t.root.pushl(d), t.metaType)
}

func (t *FTree) Pushr(d Any) FingerTree {
	return New(t.root.pushr(d), t.metaType)
}

func (t *FTree) Popl() (FingerTree, Any) {
	return New(t.root.popl(), t.metaType)
}

func (t *FTree) Popr() (FingerTree, Any) {
	return New(t.root.popr(), t.metaType)
}

func (t *FTree) Headl() Any {
	return t.root.headl()
}

func (t *FTree) Headr() Any {
	return t.root.headr()
}

func (t *FTree) Taill() FingerTree {
	return New(t.root.taill(), t.metaType)
}

func (t *FTree) Tailr() FingerTree {
	return New(t.root.tailr(), t.metaType)
}

func (t *FTree) Concatl(other FingerTree) FingerTree {
	return New(t.root.concatl(other), t.metaType)
}

func (t *FTree) Concatr(other FingerTree) FingerTree {
	return New(t.root.concatr(other), t.metaType)
}

func (t *FTree) Split(pred func(Monoid) bool) (FingerTree, FingerTree) {
	l, r = t.root.split(pred)
	return New(l, t.metaType), New(r, t.metaType)
}

func (t *FTree) IsEmpty() bool {
	return t.root.isEmpty()
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
		return Slice{makeNode2(xs[0], xs[1])}
	}
	if len(xs) == 3 {
		return Slice{makeNode3(xs[0], xs[1], xs[2])}
	}
	if len(xs) == 4 {
		return Slice{makeNode2(xs[0], xs[1]), makeNode2(xs[2], xs[3])}
	}
	if len(xs) > 4 {
		return append(nodes(xs[:3]), nodes(xs[3:])...)
	}
	return Slice{}
}

/**
 *	Join two finger trees with a 'glue' slice between them
 *	Normally calling concatr or concatl will be more useful
 */
func glue(l FingerTreeComponent, c Slice, r FingerTreeComponent) FingerTreeComponent {

	pushl := func(a FingerTreeComponent, s Slice) FingerTreeComponent {
		m := a
		s.Iterr(func(t Any) {
			m = m.pushl(t)
		})
		return m
	}

	pushr := func(a FingerTreeComponent, s Slice) FingerTreeComponent {
		m := a
		s.Iterl(func(t Any) {
			m = m.pushr(t)
		})
		return m
	}

	// If either branch is empty, it can be ignored
	if l.isEmpty() {
		return pushl(r, c)
	}
	if r.isEmpty() {
		return pushr(l, c)
	}

	// If either branch is a single, glue reduces to pushl/pushr
	s, succ := l.(*single)
	if succ {
		return pushl(r, c).pushl(s.data)
	}
	s, succ = r.(*single)
	if succ {
		return pushr(l, c).pushr(s.data)
	}

	// Otherwise, both branches are trees. We proceed recursively:
	lt, _ := l.(*ftreeTriple)
	rt, _ := r.(*ftreeTriple)

	ns := nodes(append(append(lt.right, c...), rt.left...))
	nc := glue(lt.child, ns, rt.child)
	return makeFTreeTriple(
		lt.left,
		nc,
		rt.right,
	)
}

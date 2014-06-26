package fingerTree

type Any interface{}

type FingerTreeComponent interface {
	Sliceable
	mdataContainer

	Pushl(d Any, mdtypes mdataTypeMap) FingerTreeComponent
	Pushr(d Any, mdtypes mdataTypeMap) FingerTreeComponent

	Popl(mdtypes mdataTypeMap) (FingerTreeComponent, Any)
	Popr(mdtypes mdataTypeMap) (FingerTreeComponent, Any)

	Headl() Any
	Headr() Any

	Taill(mdtypes mdataTypeMap) FingerTreeComponent
	Tailr(mdtypes mdataTypeMap) FingerTreeComponent

	/**
	 *	Note = t.Concatl(o) means 'concatenate o to the left of t'
	 *	e.g. (1).Concatl((2)) => (2 1)
	 *
	 *	Similarly Concatr:
	 *	e.g. (1).Concatr((2)) => (1 2)
	 *
	 *	For an infix-like operation equivalent to "a (+) b", use a.Concatr(b)
	 */
	Concatl(other FingerTreeComponent, mdtypes mdataTypeMap) FingerTreeComponent
	Concatr(other FingerTreeComponent, mdtypes mdataTypeMap) FingerTreeComponent

	/**
	 *	Splits a finger tree before the value specified by mdVal.
	 *	e.g., (splitting by index) [1, 2, 3].Split(0) -> ([], [1, 2, 3])
	 *							   [1, 2, 3].Split(3) -> ([1, 2, 3], [])
	 *							   [1, 2, 3].Split(1) -> ([1], [2, 3])
	 */
	// Split(mdKey string, mdVal Any) (FingerTree, FingerTree)

	IsEmpty() bool
}

func ToSlice(t FingerTreeComponent) Slice {
	app := func(a Any, b Any) Any {
		return append(a.(Slice), b)
	}
	return t.Foldl(app, make(Slice, 0)).(Slice)
}

func ToFingerTreeComponent(f Foldable, mdataTypes mdataTypeMap) FingerTreeComponent {
	push := func(tree Any, item Any) Any {
		return tree.(FingerTreeComponent).Pushr(item, mdataTypes)
	}

	return f.Foldl(push, makeEmpty()).(FingerTreeComponent)
}

/**
 *	Utility functions
 */

/**
 *	Transform a slice of elements into a slice of nodes
 */
func nodes(xs Slice, mdataTypes mdataTypeMap) Slice {
	if len(xs) == 1 {
		panic("Can't make a node from a single element.")
	}
	if len(xs) == 2 {
		return Slice{makeNode2(xs[0], xs[1], mdataTypes)}
	}
	if len(xs) == 3 {
		return Slice{makeNode3(xs[0], xs[1], xs[2], mdataTypes)}
	}
	if len(xs) == 4 {
		return Slice{makeNode2(xs[0], xs[1], mdataTypes), makeNode2(xs[2], xs[3], mdataTypes)}
	}
	if len(xs) > 4 {
		return append(nodes(xs[:3], mdataTypes), nodes(xs[3:], mdataTypes)...)
	}
	return Slice{}
}

/**
 *	Join two finger trees with a 'glue' slice between them
 *	Normally calling Concatr or Concatl will be more useful
 */
func glue(l FingerTreeComponent, c Slice, r FingerTreeComponent, mdataTypes mdataTypeMap) FingerTreeComponent {

	pushl := func(a FingerTreeComponent, s Slice) FingerTreeComponent {
		m := a
		s.Iterr(func(t Any) {
			m = m.Pushl(t, mdataTypes)
		})
		return m
	}

	pushr := func(a FingerTreeComponent, s Slice) FingerTreeComponent {
		m := a
		s.Iterl(func(t Any) {
			m = m.Pushr(t, mdataTypes)
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
		return pushl(r, c).Pushl(s.data, mdataTypes)
	}
	s, succ = r.(*single)
	if succ {
		return pushr(l, c).Pushr(s.data, mdataTypes)
	}

	// Otherwise, both branches are trees. We proceed recursively:
	lt, _ := l.(*ftreeTriple)
	rt, _ := r.(*ftreeTriple)

	ns := nodes(append(append(lt.right, c...), rt.left...), mdataTypes)
	nc := glue(lt.child, ns, rt.child, mdataTypes)
	return makeFTreeTriple(
		lt.left,
		nc,
		rt.right,
		mdataTypes,
	)
}

type FTree struct {
	tree       FingerTreeComponent
	mdataTypes mdataTypeMap
}

func New() *FTree {
	return &FTree{
		makeEmpty(),
		mdataTypeMap{
			ft_size_key: mdataField{
				0,
				1,
				func(a, b Any) Any {
					return a.(int) + b.(int)
				},
			},
		},
	}
}

func makeFTree(tree FingerTreeComponent, mdataTypes mdataTypeMap) *FTree {
	return &FTree{tree, mdataTypes}
}

func (t *FTree) Foldl(f FoldFunc, initial Any) Any {
	return t.tree.Foldl(f, initial)
}
func (t *FTree) Foldr(f FoldFunc, initial Any) Any {
	return t.tree.Foldr(f, initial)
}
func (t *FTree) Iterl(f IterFunc) {
	t.tree.Iterl(f)
}
func (t *FTree) Iterr(f IterFunc) {
	t.tree.Iterr(f)
}
func (t *FTree) ToSlice() Slice {
	return t.tree.ToSlice()
}

func (t *FTree) Pushl(d Any) *FTree {
	return makeFTree(t.tree.Pushl(d, t.mdataTypes), t.mdataTypes)
}
func (t *FTree) Pushr(d Any) *FTree {
	return makeFTree(t.tree.Pushr(d, t.mdataTypes), t.mdataTypes)
}

func (t *FTree) Popl() (*FTree, Any) {
	tree, value := t.tree.Popl(t.mdataTypes)
	return makeFTree(tree, t.mdataTypes), value
}
func (t *FTree) Popr() (*FTree, Any) {
	tree, value := t.tree.Popr(t.mdataTypes)
	return makeFTree(tree, t.mdataTypes), value
}

func (t *FTree) Headl() Any {
	return t.tree.Headl()
}
func (t *FTree) Headr() Any {
	return t.tree.Headr()
}

func (t *FTree) Taill() *FTree {
	return makeFTree(t.tree.Taill(t.mdataTypes), t.mdataTypes)
}
func (t *FTree) Tailr() *FTree {
	return makeFTree(t.tree.Tailr(t.mdataTypes), t.mdataTypes)
}

func (t *FTree) Concatl(other *FTree) *FTree {
	return makeFTree(t.tree.Concatl(other.tree, t.mdataTypes), t.mdataTypes)
}
func (t *FTree) Concatr(other *FTree) *FTree {
	return makeFTree(t.tree.Concatr(other.tree, t.mdataTypes), t.mdataTypes)
}

// Split(mdKey string, mdVal Any) (*FTree, *FTree)

func (t *FTree) IsEmpty() bool {
	return t.tree.IsEmpty()
}

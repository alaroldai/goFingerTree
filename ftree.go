package fingerTree

/**
 *	ftree structure
 */

type ftree struct {
	size int
	left  Slice
	child FingerTree
	right Slice
}

func makeFTree(left Slice, child FingerTree, right Slice) *ftree {
	sz := (Slice{left, right}).Foldl(func (init Any, curr Any) Any {
		return curr.(Slice).Foldl(func (i Any, a Any) Any {
			an, succ := a.(mdata)
			if succ {
				return i.(int) + an.ft_size()
			}
			return i.(int) + 1
		}, init)
	}, 0).(int) + child.ft_size()
	return &ftree{
		sz,
		left,
		child,
		right,
	}
}

func (t *ftree) ft_size() int {
	return t.size
}

func (t *ftree) Foldl(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldl(f, init)
	}

	var a Any = t.left.Foldl(f, initial)
	var b Any = t.child.Foldl(lift, a)
	return t.right.Foldl(f, b)
}

func (t *ftree) Foldr(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldr(f, init)
	}

	a := t.right.Foldr(f, initial)
	b := t.child.Foldr(lift, a)
	return t.left.Foldr(f, b)
}

func (t *ftree) Pushl(d Any) FingerTree {
	if len(t.left) < 4 {
		return makeFTree(
			append([]Any{d}, t.left...),
			t.child,
			t.right,
		)
	}

	var child FingerTree
	pushdown := makeNode3(
			t.left[1],
			t.left[2],
			t.left[3],
	)

	child = t.child.Pushl(pushdown)

	return makeFTree(
		Slice{d, t.left[0]},
		child,
		t.right,
	)
}

func (t *ftree) Popl() (FingerTree, Any) {
	return t.Taill(), t.Headl()
}

func (t *ftree) Popr() (FingerTree, Any) {
	return t.Tailr(), t.Headr()
}

func (t *ftree) Pushr(d Any) FingerTree {
	if len(t.right) < 4 {
		return makeFTree(
			t.left,
			t.child,
			append(t.right, d),
		)
	}

	var child FingerTree
	pushdown := makeNode3(
			t.right[0],
			t.right[1],
			t.right[2],
	)

	child = t.child.Pushr(pushdown)

	return makeFTree(
		t.left,
		child,
		[]Any{t.right[3], d},
	)
}

func (t *ftree) Iterl(f IterFunc) {
	t.Foldl(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t *ftree) Iterr(f IterFunc) {
	t.Foldr(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t *ftree) Headr() Any {
	return t.right[len(t.right)-1]
}

func (t *ftree) Headl() Any {
	return t.left[0]
}

func buildr(left Slice, m FingerTree, right Slice) FingerTree {
	if len(right) > 0 {
		return makeFTree(
			left,
			m,
			right,
		)
	}

	if m.IsEmpty() {
		return ToFingerTree(left)
	}

	return makeFTree(
		left,
		m.Tailr(),
		m.Headr().(node).ToSlice(),
	)
}

func buildl(left Slice, m FingerTree, right Slice) FingerTree {
	if len(left) > 0 {
		return makeFTree(
			left,
			m,
			right,
		)
	}

	if m.IsEmpty() {
		return ToFingerTree(right)
	}

	return makeFTree(
		m.Headl().(node).ToSlice(),
		m.Taill(),
		right,
	)
}

func (t *ftree) Tailr() FingerTree {
	return buildr(t.left, t.child, t.right[:len(t.right)-1])
}

func (t *ftree) Taill() FingerTree {
	return buildl(t.left[1:], t.child, t.right)
}

func (t *ftree) IsEmpty() bool {
	return false
}

func (t *ftree) Concatr(other FingerTree) FingerTree {
	otherAsFtree, isFTree := other.(*ftree)
	if !isFTree {
		return other.Concatl(t)
	}

	return glue(t, Slice{}, otherAsFtree)
}

func (t *ftree) Concatl(other FingerTree) FingerTree {
	return other.Concatr(t)
}

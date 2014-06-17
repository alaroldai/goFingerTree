package fingerTree

type ftree struct {
	left  Slice
	right Slice
	child FingerTree
}

func (t ftree) Foldl(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldl(f, init)
	}

	var a Any = t.left.Foldl(f, initial)
	var b Any = t.child.Foldl(lift, a)
	return t.right.Foldl(f, b)
}

func (t ftree) Foldr(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldr(f, init)
	}

	a := t.right.Foldr(f, initial)
	b := t.child.Foldr(lift, a)
	return t.left.Foldr(f, b)
}

func (t ftree) Pushl(d Any) FingerTree {
	if len(t.left) < 4 {
		return &ftree{
			append([]Any{d}, t.left...),
			t.right,
			t.child,
		}
	}

	var child FingerTree
	pushdown := &node3{
		[3]Any{
			t.left[1],
			t.left[2],
			t.left[3],
		},
	}

	child = t.child.Pushl(pushdown)

	return &ftree{
		Slice{d, t.left[0]},
		t.right,
		child,
	}
}

func (t ftree) Popl() (FingerTree, Any) {
	return t.Taill(), t.Headl()
}

func (t ftree) Popr() (FingerTree, Any) {
	return t.Tailr(), t.Headr()
}

func (t ftree) Pushr(d Any) FingerTree {
	if len(t.right) < 4 {
		return &ftree{
			t.left,
			append(t.right, d),
			t.child,
		}
	}

	var child FingerTree
	pushdown := &node3{
		[3]Any{
			t.right[0],
			t.right[1],
			t.right[2],
		},
	}

	child = t.child.Pushr(pushdown)

	return &ftree{
		t.left,
		[]Any{t.right[3], d},
		child,
	}
}

func (t ftree) Iterl(f IterFunc) {
	t.Foldl(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t ftree) Iterr(f IterFunc) {
	t.Foldr(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t ftree) Headr() Any {
	return t.right[len(t.right)-1]
}

func (t ftree) Headl() Any {
	return t.left[0]
}

func buildr(left Slice, m FingerTree, right Slice) FingerTree {
	if len(right) > 0 {
		return &ftree{left, right, m}
	}

	if m.IsEmpty() {
		return ToFingerTree(left)
	}

	return &ftree{
		left,
		m.Headr().(node).ToSlice(),
		m.Tailr(),
	}
}

func buildl(left Slice, m FingerTree, right Slice) FingerTree {
	if len(left) > 0 {
		return &ftree{left, right, m}
	}

	if m.IsEmpty() {
		return ToFingerTree(right)
	}

	return &ftree{
		m.Headl().(node).ToSlice(),
		right,
		m.Taill(),
	}
}

func (t ftree) Tailr() FingerTree {
	return buildr(t.left, t.child, t.right[:len(t.right)-1])
}

func (t ftree) Taill() FingerTree {
	return buildl(t.left[1:], t.child, t.right)
}

func (t ftree) IsEmpty() bool {
	return false
}

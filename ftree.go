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
	if len(t.left) > 1 {
		return &ftree{
				t.left[1:],
				t.right,
				t.child,
			},
			t.left[0]
	}

	nc, p := t.child.Popl()

	if p == nil {
		lright := len(t.right)
		if lright == 0 {
			return &empty{}, t.left[0]
		}
		if lright == 1 {
			return &single{t.right[0]}, t.left[0]
		}
		if lright == 2 || lright == 3 {
			return &ftree{
					t.right[:1],
					t.right[1:],
					&empty{},
				},
				t.left[0]
		}
		if lright == 4 {
			return &ftree{
					t.right[:2],
					t.right[2:],
					&empty{},
				},
				t.left[0]
		}

		panic("Invalid number of elements in right branch")
	} else {
		return &ftree{
				p.(node).ToSlice(),
				t.right,
				nc,
			},
			t.left[0]
	}
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
		m.Tailr(),
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

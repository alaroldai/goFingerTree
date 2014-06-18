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

func (t ftree) Concatr(other FingerTree) FingerTree {
	otherAsFtreePtr, isFTreePtr := other.(*ftree)
	otherAsFtreeStruct, isFTreeStruct := other.(ftree)
	if !isFTreePtr && !isFTreeStruct {
		return other.Concatl(t)
	}

	concatr := func(o *ftree) FingerTree {

		var nodes func(Slice) Slice
		nodes = func(xs Slice) Slice {
			if len(xs) == 1 {
				return Slice{&node2{[2]Any{xs[0], nil}}}
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

		var app3 func(FingerTree, Slice, FingerTree) FingerTree
		app3 = func(l FingerTree, c Slice, r FingerTree) FingerTree {

			if l.IsEmpty() {
				m := r
				for i, _ := range c {
					m = m.Pushl(c[len(c)-i-1])
				}
				return m
			}
			if r.IsEmpty() {
				m := l
				for _, t := range c {
					m = m.Pushr(t)
				}
				return m
			}

			s, succ := l.(*single)

			if succ {
				m := r
				for i, _ := range c {
					m = m.Pushl(c[len(c)-i-1])
				}
				return m.Pushl(s.data)
			}

			s, succ = r.(*single)

			if succ {
				m := l
				for _, t := range c {
					m = m.Pushr(t)
				}
				return m.Pushr(s.data)
			}

			nl := l.(*ftree).left
			nr := r.(*ftree).right
			ns := nodes(append(append(l.(*ftree).right, c...), r.(*ftree).left...))
			nc := app3(l.(*ftree).child, ns, r.(*ftree).child)
			return &ftree{nl, nr, nc}
		}

		return app3(&t, Slice{}, o)
	}

	if isFTreePtr {
		return concatr(otherAsFtreePtr)
	}
	return concatr(&otherAsFtreeStruct)
}

func (t ftree) Concatl(other FingerTree) FingerTree {
	_, isFTreePtr := other.(*ftree)
	_, isFTreeStruct := other.(ftree)
	if !isFTreePtr && !isFTreeStruct {
		return other.Concatr(t)
	}

	return other.Concatr(&t)
}

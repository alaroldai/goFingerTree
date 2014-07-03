package fingerTree

/**
 *	ftreeTriple structure
 */

type ftreeTriple struct {
	measure Monoid
	left    Slice
	child   FingerTreeComponent
	right   Slice
}

func makeFTreeTriple(left Slice, child FingerTreeComponent, right Slice) *ftreeTriple {
	mdata := left.Measure().Plus(child.Measure()).Plus(right.Measure())

	return &ftreeTriple{
		mdata,
		left,
		child,
		right,
	}
}

func (t *ftreeTriple) Measure() Monoid {
	return t.measure
}

func (t *ftreeTriple) Foldl(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldl(f, init)
	}

	var a Any = t.left.Foldl(f, initial)
	var b Any = t.child.Foldl(lift, a)
	return t.right.Foldl(f, b)
}

func (t *ftreeTriple) Foldr(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldr(f, init)
	}

	a := t.right.Foldr(f, initial)
	b := t.child.Foldr(lift, a)
	return t.left.Foldr(f, b)
}

func (t *ftreeTriple) pushl(d Any) FingerTreeComponent {
	if len(t.left) < 4 {
		return makeFTreeTriple(
			append([]Any{d}, t.left...),
			t.child,
			t.right,
		)
	}

	var child FingerTreeComponent
	pushdown := makeNode3(
		t.left[1],
		t.left[2],
		t.left[3],
	)

	child = t.child.pushl(pushdown)

	return makeFTreeTriple(
		Slice{d, t.left[0]},
		child,
		t.right,
	)
}

func (t *ftreeTriple) popl() (FingerTreeComponent, Any) {
	return t.taill(), t.headl()
}

func (t *ftreeTriple) popr() (FingerTreeComponent, Any) {
	return t.tailr(), t.headr()
}

func (t *ftreeTriple) pushr(d Any) FingerTreeComponent {
	if len(t.right) < 4 {
		return makeFTreeTriple(
			t.left,
			t.child,
			append(t.right, d),
		)
	}

	var child FingerTreeComponent
	pushdown := makeNode3(
		t.right[0],
		t.right[1],
		t.right[2],
	)

	child = t.child.pushr(pushdown)

	return makeFTreeTriple(
		t.left,
		child,
		[]Any{t.right[3], d},
	)
}

func (t *ftreeTriple) Iterl(f IterFunc) {
	t.Foldl(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t *ftreeTriple) Iterr(f IterFunc) {
	t.Foldr(func(_ Any, b Any) Any {
		f(b)
		return nil
	}, nil)
}

func (t *ftreeTriple) headr() Any {
	return t.right[len(t.right)-1]
}

func (t *ftreeTriple) headl() Any {
	return t.left[0]
}

func buildr(left Slice, m FingerTreeComponent, right Slice) FingerTreeComponent {
	if len(right) > 0 {
		return makeFTreeTriple(
			left,
			m,
			right,
		)
	}

	if m.isEmpty() {
		return ToFingerTreeComponent(left)
	}

	return makeFTreeTriple(
		left,
		m.tailr(),
		m.headr().(node).ToSlice(),
	)
}

func buildl(left Slice, m FingerTreeComponent, right Slice) FingerTreeComponent {
	if len(left) > 0 {
		return makeFTreeTriple(
			left,
			m,
			right,
		)
	}

	if m.isEmpty() {
		return ToFingerTreeComponent(right)
	}

	return makeFTreeTriple(
		m.headl().(node).ToSlice(),
		m.taill(),
		right,
	)
}

func (t *ftreeTriple) tailr() FingerTreeComponent {
	return buildr(t.left, t.child, t.right[:len(t.right)-1])
}

func (t *ftreeTriple) taill() FingerTreeComponent {
	return buildl(t.left[1:], t.child, t.right)
}

func (t *ftreeTriple) isEmpty() bool {
	return false
}

func (t *ftreeTriple) concatr(other FingerTreeComponent) FingerTreeComponent {
	otherAsFtree, isFTreeTriple := other.(*ftreeTriple)
	if !isFTreeTriple {
		return other.concatl(t)
	}

	return glue(t, Slice{}, otherAsFtree)
}

func (t *ftreeTriple) concatl(other FingerTreeComponent) FingerTreeComponent {
	return other.concatr(t)
}

func splitDigit(digit Slice, pred func(Monoid) bool, init Monoid) (Slice, Any, Slice) {
	var measure = init
	for i, a := range digit {
		measure = measure.Plus(Measure(a))
		if pred(measure) {
			return digit[:i], digit[i], digit[i+1:]
		}
	}
	return digit[:len(digit)-1], digit[len(digit)-1], Slice{}
}

func (t ftreeTriple) splitTree(pred func(Monoid) bool, init Monoid) (FingerTreeComponent, Any, FingerTreeComponent) {
	vleft := init.Plus(t.left.Measure())
	if pred(vleft) {
		l, x, r := splitDigit(t.left, pred, init)
		return ToFingerTreeComponent(l), x, buildl(r, t.child, t.right)
	}

	vchild := vleft.Plus(t.child.Measure())
	if pred(vchild) {
		ml, mx, mr := t.child.splitTree(pred, vleft)
		l, x, r := splitDigit(mx.(node).ToSlice(), pred, vleft.Plus(ml.Measure()))
		return buildr(t.left, ml, l), x, buildl(r, mr, t.right)
	}

	l, x, r := splitDigit(t.right, pred, vchild)
	return buildr(t.left, t.child, l), x, ToFingerTreeComponent(r)
}

func (t ftreeTriple) split(pred func(Monoid) bool) (FingerTreeComponent, FingerTreeComponent) {
	if pred(t.Measure()) {
		l, x, r := t.splitTree(pred, Zero)
		return l, r.pushl(x)
	} else {
		return &t, &empty{}
	}
}

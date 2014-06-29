package fingerTree

func splitDigit(digit Slice, pred func(Monoid)bool, init Monoid) (Slice, Any, Slice) {
	var measure = init
	for i, a := range digit {
		measure = measure.Plus(Measure(a))
		if pred(measure) {
			return digit[:i], digit[i], digit[i+1:]
		}
	}
	return digit[:len(digit)-1], digit[len(digit)-1], Slice{}
}

func (e empty) splitTree(pred func(Monoid)bool, init Monoid) (FingerTree, Any, FingerTree) {
	panic("splitTree called on empty tree")
}

func (s single) splitTree(pred func(Monoid)bool, init Monoid) (FingerTree, Any, FingerTree) {
	return makeEmpty(), s.data, makeEmpty()
}

func (t ftree) splitTree(pred func(Monoid)bool, init Monoid) (FingerTree, Any, FingerTree) {
	vleft := init.Plus(t.left.Measure())
	if pred(vleft) {
		l, x, r := splitDigit(t.left, pred, init)
		return ToFingerTree(l), x, buildl(r, t.child, t.right)
	}

	vchild := vleft.Plus(t.child.Measure())
	if pred(vchild) {
		ml, mx, mr := t.child.splitTree(pred, vleft)
		l, x, r := splitDigit(mx.(node).ToSlice(), pred, vleft.Plus(ml.Measure()))
		return buildr(t.left, ml, l), x, buildl(r, mr, t.right)
	}

	l, x, r := splitDigit(t.right, pred, vchild)
	return buildr(t.left, t.child, l), x, ToFingerTree(r)
}


func (e empty) Split(pred func(Monoid)bool) (FingerTree, FingerTree) {
	return &empty{}, &empty{}
}

func (s single) Split(pred func(Monoid)bool) (FingerTree, FingerTree) {
	if pred(s.Measure()) {
		return &empty{}, &s
	} else {
		return &s, &empty{}
	}
}

func (t ftree) Split(pred func(Monoid)bool) (FingerTree, FingerTree) {
	if pred(t.Measure()) {
		l, x, r := t.splitTree(pred, Zero)
		return l, r.Pushl(x)
	} else {
		return &t, &empty{}
	}
}

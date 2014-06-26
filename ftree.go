package fingerTree

/**
 *	ftreeTriple structure
 */

type ftreeTriple struct {
	metadata map[string]Any
	left     Slice
	child    FingerTreeComponent
	right    Slice
}

func makeFTreeTriple(left Slice, child FingerTreeComponent, right Slice, mdataTypes mdataTypeMap) *ftreeTriple {
	meta := make(map[string]Any)
	for k, v := range mdataTypes {
		lz := mdataComposeFromSliceWithKey(left, k, mdataTypes)
		rz := mdataComposeFromSliceWithKey(right, k, mdataTypes)
		cz := child.mdataForKey(k, mdataTypes)
		sz := Slice{lz, rz, cz}.Foldl(func(a, b Any) Any { return v.compose(a, b) }, v.identity)
		meta[k] = sz
	}
	return &ftreeTriple{
		meta,
		left,
		child,
		right,
	}
}

func (t *ftreeTriple) mdataForKey(key string, mdataTypes mdataTypeMap) Any {
	return t.metadata[key]
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

func (t *ftreeTriple) Pushl(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	if len(t.left) < 4 {
		return makeFTreeTriple(
			append([]Any{d}, t.left...),
			t.child,
			t.right,
			mdataTypes,
		)
	}

	var child FingerTreeComponent
	pushdown := makeNode3(
		t.left[1],
		t.left[2],
		t.left[3],
		mdataTypes,
	)

	child = t.child.Pushl(pushdown, mdataTypes)

	return makeFTreeTriple(
		Slice{d, t.left[0]},
		child,
		t.right,
		mdataTypes,
	)
}

func (t *ftreeTriple) Popl(mdataTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return t.Taill(mdataTypes), t.Headl()
}

func (t *ftreeTriple) Popr(mdataTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return t.Tailr(mdataTypes), t.Headr()
}

func (t *ftreeTriple) Pushr(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	if len(t.right) < 4 {
		return makeFTreeTriple(
			t.left,
			t.child,
			append(t.right, d),
			mdataTypes,
		)
	}

	var child FingerTreeComponent
	pushdown := makeNode3(
		t.right[0],
		t.right[1],
		t.right[2],
		mdataTypes,
	)

	child = t.child.Pushr(pushdown, mdataTypes)

	return makeFTreeTriple(
		t.left,
		child,
		[]Any{t.right[3], d},
		mdataTypes,
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

func (t *ftreeTriple) Headr() Any {
	return t.right[len(t.right)-1]
}

func (t *ftreeTriple) Headl() Any {
	return t.left[0]
}

func buildr(left Slice, m FingerTreeComponent, right Slice, mdataTypes mdataTypeMap) FingerTreeComponent {
	if len(right) > 0 {
		return makeFTreeTriple(
			left,
			m,
			right,
			mdataTypes,
		)
	}

	if m.IsEmpty() {
		return ToFingerTreeComponent(left, mdataTypes)
	}

	return makeFTreeTriple(
		left,
		m.Tailr(mdataTypes),
		m.Headr().(node).ToSlice(),
		mdataTypes,
	)
}

func buildl(left Slice, m FingerTreeComponent, right Slice, mdataTypes mdataTypeMap) FingerTreeComponent {
	if len(left) > 0 {
		return makeFTreeTriple(
			left,
			m,
			right,
			mdataTypes,
		)
	}

	if m.IsEmpty() {
		return ToFingerTreeComponent(right, mdataTypes)
	}

	return makeFTreeTriple(
		m.Headl().(node).ToSlice(),
		m.Taill(mdataTypes),
		right,
		mdataTypes,
	)
}

func (t *ftreeTriple) Tailr(mdataTypes mdataTypeMap) FingerTreeComponent {
	return buildr(t.left, t.child, t.right[:len(t.right)-1], mdataTypes)
}

func (t *ftreeTriple) Taill(mdataTypes mdataTypeMap) FingerTreeComponent {
	return buildl(t.left[1:], t.child, t.right, mdataTypes)
}

func (t *ftreeTriple) IsEmpty() bool {
	return false
}

func (t *ftreeTriple) Concatr(other FingerTreeComponent, mdataTypes mdataTypeMap) FingerTreeComponent {
	otherAsFtree, isFTree := other.(*ftreeTriple)
	if !isFTree {
		return other.Concatl(t, mdataTypes)
	}

	return glue(t, Slice{}, otherAsFtree, mdataTypes)
}

func (t *ftreeTriple) Concatl(other FingerTreeComponent, mdataTypes mdataTypeMap) FingerTreeComponent {
	return other.Concatr(t, mdataTypes)
}

func (t *ftreeTriple) ToSlice() Slice {
	return ToSlice(t)
}

package fingerTree23

type Any interface{}

type FoldFunc func(a Any, b Any) Any
type Foldable interface {
	Foldl(f FoldFunc, initial Any) Any
}

type FingerTree interface {
	Foldable

	Pushf(d Any) FingerTree
	Pushb(d Any) FingerTree
}

type node interface {
	Foldable
}

type node2 struct {
	data [2]Any
}

func (n node2) Foldl(f FoldFunc, initial Any) Any {
	return Foldl(f, initial, n.data[:])
}

type node3 struct {
	data [3]Any
}

func (n node3) Foldl(f FoldFunc, initial Any) Any {
	return Foldl(f, initial, n.data[:])
}

type ftree struct {
	left  []Any
	right []Any
	child FingerTree
}

func (t ftree) Foldl(f FoldFunc, initial Any) Any {
	lift := func(init Any, data Any) Any {
		n := data.(node)
		return n.Foldl(f, init)
	}

	var a interface{} = Foldl(f, initial, t.left)
	var b interface{} = t.child.Foldl(lift, a)
	return Foldl(f, b, t.right)
}

func (t ftree) Pushf(d Any) FingerTree {
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

	child = t.child.Pushf(pushdown)

	return &ftree{
		[]Any{d, t.left[0]},
		t.right,
		child,
	}
}

func (t ftree) Pushb(d Any) FingerTree {
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

	child = t.child.Pushb(pushdown)

	return &ftree{
		t.left,
		[]Any{t.right[3], d},
		child,
	}
}

type single struct {
	data Any
}

func (s single) Foldl(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}
func (s single) Pushf(d Any) FingerTree {
	return &ftree{[]Any{d, s.data}, []Any{}, empty{}}
}

func (s single) Pushb(d Any) FingerTree {
	return &ftree{[]Any{s.data, d}, []Any{}, empty{}}
}

type empty struct {}

func (tree empty) Foldl(f FoldFunc, initial Any) Any {
	return initial
}

func (tree empty) Pushf(d Any) FingerTree {
	return &single{d};
}

func (tree empty) Pushb(d Any) FingerTree {
	return &single{d};
}

func ToSlice(t FingerTree) []Any {
	app := func(a Any, b Any) Any {
		return append(a.([]Any), b)
	}
	return t.Foldl(app, make([]Any, 0)).([]Any)
}

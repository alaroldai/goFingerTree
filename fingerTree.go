package fingerTree23

type Data interface{}

type FoldFunc func(interface{}, Data) interface{}
type Foldable interface {
	Foldl(f FoldFunc, initial interface{}) interface{}
}

type BackQueue interface {
}

type FingerTree interface {
	Foldable

	Pushf(d Data) FingerTree
	Pushb(d Data) FingerTree
}

type node interface {
	Foldable
}

type node2 struct {
	data [2]Data
}

func (n node2) Foldl(f FoldFunc, initial interface{}) interface{} {
	return Foldl(f, initial, n.data[0:], 2)
}

type node3 struct {
	data [3]Data
}

func (n node3) Foldl(f FoldFunc, initial interface{}) interface{} {
	return Foldl(f, initial, n.data[0:], 3)
}

type ftree struct {
	left  []Data
	right []Data
	child FingerTree
}

func (t ftree) Foldl(f FoldFunc, initial interface{}) interface{} {
	lift := func(init interface{}, data Data) interface{} {
		n := data.(node)
		return n.Foldl(f, init)
	}

	var lleft int = len(t.left)
	var lright int = len(t.right)

	var a interface{} = Foldl(f, initial, t.left, lleft)
	var b interface{}
	if t.child != nil {
		b = t.child.Foldl(lift, a)
	} else {
		b = a
	}
	return Foldl(f, b, t.right, lright)
}

func (t ftree) Pushf(d Data) FingerTree {
	if len(t.left) < 4 {
		return &ftree{
			append([]Data{d}, t.left...),
			t.right,
			t.child,
		}
	}

	var child FingerTree
	pushdown := &node3{
		[3]Data{
			t.left[1],
			t.left[2],
			t.left[3],
		},
	}

	if t.child != nil {
		child = t.child.Pushf(pushdown)
	} else {
		child = &single{pushdown}
	}

	return &ftree{
		[]Data{d, t.left[0]},
		t.right,
		child,
	}
}

func (t ftree) Pushb(d Data) FingerTree {
	if len(t.right) < 4 {
		return &ftree{
			t.left,
			append(t.right, []Data{d}),
			t.child,
		}
	}

	var child FingerTree
	pushdown := &node3{
		[3]Data{
			t.right[0],
			t.right[1],
			t.right[2],
		},
	}

	if t.child != nil {
		child = t.child.Pushb(pushdown)
	} else {
		child = &single{pushdown}
	}
	return &ftree{
		t.left,
		[]Data{t.right[3], d},
		child,
	}
}

type single struct {
	data Data
}

func (s single) Foldl(f FoldFunc, initial interface{}) interface{} {
	return f(initial, s.data)
}

func (s single) Pushf(d Data) FingerTree {
	return &ftree{[]Data{d, s.data}, []Data{}, nil}
}

func (s single) Pushb(d Data) FingerTree {
	return &ftree{[]Data{s.data, d}, []Data{}, nil}
}

func ToSlice(t FingerTree) []Data {
	app := func(a interface{}, b Data) interface{} {
		_, s := b.([]Data)
		if s {
			return append(a.([]Data), b.([]Data)...)
		}
		return append(a.([]Data), b)
	}
	return t.Foldl(app, make([]Data, 0)).([]Data)
}
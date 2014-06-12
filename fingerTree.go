package fingerTree23

type Data interface{}

type IterFunc func(Data)
type Iterable interface {
	Iterl(IterFunc)
	Iterr(IterFunc)
}

type FoldFunc func(interface{}, Data) interface{}
type Foldable interface {
	Iterable
	Foldl(f FoldFunc, initial interface{}) interface{}
	Foldr(f FoldFunc, initial interface{}) interface{}
}

type FingerTree interface {
	Foldable

	Pushl(d Data) FingerTree
	Pushr(d Data) FingerTree

	Popl() (FingerTree, Data)
}

type node interface {
	Foldable
	toDigit() []Data
}

type node2 struct {
	data [2]Data
}

func (n node2) toDigit() []Data {
	return n.data[:]
}

func (n node2) Foldl(f FoldFunc, initial interface{}) interface{} {
	return Foldl(f, initial, n.data[:], 2)
}

func (n node2) Foldr(f FoldFunc, initial interface{}) interface{} {
	return Foldr(f, initial, n.data[:], 2)
}

func (n node2) Iterl(f IterFunc) {
	Iterl(f, n.data[:], len(n.data))
}

func (n node2) Iterr(f IterFunc) {
	Iterr(f, n.data[:], len(n.data))
}

type node3 struct {
	data [3]Data
}

func (n node3) toDigit() []Data {
	return n.data[:]
}

func (n node3) Foldl(f FoldFunc, initial interface{}) interface{} {
	return Foldl(f, initial, n.data[:], 3)
}

func (n node3) Foldr(f FoldFunc, initial interface{}) interface{} {
	return Foldr(f, initial, n.data[:], 3)
}

func (n node3) Iterl(f IterFunc) {
	Iterl(f, n.data[:], len(n.data))
}

func (n node3) Iterr(f IterFunc) {
	Iterr(f, n.data[:], len(n.data))
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
	var b interface{} = t.child.Foldl(lift, a)
	return Foldl(f, b, t.right, lright)
}

func (t ftree) Foldr(f FoldFunc, initial interface{}) interface{} {
	lift := func(init interface{}, data Data) interface{} {
		n := data.(node)
		return n.Foldr(f, init)
	}

	lleft := len(t.left)
	lright := len(t.right)

	a := Foldr(f, initial, t.right, lright)
	b := t.child.Foldr(lift, a)
	return Foldr(f, b, t.left, lleft)
}

func (t ftree) Pushl(d Data) FingerTree {
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

	child = t.child.Pushl(pushdown)

	return &ftree{
		[]Data{d, t.left[0]},
		t.right,
		child,
	}
}

func (t ftree) Popl() (FingerTree, Data) {
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
				p.(node).toDigit(),
				t.right,
				nc,
			},
			t.left[0]
	}
}

func (t ftree) Pushr(d Data) FingerTree {
	if len(t.right) < 4 {
		return &ftree{
			t.left,
			append(t.right, d),
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

	child = t.child.Pushr(pushdown)

	return &ftree{
		t.left,
		[]Data{t.right[3], d},
		child,
	}
}

func (t ftree) Iterl(f IterFunc) {
	t.Foldl(func(_ interface{}, b Data) interface{} {
		f(b)
		return nil
	}, nil)
}

func (t ftree) Iterr(f IterFunc) {
	t.Foldr(func(_ interface{}, b Data) interface{} {
		f(b)
		return nil
	}, nil)
}

type single struct {
	data Data
}

func (s single) Foldl(f FoldFunc, initial interface{}) interface{} {
	return f(initial, s.data)
}

func (s single) Foldr(f FoldFunc, initial interface{}) interface{} {
	return f(initial, s.data)
}

func (s single) Pushl(d Data) FingerTree {
	return &ftree{[]Data{d, s.data}, []Data{}, empty{}}
}

func (s single) Popl() (FingerTree, Data) {
	return &empty{}, s.data
}

func (s single) Pushr(d Data) FingerTree {
	return &ftree{[]Data{s.data, d}, []Data{}, empty{}}
}

func (s single) Iterl(f IterFunc) {
	f(s.data)
}

func (s single) Iterr(f IterFunc) {
	f(s.data)
}

type empty struct{}

func (tree empty) Foldl(f FoldFunc, initial interface{}) interface{} {
	return initial
}

func (e empty) Foldr(f FoldFunc, initial interface{}) interface{} {
	return initial
}

func (tree empty) Pushl(d Data) FingerTree {
	return &single{d}
}

func (e empty) Popl() (FingerTree, Data) {
	return &empty{}, nil
}

func (tree empty) Pushr(d Data) FingerTree {
	return &single{d}
}

func (e empty) Iterl(f IterFunc) {
	return
}

func (e empty) Iterr(f IterFunc) {
	return
}

func ToSlice(t FingerTree) []Data {
	app := func(a interface{}, b Data) interface{} {
		return append(a.([]Data), b)
	}
	return t.Foldl(app, make([]Data, 0)).([]Data)
}

package fingerTree23

type Data interface{}

type FoldFunc func(interface{}, Data) interface{}
type Foldable interface {
	Foldl(f FoldFunc, initial interface{}) interface{}
}

type FingerTree interface {
	Foldable
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

	a := Foldl(f, initial, t.left, lleft)
	b := t.child.Foldl(lift, a)
	return Foldl(f, b, t.right, lright)
}

type single struct {
	data Data
}

func (s single) Foldl(f FoldFunc, initial interface{}) interface{} {
	return f(initial, s)
}

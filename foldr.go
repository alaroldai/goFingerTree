package fingerTree23

func Foldr(f FoldFunc, s []Any, initial Any) Any {
	if len(s) > 0 {
		return f(s[0], Foldr(f, s[1:], initial))
	} else {
		return initial
	}
}

func (n node2) Foldr(f FoldFunc, initial Any) Any {
	return Foldr(f, n.data[:], initial)
}


func (n node3) Foldr(f FoldFunc, initial Any) Any {
	return Foldr(f, n.data[:], initial)
}


// -- foldr for finger trees --

func (e empty) Foldr(f FoldFunc, initial Any) Any {
	return initial
}

func (s single) Foldr(f FoldFunc, initial Any) Any {
	return f(s.data, initial)
}

func (t ftree) Foldr(f FoldFunc, initial Any) Any {
	lift := func(data Any, init Any) Any {
		n := data.(node)
		return n.Foldr(f, init)
	}

	var a interface{} = Foldr(f, t.right, initial)
	var b interface{} = t.child.Foldr(lift, a)
	return Foldr(f, t.left, b)
}

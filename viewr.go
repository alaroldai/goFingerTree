package fingerTree23

// -- Head --

func (e empty) Headr() Any {
	return nil
}

func (s single) Headr() Any {
	return s.data
}

func (t ftree) Headr() Any {
	return t.right[len(t.right)-1]
}

// -- Tail --

func (e empty) Tailr() FingerTree {
	// Not sure if this makes sense
	return nil
}

func (s single) Tailr() FingerTree {
	return empty{}
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
		m.Headr().(Sliceable).ToSlice(),
		m.Tailr(),
	}
}

func (t ftree) Tailr() FingerTree {
	return buildr(t.left, t.child, t.right[:len(t.right)-1])
}

func (e empty) IsEmpty() bool {
	return true
}
func (s single) IsEmpty() bool {
	return false
}
func (t ftree) IsEmpty() bool {
	return false
}

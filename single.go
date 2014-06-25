package fingerTree

type single struct {
	data Any
}

func (s *single) Foldl(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Foldr(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Pushl(d Any) FingerTree {
	return &ftree{
		[]Any{d},
		&empty{},
		[]Any{s.data},
	}
}

func (s *single) Popl() (FingerTree, Any) {
	return &empty{}, s.data
}

func (s *single) Popr() (FingerTree, Any) {
	return &empty{}, s.data
}

func (s *single) Pushr(d Any) FingerTree {
	return &ftree{
		[]Any{s.data},
		&empty{},
		[]Any{d},
	}
}

func (s *single) Iterl(f IterFunc) {
	f(s.data)
}

func (s *single) Iterr(f IterFunc) {
	f(s.data)
}

func (s *single) Headr() Any {
	return s.data
}

func (s *single) Headl() Any {
	return s.data
}

func (s *single) Tailr() FingerTree {
	return &empty{}
}

func (s *single) Taill() FingerTree {
	return &empty{}
}

func (s *single) IsEmpty() bool {
	return false
}

// Concat t to the right of the receiver
func (s *single) Concatr(t FingerTree) FingerTree {
	return t.Pushl(s.data)
}

func (s *single) Concatl(t FingerTree) FingerTree {
	return t.Pushr(s.data)
}

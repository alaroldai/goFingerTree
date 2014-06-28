package fingerTree

type single struct {
	measure Monoid
	data Any
}

func makeSingle(d Any) *single {
	return &single{Measure(d), d}
}

func (s *single) Measure() Monoid {
	return s.measure
}

func (s *single) Foldl(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Foldr(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Pushl(d Any) FingerTree {
	return makeFTree(
		[]Any{d},
		makeEmpty(),
		[]Any{s.data},
	)
}

func (s *single) Popl() (FingerTree, Any) {
	return makeEmpty(), s.data
}

func (s *single) Popr() (FingerTree, Any) {
	return makeEmpty(), s.data
}

func (s *single) Pushr(d Any) FingerTree {
	return makeFTree(
		[]Any{s.data},
		makeEmpty(),
		[]Any{d},
	)
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
	return makeEmpty()
}

func (s *single) Taill() FingerTree {
	return makeEmpty()
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

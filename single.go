package fingerTree

type single struct {
	data Any
}

func makeSingle(d Any) *single {
	return &single{d}
}

func (s *single) Measure() Monoid {
	return Measure(s.data)
}

func (s *single) Foldl(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) Foldr(f FoldFunc, initial Any) Any {
	return f(initial, s.data)
}

func (s *single) pushl(d Any) FingerTreeComponent {
	return makeFTreeTriple(
		[]Any{d},
		makeEmpty(),
		[]Any{s.data},
	)
}

func (s *single) popl() (FingerTreeComponent, Any) {
	return makeEmpty(), s.data
}

func (s *single) popr() (FingerTreeComponent, Any) {
	return makeEmpty(), s.data
}

func (s *single) pushr(d Any) FingerTreeComponent {
	return makeFTreeTriple(
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

func (s *single) headr() Any {
	return s.data
}

func (s *single) headl() Any {
	return s.data
}

func (s *single) tailr() FingerTreeComponent {
	return makeEmpty()
}

func (s *single) taill() FingerTreeComponent {
	return makeEmpty()
}

func (s *single) isEmpty() bool {
	return false
}

// concat t to the right of the receiver
func (s *single) concatr(t FingerTreeComponent) FingerTreeComponent {
	return t.pushl(s.data)
}

func (s *single) concatl(t FingerTreeComponent) FingerTreeComponent {
	return t.pushr(s.data)
}

func (s single) splitTree(pred func(Monoid) bool, init Monoid) (FingerTreeComponent, Any, FingerTreeComponent) {
	return makeEmpty(), s.data, makeEmpty()
}

func (s single) split(pred func(Monoid) bool) (FingerTreeComponent, FingerTreeComponent) {
	if pred(s.Measure()) {
		return &empty{}, &s
	} else {
		return &s, &empty{}
	}
}

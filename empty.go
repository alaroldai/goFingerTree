package fingerTree

type empty struct{}

func makeEmpty() *empty {
	return &empty{}
}

func (e *empty) Measure() Monoid {
	return Zero
}

func (e *empty) Foldl(f FoldFunc, initial Any) Any {
	return initial
}

func (e *empty) Foldr(f FoldFunc, initial Any) Any {
	return initial
}

func (e *empty) pushl(d Any) FingerTreeComponent {
	return makeSingle(d)
}

func (e *empty) popl() (FingerTreeComponent, Any) {
	return makeEmpty(), nil
}

func (e *empty) popr() (FingerTreeComponent, Any) {
	return makeEmpty(), nil
}

func (e *empty) pushr(d Any) FingerTreeComponent {
	return makeSingle(d)
}

func (e *empty) Iterl(f IterFunc) {
	return
}

func (e *empty) Iterr(f IterFunc) {
	return
}

func (e *empty) headr() Any {
	return nil
}

func (e *empty) headl() Any {
	return nil
}

func (e *empty) tailr() FingerTreeComponent {
	// Not sure if this makes sense
	return nil
}

func (e *empty) taill() FingerTreeComponent {
	return nil
}

func (e *empty) isEmpty() bool {
	return true
}

func (e *empty) concatr(t FingerTreeComponent) FingerTreeComponent {
	return t
}

func (e *empty) concatl(t FingerTreeComponent) FingerTreeComponent {
	return t
}

func (e empty) splitTree(pred func(Monoid) bool, init Monoid) (FingerTreeComponent, Any, FingerTreeComponent) {
	panic("splitTree called on empty tree")
}

func (e empty) split(pred func(Monoid) bool) (FingerTreeComponent, FingerTreeComponent) {
	return &empty{}, &empty{}
}

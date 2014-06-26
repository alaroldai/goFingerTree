package fingerTree

type empty struct{}

func makeEmpty() *empty {
	return &empty{}
}

func (e *empty) mdataForKey(key string, mdataTypes mdataTypeMap) Any {
	return mdataTypes[key].identity
}

func (e *empty) Foldl(f FoldFunc, initial Any) Any {
	return initial
}

func (e *empty) Foldr(f FoldFunc, initial Any) Any {
	return initial
}

func (e *empty) Pushl(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	return makeSingle(d, mdataTypes)
}

func (e *empty) Popl(mdataTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return makeEmpty(), nil
}

func (e *empty) Popr(mdataTypes mdataTypeMap) (FingerTreeComponent, Any) {
	return makeEmpty(), nil
}

func (e *empty) Pushr(d Any, mdataTypes mdataTypeMap) FingerTreeComponent {
	return makeSingle(d, mdataTypes)
}

func (e *empty) Iterl(f IterFunc) {
	return
}

func (e *empty) Iterr(f IterFunc) {
	return
}

func (e *empty) Headr() Any {
	return nil
}

func (e *empty) Headl() Any {
	return nil
}

func (e *empty) Tailr(mdataTypes mdataTypeMap) FingerTreeComponent {
	// Not sure if this makes sense
	return nil
}

func (e *empty) Taill(mdataTypes mdataTypeMap) FingerTreeComponent {
	return nil
}

func (e *empty) IsEmpty() bool {
	return true
}

func (e *empty) Concatr(t FingerTreeComponent, mdataTypes mdataTypeMap) FingerTreeComponent {
	return t
}

func (e *empty) Concatl(t FingerTreeComponent, mdataTypes mdataTypeMap) FingerTreeComponent {
	return t
}

func (e *empty) ToSlice() Slice {
	return ToSlice(e)
}

package fingerTree

type Measured interface {
	Measure() Monoid
}

// Uses the trivial measure (which returns Zero) for non-Measured things
func Measure(m Any) Monoid {
	ms, success := m.(Measured)
	if success {
		return ms.Measure()
	} else {
		return Zero
	}
}

// The free measure, that just makes a list of its contents
type mfree struct {
	v Any
}

func (c mfree) Measure() Monoid {
	return Slice{c.v}
}

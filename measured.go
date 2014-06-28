package fingerTree

type Measured interface {
	Measure() Monoid
}

func Measure(m Any) Monoid {
	ms, success := m.(Measured)
	if success {
		return ms.Measure()
	} else {
		return Zero
	}
}

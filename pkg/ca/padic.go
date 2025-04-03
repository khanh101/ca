package ca

// PAdic : p-adic integers
type PAdic interface {
	Get(int) int
	Add(PAdic) PAdic
	Neg() PAdic
	Sub(PAdic) PAdic
	Iter() Iter
	Mul(PAdic) PAdic
	Norm() int
	Approx(n int) (int, []int)
}

type padic struct {
	prime int
	iter  func() int
	cache []int
}

func NewPAdic(prime int, iter func() int) PAdic {
	carry := 0
	return &padic{
		prime: prime,
		iter: func() int {
			v := iter()
			q, r := divmod(carry+v, prime)
			carry = q
			return r
		},
		cache: nil,
	}
}

func NewPAdicFromInt(prime int, v int) PAdic {
	return NewPAdic(prime, NewIterFromList([]int{v}, 0))
}

func (a *padic) Get(i int) int {
	for len(a.cache) <= i {
		a.cache = append(a.cache, a.iter())
	}
	return a.cache[i]
}

func (a *padic) Add(B PAdic) PAdic {
	b := B.(*padic)
	if a.prime != b.prime {
		panic("different bases")
	}
	i := 0
	carry := 0
	return &padic{
		prime: a.prime,
		iter: func() int {
			q, r := divmod(carry+a.Get(i)+b.Get(i), a.prime)
			carry = q
			i++
			return r
		},
	}
}

func (a *padic) Neg() PAdic {
	i := 0
	return (&padic{
		prime: a.prime,
		iter: func() int {
			r := a.prime - a.Get(i) - 1
			i++
			return r
		},
	}).Add(&padic{
		prime: a.prime,
		iter:  NewIterFromList([]int{1}, 0),
	})
}

func (a *padic) Sub(B PAdic) PAdic {
	return a.Add(B.Neg())
}

func (a *padic) Iter() Iter {
	i := 0
	return func() int {
		v := a.Get(i)
		i++
		return v
	}
}

func (a *padic) mulDigit(b int) PAdic {
	i := 0
	carry := 0
	return &padic{
		prime: a.prime,
		iter: func() int {
			q, r := divmod(carry+a.Get(i)*b, a.prime)
			carry = q
			i++
			return r
		},
	}
}

func (a *padic) Mul(B PAdic) PAdic {
	b := B.(*padic)
	if a.prime != b.prime {
		panic("different bases")
	}
	var partial []PAdic
	i := 0
	carry := 0
	return NewPAdic(a.prime, func() int {
		partial = append(partial, NewPAdic(
			a.prime,
			shift(i, 0, a.mulDigit(b.Get(i)).Iter())),
		)
		s := 0
		for _, p := range partial {
			s += p.Get(i)
		}
		q, r := divmod(carry+s, a.prime)
		carry = q
		i++
		return r
	})
}

func (a *padic) Approx(n int) (int, []int) {
	approx := make([]int, n)
	s := 0
	x := 1
	for i := 0; i < n; i++ {
		s += a.Get(i) * x
		x *= a.prime
		approx[i] = a.Get(i)
	}
	return s, approx
}

func (a *padic) Norm() int {
	i := 0
	for {
		v := a.Get(i)
		if v != 0 {
			return i
		}
		i++
	}
}

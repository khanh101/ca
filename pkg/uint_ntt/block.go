package uint_ntt

const (
	base  = 1 << 16 // pick base = 2^d, max_n * base * base < p so that multiplication won't overflow
	max_n = 4294967294
)

// Block : polynomial in F_p[X]
type Block []uint64

func (b Block) get(i int) uint64 {
	if i >= len(b) {
		return 0
	}
	return b[i]
}

func (b Block) set(i int, v uint64) Block {
	for i >= len(b) {
		b = append(b, 0)
		if len(b) > max_n {
			panic("too many blocks")
		}
	}
	b[i] = v
	return b
}

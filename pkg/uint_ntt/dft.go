package uint_ntt

import (
	"ca/pkg/vec"
	"math/bits"
)

// CooleyTukeyFFT :Cooley-Tukey algorithm
func CooleyTukeyFFT(x vec.Vec[uint64], omega uint64) vec.Vec[uint64] {
	n := x.Len()
	if n == 1 {
		return x
	}
	if n <= 0 || n%2 != 0 {
		panic("n must be power of 2")
	}
	// even and odd values of x
	e, o := vec.Make[uint64](n/2), vec.Make[uint64](n/2)
	for i := 0; i < n/2; i++ {
		e = e.Set(i, x.Get(2*i))
		o = o.Set(i, x.Get(2*i+1))
	}
	omega_2 := mul(omega, omega)
	eFFT := CooleyTukeyFFT(e, omega_2)
	oFFT := CooleyTukeyFFT(o, omega_2)

	y := vec.Make[uint64](n)
	var omega_n uint64 = 1 // omega^0
	for i := 0; i < n/2; i++ {
		t := mul(omega_n, oFFT.Get(i))
		y = y.Set(i, add(eFFT.Get(i), t))
		y = y.Set(i+n/2, sub(eFFT.Get(i), t))
		omega_n = mul(omega_n, omega)
	}
	return y
}

// nextPowerOfTwo : return the smallest power of 2 greater than x (code generated by chatgpt)
func nextPowerOfTwo(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	if x > 1<<63 {
		panic("next power of 2 overflows uint64")
	}
	return 1 << (64 - bits.LeadingZeros64(x-1))
}

func time2freq(time vec.Vec[uint64], length uint64) vec.Vec[uint64] {
	// extend  into powers of 2
	n := nextPowerOfTwo(length)
	time = time.Slice(0, int(n)) // extend to length n

	omega := getPrimitiveRoot(n)
	freq := trimZeros(CooleyTukeyFFT(time, omega))
	return freq
}

func freq2time(freq vec.Vec[uint64], length uint64) vec.Vec[uint64] {
	// extend  into powers of 2
	n := nextPowerOfTwo(length)
	freq = freq.Slice(0, int(n)) // extend to length n
	omega := getPrimitiveRoot(n)
	il := inv(n)

	time := CooleyTukeyFFT(freq, inv(omega))
	for i := 0; i < time.Len(); i++ {
		f := time.Get(i)
		time = time.Set(i, mul(f, il))
	}

	time = trimZeros(time)
	return time
}

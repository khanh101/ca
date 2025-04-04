package main

import (
	"ca/pkg/ca"
	"fmt"
)

const N = 10

func main() {
	x := ca.NewPAdicFromInt(3, 23)
	y := ca.NewPAdicFromInt(3, 27)
	z := ca.NewPAdicFromInt(3, 92)
	fmt.Println(x.Approx(N))
	fmt.Println(y.Approx(N))
	fmt.Println(y.Mul(x).Approx(N)) // print 27 x 23 = 621
	fmt.Println(y.Sub(x).Approx(N)) // print 27 - 23 = 4
	fmt.Println(z.Div(x).Approx(N)) // print 92 / 23 = 4
}

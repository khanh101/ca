package main

import (
	"ca/pkg/fib"
	"ca/pkg/integer"
	"ca/pkg/padic"
	"ca/pkg/uint1792"
	"fmt"
	"os"
	"strconv"
)

func testPAdic() {
	const N = 10
	x := padic.NewPAdicFromInt(3, 23)
	y := padic.NewPAdicFromInt(3, 27)
	z := padic.NewPAdicFromInt(3, 92)
	fmt.Println(x.Approx(N))
	fmt.Println(y.Approx(N))
	fmt.Println(y.Mul(x).Approx(N)) // print 27 x 23 = 621
	fmt.Println(y.Sub(x).Approx(N)) // print 27 - 23 = 4
	fmt.Println(z.Div(x).Approx(N)) // print 92 / 23 = 4
}

func testUint1792() {
	x := uint1792.FromString("0x318346193417412890342342")
	z := uint1792.FromString("0x484723895378245789")
	y := uint1792.FromString(x.String())
	fmt.Println(x.Add(z).Mod(y)) // (x + integer) % x
}

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <n>")
	}

	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(fib.Fib(integer.Zero, uint64(n)))
}

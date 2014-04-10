package main

import (
	"euler"
	"fmt"
	"math"
	"math/big"
	"time"
)

func ctdFrac(list []int) (string, string) {
	current := big.NewRat(int64(list[len(list)-1]), 1)
	temp := big.NewRat(0, 1)

	for i := len(list) - 2; i >= 0; i-- {
		temp.SetFrac64(int64(list[i]), 1)
		current.Inv(current)

		current.Add(current, temp)
	}

	return current.Num().String(), current.Denom().String()
}

//fraction of the form (a +  b \sqrt R) / d
type frac struct {
	a int
	b int
	R int
	d int
}

func flip(F frac) frac {
	return frac{
		a: F.a * F.d,
		b: -F.b * F.d,
		R: F.R,
		d: (F.a * F.a) - (F.b * F.b * F.R)}
}

func reduce(F frac) frac {
	gcd := int(euler.GCD(int64(F.a), int64(F.b)))
	gcd = int(euler.GCD(int64(gcd), int64(F.d)))
	return frac{F.a / gcd, F.b / gcd, F.R, F.d / gcd}
}

func nextFrac(F frac) (n int, next frac) {
	total := (float64(F.a) + (float64(F.b) * math.Sqrt(float64(F.R)))) / float64(F.d)

	n = int(total)

	if n != 0 {
		next = frac{
			a: F.a - (n * F.d),
			b: F.b,
			R: F.R,
			d: F.d}
	} else {
		n = 1
		next = frac{
			a: F.a + F.d,
			b: F.b,
			R: F.R,
			d: F.d}
	}

	next = flip(next)
	next = reduce(next)

	return
}

func isSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return true
	}
	return false
}

//"Pell Equation"
func main() {
	starttime := time.Now()

	record, bestD := 0, 0

	for rad := 2; rad <= 1000; rad++ {

		if !isSquare(rad) {

			convergentList := make([]int, 1)

			test := frac{a: 0, b: 1, R: rad, d: 1}
			n := 0

			n, test = nextFrac(test)

			convergentList[0] = n

			for i := 0; n != 2*convergentList[0]; i++ {
				n, test = nextFrac(test)
				convergentList = append(convergentList, n)

			}

			fLength := len(convergentList) - 1

			x := ""

			if len(convergentList)%2 == 1 {
				x, _ = ctdFrac(convergentList[:fLength])

			} else {

				newt := make([]int, 2*fLength)
				copy(newt[:fLength+1], convergentList)
				copy(newt[fLength+1:(2*fLength)], convergentList[1:fLength])
				x, _ = ctdFrac(newt[:len(newt)])

			}

			if len(x) >= record {
				fmt.Println("d=", rad, ":", x)
				record, bestD = len(x), rad
			}

		}
	}

	fmt.Printf("%d\n", bestD)

	fmt.Println("Elapsed time:", time.Since(starttime))

}

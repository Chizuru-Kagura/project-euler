package euler

import "math"

const (
	primeTableLength = 50000000 //+1!!!!!!!!!!!
	//lastPrime = Prime[primeTableLength - 1]
	lastPrime          = 982451629
	totientTableLength = 100000
)

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64
var totientTable [totientTableLength]int64
var factorialtable = make(map[int64]int64)

func Divisors(n int64) int64 {
	factors := Factors(n)
	div := int64(1)
	for i := 0; i < len(factors); i++ {
		div *= factors[i][1] + 1
	}
	return div
}

//returns an ordered list of distinct factors
func Factor(n int64) []int64 {
	var answer = []int64{}

	current := n

	i := int64(1)
	for !IsPrime(current) {
		if current%Prime(i) == 0 {
			answer = append(answer, Prime(i))
			current = current / Prime(i)
			i = 0
		}
		i++
	}

	answer = append(answer, current)

	return answer
}

//ax + by = gcd(a,b)
func ExtendedEuclidean(a, b int64) (x, y int64) {
	x, lastx := int64(0), int64(1)
	y, lasty := int64(1), int64(0)
	for b != int64(0) {
		quotient := a / b
		a, b = b, a%b
		x, lastx = lastx-quotient*x, x
		y, lasty = lasty-quotient*y, y
	}
	return lastx, lasty
}

//what's X^-1 mod n? (assuming of course x, n coprime)
func InverseMod(x, n int64) int64 {
	ans, _ := ExtendedEuclidean(x, n)
	return ans
}

//find a number equal to a mod n. N are assumed to be coprime
func ChineseRemainder(a, n []int64) int64 {
	N := int64(1)
	for _, en := range n {
		N *= en
	}

	ans := int64(0)
	for i := range a {
		summand := a[i]
		summand *= N
		summand /= n[i]
		summand *= InverseMod(N/n[i], n[i])

		ans += summand
	}

	return ans

}

//an ordered list of prime factors, together with their degrees
func Factors(n int64) [][2]int64 {
	factorList := Factor(n)
	factors := [][2]int64{[2]int64{factorList[0], 1}}
	for i := 1; i < len(factorList); i++ {
		if factorList[i] == factors[len(factors)-1][0] {
			factors[len(factors)-1][1]++
		} else {
			factors = append(factors, [2]int64{factorList[i], 1})
		}
	}
	return factors
}

func Factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	if answer, ok := factorialtable[n]; ok {
		return answer
	}

	answer := Factorial(n-1) * n

	factorialtable[n] = answer

	return answer
}

func IntSqrt(n int64) (sqrt int64, square bool) {
	sqrt = int64(math.Sqrt(float64(n)))
	square = sqrt*sqrt == n
	return
}

//number of primes less than or equal to n
func PrimePi(n int64) int64 {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n < primeTableLength && primepilist[n] != 0 {
		return primepilist[n]
	}

	var answer int64
	if IsPrime(n) {
		answer = 1 + PrimePi(n-1)
		if answer < primeTableLength {
			primeTable[answer] = n
		}
	} else {
		answer = PrimePi(n - 1)

	}

	if n < primeTableLength {
		primepilist[n] = answer
	}
	return answer
}

//Finds/caches all primes below n using a seive
func PrimeCache(n int64) {
	composite := make([]bool, n)
	composite[0], composite[1] = true, true

	p := int64(2)

	for p < n {
		marker := p + p
		for ; marker < n; marker += p {
			composite[marker] = true
		}
		p++
		for ; p < n && composite[p]; p++ {
		}
	}

	m := 1
	for i := int64(0); i < n; i++ {
		if !composite[i] {
			primeTable[m] = i
			primepilist[i] = int64(m)
			m++
		}
	}
}

//if p is the nth prime
func PrimeN(p int64) int64 {
	if !IsPrime(p) {
		return -1
	}
	return PrimePi(p)
}

func IsPrime(n int64) bool {

	if n == 1 {
		return false
	}

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	//If we need to pass the end of the prime table brute force
	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}

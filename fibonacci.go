package main

var fibonacciRecIterCount = 0
var fibonacciRecInstCount = 0

func fibonacciRec(n uint) uint {
	defer func() {
		fibonacciRecIterCount++
		fibonacciRecInstCount++ // switch comparison
	}()

	switch {
	case n <= 1:
		return n
	default:
		a := fibonacciRec(n - 1)
		b := fibonacciRec(n - 2)
		fibonacciRecInstCount++ // n-1
		fibonacciRecInstCount++ // set a
		fibonacciRecInstCount++ // n-2
		fibonacciRecInstCount++ // set b

		defer func() {
			fibonacciRecInstCount++ // sum a+b
		}()

		return a + b
	}
}

func fibonacciDyn(n uint) uint {
	f := make([]uint, n+1)
	f[0], f[1] = 0, 1
	fibonacciRecInstCount++ // find in 0
	fibonacciRecInstCount++ // assign to 0
	fibonacciRecInstCount++ // find in 1
	fibonacciRecInstCount++ // assign to 1

	for i := uint(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
		fibonacciRecInstCount++ // find position i
		fibonacciRecInstCount++ // i-1
		fibonacciRecInstCount++ // find position i-1
		fibonacciRecInstCount++ // i-2
		fibonacciRecInstCount++ // find position i-2
		fibonacciRecInstCount++ // find position i
		fibonacciRecInstCount++ // assign

		fibonacciRecIterCount++
	}

	defer func() {
		fibonacciRecInstCount++ // return (access index)
	}()
	return f[n]
}

func fibonacciMemo(n uint) uint {
	f := make([]*uint, n+1)
	return lookupFibonacci(f, n)
}

func lookupFibonacci(f []*uint, n uint) uint {
	defer func() {
		fibonacciRecIterCount++
		fibonacciRecInstCount++ // base case n index access
		fibonacciRecInstCount++ // base case comparison
	}()

	if f[n] != nil {
		defer func() {
			fibonacciRecInstCount++ // n index access
		}()
		return *f[n]
	}

	defer func() {
		fibonacciRecInstCount++ // comparison
	}()

	switch {
	case n <= 1:
		f[n] = newUint(n)

		fibonacciRecInstCount++ // n index access
		fibonacciRecInstCount++ // assignment
	default:
		f[n] = newUint(lookupFibonacci(f, n-1) + lookupFibonacci(f, n-2))

		fibonacciRecInstCount++ // n-1
		fibonacciRecInstCount++ // n-2
		fibonacciRecInstCount++ // n index access
		fibonacciRecInstCount++ // assignment
	}

	defer func() {
		fibonacciRecInstCount++ // n index access
	}()
	return *f[n]
}

func newUint(n uint) *uint { return &n }

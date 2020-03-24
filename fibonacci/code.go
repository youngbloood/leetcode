package fibonacci

func Fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return Fib(N-1) + Fib(N-2)
}

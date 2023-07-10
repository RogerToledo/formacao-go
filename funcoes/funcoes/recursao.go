package funcoes

func Fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fatorial(n-1)
}
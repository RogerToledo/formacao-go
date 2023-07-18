package calculadora

import "errors"

func Soma(numeros []int) int {
	soma := 0

	for _, n := range numeros {
		soma += n
	}

	return soma
}

func Subtracao(numeros []int) int {
	sub := 0

	for i, n := range numeros {
		if i == 0 {
			sub = n
		} else {
			sub -= n
		}
	}

	return sub
}

func Multiplicacao(numeros []int) int {
	mult := 0

	for i, n := range numeros {
		if i == 0 {
			mult = n
		} else {
			mult *= n
		}
	}

	return mult
}

func Divisao(numeros []int) (float64, error) {
	div := 0.0

	for i, n := range numeros {
		if i == 0 {
			div = float64(n)
			continue
		} 

		if n == 0 {
			return 0, errors.New("O divisor não pode ser ZERO!")
		}

		div /= float64(n)
	}

	return div, nil
}
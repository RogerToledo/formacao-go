package funcoes

func Media(notas []float64) float64 {
	var soma float64

	for _, nota := range notas {
		soma += nota
	}

	return soma / float64(len(notas))
}
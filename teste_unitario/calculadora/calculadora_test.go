package calculadora

import (
	"testing"
)

func TestSoma(t *testing.T) {
	testes := []struct {
		descricao string
		input     []int
		esperado  int
	}{
		{
			descricao: "Deve retornar 0 quando todas as entradas são zero",
			input:     []int{0, 0, 0},
			esperado:  0,
		},
		{
			descricao: "Deve retornar 6 quando as entradas forem 1, 2, 3",
			input:     []int{1, 2, 3},
			esperado:  6,
		},
		{
			descricao: "Deve retornar 1 quando as entradas forem 1, 2, 3, -5",
			input:     []int{1, 2, 3, -5},
			esperado:  1,
		},
		{
			descricao: "Deve retornar -1 quando a entrada for -1",
			input:     []int{-1},
			esperado:  -1,
		},
	}

	for _, teste := range testes {
		t.Log(teste.descricao)

		soma := Soma(teste.input)
		if soma != teste.esperado {
			t.Fatalf("Soma: %d -> Esperado: %d", soma, teste.esperado)
		}

	}
}

func TestSubtracao(t *testing.T) {
	testes := []struct {
		descricao string
		input     []int
		esperado  int
	}{
		{
			descricao: "Deve retornar 0 quando todas as entradas são zero",
			input:     []int{0, 0, 0},
			esperado:  0,
		},
		{
			descricao: "Deve retornar -4 quando as entradas forem 1, 2, 3",
			input:     []int{1, 2, 3},
			esperado:  -4,
		},
		{
			descricao: "Deve retornar 5 quando as entradas forem 10, 5",
			input:     []int{10, 5},
			esperado:  5,
		},
		{
			descricao: "Deve retornar -1 quando a entrada for -1",
			input:     []int{-1},
			esperado:  -1,
		},
	}

	for _, teste := range testes {
		t.Log(teste.descricao)

		sub := Subtracao(teste.input)
		if sub != teste.esperado {
			t.Fatalf("Subtração: %d -> Esperado: %d", sub, teste.esperado)
		}

	}
}

func TestMultiplicacao(t *testing.T) {
	testes := []struct {
		descricao string
		input     []int
		esperado  int
	}{
		{
			descricao: "Deve retornar 0 quando todas as entradas são zero",
			input:     []int{0, 0, 0},
			esperado:  0,
		},
		{
			descricao: "Deve retornar 0 quando uma das entradas for 0",
			input:     []int{1, 2, 0},
			esperado:  0,
		},
		{
			descricao: "Deve retornar 50 quando as entradas forem 10, 5, 1",
			input:     []int{10, 5, 1},
			esperado:  50,
		},
		{
			descricao: "Deve retornar -50 quando as entradas forem 10, 5, -1",
			input:     []int{10, 5, -1},
			esperado:  -50,
		},
		{
			descricao: "Deve retornar -1 quando a entrada for -1",
			input:     []int{-1},
			esperado:  -1,
		},
	}

	for _, teste := range testes {
		t.Log(teste.descricao)

		mult := Multiplicacao(teste.input)
		if mult != teste.esperado {
			t.Fatalf("Subtração: %d -> Esperado: %d", mult, teste.esperado)
		}

	}
}

func TestDivisao(t *testing.T) {
	testes := []struct {
		descricao string
		input     []int
		esperado  float64
		erro bool
	}{
		{
			descricao: "Deve retornar 0 quando as entradas são 0, 2",
			input:     []int{0, 2},
			esperado:  0.0,
			erro: false,
		},
		{
			descricao: "Deve retornar erro quando as entradas são 10, 0",
			input:     []int{10, 0},
			esperado:  0.0,
			erro: true,
		},
		{
			descricao: "Deve retornar erro quando as entradas são 10, 5, 0",
			input:     []int{10, 5, 0},
			esperado:  0.0,
			erro: true,
		},
		{
			descricao: "Deve retornar 1 quando as entradas são 10, 5, 2",
			input:     []int{10, 5, 2},
			esperado:  1.0,
			erro: false,
		},
		{
			descricao: "Deve retornar 1 quando as entradas são 10, 5, 2",
			input:     []int{100, 5, 2},
			esperado:  10.0,
			erro: false,
		},
	}

	for _, teste := range testes {
		t.Log(teste.descricao)

		div, erro := Divisao(teste.input)
		if erro != nil && teste.erro == false {
			t.Fatal(erro)
		}

		if div != float64(teste.esperado) {
			t.Fatalf("Divisão: %f -> Esperado: %f", div, teste.esperado)
		}
	}
}

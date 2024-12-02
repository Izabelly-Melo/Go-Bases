package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Exercicio 1")
	fmt.Println(exerc1(140000.00))

	fmt.Println("\nExercicio 2")
	fmt.Println("A média das notas é:", exerc2(-10, 10, 5, 10))

	fmt.Println("\nExercicio 3")
	fmt.Printf("Total do salário %.2f\n", exerc3(60, "C"))

	fmt.Println("\nExercicio 4")
	exerc4()

	fmt.Println("\nExercicio 5")
	exerc5()
}

/*
1 - Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento
de depositar o salário. Para cumprir o objetivo, é necessário criar uma função que retorne o imposto de um salário.
Levando em conta que se a pessoa ganha mais de US$ 50.000, 17% do salário será
deduzido e se a pessoa ganha mais de US$ 150.000, 10% também será deduzido (27% no total).
*/

func exerc1(salario float64) float64 {
	var imposto float64

	if salario > 150000.00 {
		imposto = 0.27 * salario
	} else if salario > 50000.00 {
		imposto = 0.17 * salario
	}

	return imposto
}

/*
2 - Uma escola precisa calcular a média (por aluno) de suas notas. É solicitado que ela
gere uma função na qual possam ser passados N números inteiros e que retorne a média.
Não é possível inserir notas negativas.
*/
func exerc2(nota ...float64) float64 {
	var result float64

	for _, notas := range nota {
		if notas < 0 {
			fmt.Println("Nota inválida!")
			return 0
		} else {
			result += notas
		}
	}

	return result / float64(len(nota))
}

/*
3 - Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria.
Categoria C, o salário deles é de US$ 1.000 por hora.
Categoria B, o salário deles é de US$ 1.500 por hora, mais 20% do salário mensal.
Categoria A, o salário deles é de US$ 3.000 por hora, mais 50% do salário mensal.
Você deve gerar uma função que receba como parâmetro o número de minutos
trabalhados por mês, a categoria e retorne o salário.
*/
func exerc3(minutosTrabalhadas int, categoria string) float64 {
	horasTrabalhadas := minutosTrabalhadas / 60

	var salario float64
	switch categoria {
	case "C":
		salario = float64(1000 * horasTrabalhadas)
	case "B":
		salario = float64(1500 * horasTrabalhadas)
		salario += 0.2 * salario
	case "A":
		salario = float64(3000 * horasTrabalhadas)
		salario += 0.5 * salario
	}
	return salario
}

/*
4 - Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas
de notas para os alunos de um curso. Para isso, eles precisam gerar uma função que
indique o tipo de cálculo que desejam realizar (mínimo, máximo ou médio) e que retorne
outra função e uma mensagem (caso o cálculo não esteja definido) que possa receber
um número N de inteiros e retorne o cálculo indicado na função anterior.
*/
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func exerc4() {
	minFunc, err := operacao(minimum)
	tratandoError(err)

	averageFunc, err := operacao(average)
	tratandoError(err)

	maxFunc, err := operacao(maximum)
	tratandoError(err)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Valor minimo", minValue)
	fmt.Println("Valor medio", averageValue)
	fmt.Println("Valor maximo", maxValue)
}

func operacao(tipo string) (func(...float64) float64, error) {
	switch tipo {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("função inválida")
	}
}

func maxFunc(nota ...float64) float64 {
	var max float64
	for _, valor := range nota {
		if max <= valor {
			max = valor
		}
	}

	return max
}

func minFunc(nota ...float64) float64 {
	var min = nota[0]
	for _, valor := range nota {
		if min >= valor {
			min = valor
		}
	}
	return min
}

func averageFunc(nota ...float64) float64 {
	var media float64

	for _, valor := range nota {
		media += valor
	}

	return media / float64(len(nota))
}

func tratandoError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*
5 - Um abrigo de animais precisa calcular a quantidade de alimentos que precisa
comprar para seus animais de estimação. No momento, eles só têm tarântulas,
hamsters, cães e gatos, mas esperam poder abrigar muito mais animais.
Eles têm os seguintes dados:
1 - Cão 10 kg de comida.
2 - Gato 5 kg de comida.
3 - Hamster 250 g de comida.
4 - Tarântula 150 g de comida.
É solicitado que você:
1- Implemente uma função Animal que receba como parâmetro um valor de texto
com o animal especificado e retorne uma função e uma mensagem (caso o animal não exista).
2 - Uma função para cada animal que calcule a quantidade de comida com base
na quantidade do tipo de animal especificado.
*/

func exerc5() {
	tipoAnimal := "gato"
	qtd := 2

	calculo, err := Animal(tipoAnimal)

	if err != nil {
		fmt.Println(err)
		return
	}

	result := calculo(qtd)

	fmt.Printf("Para %d %s(s), vai precisar de %.2f Kg de comida!\n", qtd, tipoAnimal, result)
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case "cachorro":
		return cachorro, nil
	case "gato":
		return gato, nil
	case "hamster":
		return hamster, nil
	case "tarantula":
		return tarantula, nil
	default:
		return nil, errors.New("Animal não encontrado!")
	}
}

func cachorro(qtdAnimais int) float64 {
	return float64(qtdAnimais) * 10.0
}

func gato(qtdAnimais int) float64 {
	return float64(qtdAnimais) * 5.0
}

func hamster(qtdAnimais int) float64 {
	return float64(qtdAnimais) * 0.250
}

func tarantula(qtdAnimais int) float64 {
	return float64(qtdAnimais) * 0.150
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	exerc5(30, 10)
}

/*
Vamos tornar nosso programa um pouco mais complexo e útil.

Desenvolver as funções necessárias para permitir que a empresa faça cálculos:
Salário mensal de um trabalhador de acordo com o número de horas trabalhadas. A função deve receber as horas
trabalhadas no mês e o valor da hora como argumento.
A função deve retornar mais de um valor (salário calculado e erro). Caso o salário mensal seja igual ou
superior a US$ 150.000, será deduzido um imposto de 10%. Caso o número de horas mensais inserido seja inferior
a 80 ou um número negativo, a função deve retornar um
erro. Ela indicará “Error: the worker cannot have worked less than 80 hours per month”.
*/
type ErroSalary struct {
	mensagem string
}

func (e ErroSalary) Error() string {
	return e.mensagem
}

func validaSalario(salario float64, horas int) (float64, error) {
	var result float64
	var erro error

	if salario >= 150000 {
		result = salario * 0.10
	} else {
		result = 0.0
		erro = errors.New("Salary is not fine!")
	}

	if horas <= 80 {
		result = 0.0
		erro = errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	return result, erro
}

func exerc5(salary float64, horas int) {
	result, erro := validaSalario(salary, horas)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Printf("Salary is ok! %.2f", result)
	}
}

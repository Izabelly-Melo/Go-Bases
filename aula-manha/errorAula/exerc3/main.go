package main

import (
	"errors"
	"fmt"
)

func main() {
	exerc3(1000)
}

/*
Faça o mesmo que no exercício anterior,
mas reformule o código de modo que, em vez de "Error()", seja implementado "errors.New()".
*/

type ErroSalary struct {
	mensagem string
}

func (e ErroSalary) Error() string {
	return e.mensagem
}

func validaSalario(salario int) error {
	if salario <= 10000 {
		return errors.New("Error: the salary entered does not reach the taxable minimum")
	}
	return nil
}

func exerc3(salary int) {
	e := validaSalario(salary)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Salary is ok!")
	}
}

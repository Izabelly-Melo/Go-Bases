package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 1500000
	exerc1(salary)

}

/*
Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".

Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: the salary
entered does not reach the taxable minimum" e acione-a caso "salary" seja
menor que 150.000. Caso contrário, será necessário imprimir no console a mensagem "Must pay tax".
*/

func exerc1(salary int) {
	erro := errorSalary{}
	erro2 := fmt.Errorf("error: %w", erro)
	if salary < 150000 {
		fmt.Println(errors.Unwrap(erro2))
	} else {
		fmt.Println("Must pay tax")
	}
}

type errorSalary struct{}

func (e errorSalary) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

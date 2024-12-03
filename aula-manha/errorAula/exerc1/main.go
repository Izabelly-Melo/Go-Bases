package main

import (
	"fmt"
)

func main() {
	salary := 150000
	exerc1(salary)

}

/*
Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".

Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: the salary
entered does not reach the taxable minimum" e acione-a caso "salary" seja
menor que 150.000. Caso contrário, será necessário imprimir no console a mensagem "Must pay tax".
*/

func exerc1(salary int) {
	erro := valida(salary)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("Must pay tax")
	}

}

type ErrorSalary struct {
	mensagem string
}

func (e ErrorSalary) Error() string {
	return e.mensagem
}

func valida(salario int) error {
	if salario < 150000 {
		return ErrorSalary{"Error: the salary entered does not reach the taxable minimum"}
	}
	return nil
}

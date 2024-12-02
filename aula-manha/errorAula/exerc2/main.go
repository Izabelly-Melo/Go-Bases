package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 10000
	exerc2(salary)
}

/*
Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".
Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: salary
is less than 10000" e a
inicie caso "salary" seja menor ou igual a 10000. A validação deve ser feita com a função "Is()" dentro do "main".
*/

type ErroSalary struct {
	mensagem string
}

func (e ErroSalary) Error() string {
	return e.mensagem
}

func validaSalario(salario int) error {
	if salario <= 10000 {
		return ErroSalary{mensagem: "Error: the salary entered does not reach the taxable minimum"}
	}
	return nil
}

func exerc2(salary int) {
	estrutura := ErroSalary{mensagem: "Error: the salary entered does not reach the taxable minimum"}
	e := validaSalario(salary)

	if e != nil {
		if errors.Is(e, estrutura) {
			fmt.Println(e.Error())
		}
	} else {
		fmt.Println("Salary is ok!")
	}

}

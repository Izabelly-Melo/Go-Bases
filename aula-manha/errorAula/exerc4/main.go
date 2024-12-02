package main

import (
	"fmt"
)

func main() {
	exerc4(1500)
}

/*
Repeti o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de erro receba como
parâmetro o valor de "salary" indicando que ele não atinge o mínimo tributável (a mensagem exibida pelo console
deve dizer:“Error: the minimum taxable amount is 150,000
and the salary entered is: [salary]”, sendo  [salary] o valor do tipo int passado pelo parâmetro).
*/

type ErroSalary struct {
	mensagem string
}

func (e ErroSalary) Error() string {
	return e.mensagem
}

func validaSalario(salario int) error {
	if salario < 150000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salario)
	}
	return nil
}

func exerc4(salary int) {
	e := validaSalario(salary)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Salary is ok!")
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Exercicio 1")
	exerc1()
	fmt.Println("Exercicio 2")
	exerc2()
	fmt.Println("Exercicio 3")
	exerc3()
	fmt.Println("Exercicio 4")
	exerc4()
}

/*
1 - Criar um aplicativo que receba uma variável com a palavra e imprima o número de letras que ela contém.
Em seguida, imprima cada uma das letras.
*/

func exerc1() {
	palavra := "Brasil"

	fmt.Printf("A palavra é: %s e contém %d caracteres\n", palavra, len(palavra))

	for _, caracter := range palavra {
		fmt.Println(string(caracter))
	}
}

/*
2 - Um banco deseja conceder empréstimos a seus clientes, mas nem todos têm
 acesso a eles. Ele tem certas regras para determinar
quais clientes podem receber empréstimos. O banco só concede empréstimos a
clientes com mais de 22 anos de idade, que estejam
empregados e que estejam em seu emprego há mais de um ano. Dentro dos empréstimos
concedidos, o banco não cobrará juros daqueles
que tiverem um salário superior a US$ 100.000.
Você precisa criar um aplicativo que receba essas variáveis e imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos três mensagens diferentes.
*/

func exerc2() {
	var (
		nome            = "Taeyeon"
		salario         = 100000.00
		idade           = 30
		tempoDeTrabalho = 2
	)

	if tempoDeTrabalho > 1 && idade >= 22 {
		if salario >= 100000.00 {
			fmt.Println("Empréstimo aprovado sem juros!")
		} else {
			fmt.Println("Empréstimo aprovado com juros!")
		}

	} else {
		fmt.Println(nome, "Não possui os critérios necessários!")
	}
}

/*
3 - Crie um aplicativo que receba uma variável com o número do mês.
Dependendo do número, imprima o mês correspondente no texto.
Você consegue pensar em maneiras diferentes de resolver isso? Qual delas você escolheria e por quê?
Ex: 7, Julio.
Observação: verifique se o número do mês está correto.
*/

func exerc3() {
	mes := 1
	meses := map[int]string{1: "Jan", 2: "Fev", 3: "Mar", 4: "Abr", 5: "Maio", 6: "Jun", 7: "Jul", 8: "Ago", 9: "Set", 10: "Out", 11: "Nov", 12: "Dez"}

	if mes < 1 || mes > 12 {
		fmt.Println("Este mês não existe!")
	}

	fmt.Println("Número do Mês:", mes)
	fmt.Println("Mês:", meses[mes])
}

/*
4 - Um funcionário de uma empresa quer saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.
Por outro lado, também é necessário:
- Saber quantos de seus funcionários têm mais de 21 anos.
- Adicionar um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Remover Pedro do mapa.
*/
func exerc4() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	fmt.Println("Benjamin tem ", employees["Benjamin"])

	var qtdIdade int

	for _, idade := range employees {
		if idade >= 21 {
			qtdIdade++
		}
	}

	fmt.Println("A quantidade de funcionário que tem mais de 21 anos é", qtdIdade)

	employees["Federico"] = 35
	fmt.Println("Adicionando Federico a lista fica", employees)

	delete(employees, "Pedro")
	fmt.Println("Deletando Pedro a lista fica", employees)
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercicio 1")
	exerc1()
	fmt.Println("\nExercicio 2")
	exerc2()
}

/*
Uma universidade precisa registrar os alunos e gerar uma funcionalidade
para imprimir os detalhes dos dados de cada aluno, como segue:

Nome: [Primeiro nome do aluno]
Sobrenome: [Sobrenome do aluno]
ID: [ID do aluno]
Data: [Data de admissão do aluno]

Os valores entre colchetes devem ser substituídos pelos dados fornecidos
pelos alunos. Para isso, é necessário gerar uma estrutura Students com as variáveis
Name, Surname, DNI, Date e que tenha um método de detalhamento
*/

type Students struct {
	Nome      string
	Sobrenome string
	ID        int
	Data      string
}

func (a *Students) Informacao() {
	fmt.Println("Nome Completo:", a.Nome, a.Sobrenome, "\nID:", a.ID, "\nData:", a.Data)
}

type InfoAluno interface {
	Informacao()
}

func PreencheDados(i InfoAluno) {
	i.Informacao()
}

func exerc1() {
	aluno1 := Students{"Taeyeon", "Kim", 1, "29/11/2024"}
	PreencheDados(&aluno1)
}

/*
Algumas lojas de comércio eletrônico precisam criar uma funcionalidade no Go para gerenciar produtos e retornar o
valor do preço total. A empresa tem três tipos de produtos: Pequeno, Médio e Grande (muitos outros são esperados).
E os custos adicionais são:
Pequeno: apenas o custo do produto
Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500.
O custo de manter o produto em estoque na loja é uma porcentagem do preço do produto.
É necessária uma função factory que receba o tipo de produto e o preço e retorne uma interface Product que tenha o método Price.
Deve ser possível executar o método Price e fazer com que o método retorne o preço total com base no custo do produto
e em quaisquer custos adicionais.
*/
const (
	Grande  = "Grande"
	Medio   = "Medio"
	Pequeno = "Pequeno"
)

type IProduto interface {
	Valor() float64
}

type Produto struct {
	tipo  string
	price float64
}

func (p *Produto) Valor() float64 {
	switch p.tipo {
	case Grande:
		return p.price + (p.price * 0.06) + 2500
	case Medio:
		taxa := 0.03
		var custo = p.price * taxa
		return p.price + (custo * 2)
	case Pequeno:
		return p.price
	}
	return 0.0
}

func NovoProduto(tipo string, price float64) IProduto {
	return &Produto{
		tipo:  tipo,
		price: price,
	}
}

func ExibePrice(p IProduto) {
	fmt.Printf("Valor Total: %.2f\n", p.Valor())
}

func exerc2() {
	mac := NovoProduto(Grande, 22000)
	lapis := NovoProduto(Pequeno, 1.99)
	caderno := NovoProduto(Medio, 20)

	fmt.Println("MacBook:")
	ExibePrice(mac)
	fmt.Println("\nLápis:")
	ExibePrice(lapis)
	fmt.Println("\nCaderno:")
	ExibePrice(caderno)
}

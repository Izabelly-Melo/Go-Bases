package main

import "fmt"

func main() {
	fmt.Println("Exercicio 1")
	exerc1()
	fmt.Println("\nExercicio 2")
	exerc2()
}

/*
1 - Crie um programa que atenda aos seguintes pontos:
- Ter uma estrutura chamada Product com os campos ID, Name, Price, Description e Category.
- Ter uma slice de Produto chamada Produtos instanciada com valores. 2 métodos
associados à estrutura Produto: Save(), GetAll(). O método Save() deve pegar a fatia de Products e
adicionar o produto a partir do qual o método é chamado. O método GetAll() imprime todos os
produtos salvos na fatia Products.
- Uma função getById() para a qual um INT deve ser passado como parâmetro e retorna o produto
correspondente ao parâmetro passado.
- Execute pelo menos uma vez cada método e função definidos em main().
*/
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Produtos = []Product{
	{
		ID:          1,
		Name:        "A linguagem de programação",
		Price:       100.00,
		Description: "Tudo sobre GO do zero",
		Category:    "Livros",
	},
	{
		ID:          2,
		Name:        "Macbook",
		Price:       250000.00,
		Description: "Macbook pro M3",
		Category:    "Eletronicos",
	},
}

func (p *Product) Save() {
	Produtos = append(Produtos, *p)
}

func (p *Product) getAll() {
	fmt.Println("\nLista de Produtos:")
	for _, produtos := range Produtos {
		fmt.Println("\nID:", produtos.ID)
		fmt.Println("Name:", produtos.Name)
		fmt.Println("Price:", produtos.Price)
		fmt.Println("Description:", produtos.Description)
		fmt.Println("Category:", produtos.Category)
	}

}

func getById(id int) *Product {
	for _, prod := range Produtos {
		if prod.ID == id {
			return &prod
		}
	}
	return nil
}

func exerc1() {
	album := Product{
		ID:          3,
		Name:        "My Voice - Taeyeon",
		Price:       250.00,
		Description: "Full Album",
		Category:    "Música",
	}

	celular := Product{
		ID:          4,
		Name:        "Samsung galaxy 24",
		Price:       6000.00,
		Description: "Celular",
		Category:    "Eletronicos",
	}

	album.Save()
	fmt.Printf("O Produto %s foi salvo!\n", album.Name)

	celular.Save()
	fmt.Printf("O Produto %s foi salvo!\n", celular.Name)

	album.getAll()

	existeProduto := getById(3)

	if existeProduto != nil {
		fmt.Println("\nO Produto escolhido")
		fmt.Println(*existeProduto)
	} else {
		fmt.Println("\nProduto não cadastrado!")
	}
}

/*
2 - Uma empresa precisa fazer um bom gerenciamento de seus funcionários. Para isso, criaremos um pequeno programa
que nos ajudará a gerenciar corretamente esses funcionários. Os objetivos são:
- Criar uma estrutura Person com os campos ID, Name, DateOfBirth.
- Criar uma estrutura Employee com os campos: ID, Position e uma composição com a estrutura Person.
- Criar um método para a estrutura Employee chamado PrintEmployee(), que imprimirá os campos de um funcionário.
- Instanciar na função main() uma Person e um Employee carregando seus respectivos campos e, finalmente, executar o método PrintEmployee().
- Se você conseguir realizar esse pequeno programa, poderá ajudar a empresa a resolver o problema de gerenciamento dos funcionários.
*/
type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Println("\nID:", e.ID)
	fmt.Println("Position:", e.Position)
	fmt.Println("Person ID:", e.Person.ID)
	fmt.Println("Name:", e.Person.Name)
	fmt.Println("Birthday:", e.Person.DateOfBirth)
}

func exerc2() {
	techLeader := Employee{}
	techLeader.ID = 10
	techLeader.Position = "Tech Leader"
	techLeader.Person = Person{ID: 10, Name: "Im Nayeon", DateOfBirth: "22/09/1995"}

	dev := Employee{}
	dev.ID = 22
	dev.Position = "Developer"
	dev.Person = Person{ID: 1, Name: "Momo Hirai", DateOfBirth: "09/11/1996"}

	techLeader.PrintEmployee()
	dev.PrintEmployee()
}

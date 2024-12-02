package main

import "fmt"

func main() {
	fmt.Println("Exercício 1")
	exerc1()
	fmt.Println("\nExercício 2")
	exerc2()
	fmt.Println("\nExercício 3 está comentado")
	fmt.Println("\nExercício 4 está comentado")
}

/*
Crie um aplicativo em que você tenha como variáveis seu nome e endereço.
Imprimir no console o valor de cada uma das variáveis.
*/
func exerc1() {
	var (
		nome     = "Izabelly"
		endereco = "São Paulo - SP"
	)
	fmt.Println("Meu nome é", nome)
	fmt.Println("Moro em", endereco)
}

/*
Uma empresa de meteorologia deseja um aplicativo em que possa ter a temperatura, a umidade e a pressão atmosférica de diferentes locais.

Declare 3 variáveis especificando o tipo de dados; como valor, elas devem ter a temperatura, a umidade e a pressão do local onde você está.
Imprima os valores das variáveis no console.
Que tipo de dados você atribuiria às variáveis?
R: atribuiria o valor de int para as variáveis
*/
func exerc2() {
	var (
		temperatura int    = 33
		umidade     int    = 34
		pressao     int    = 1008
		local       string = "São Paulo"
	)

	fmt.Println("Local:", local)
	fmt.Println("Temperatura:", temperatura)
	fmt.Println("Pressão:", pressao)
	fmt.Println("Umidade:", umidade)
}

/*
Um professor de programação está corrigindo os exames de seus alunos na disciplina Programação I para dar a eles os
retornos correspondentes. Um dos itens do exame é declarar diferentes variáveis.

Preciso de ajuda para:
Detecte quais dessas variáveis declaradas pelo aluno estão corretas.
Corrija as incorretas.

RESPOSTA:
- var 1firstName string: está errada, pois começa com número o certo: var firstName string
- var lastName string está correta
- var int age está errada, pois a variável está esperando um valor numérico o certo: var int 2
- 1lastName := 6 está errada, pois começa com número o certo: lastName := 6
- var driver_license = true está certo, mas não é uma boa prática o melhor seria driverLicense
- var person height int está errado o certo: var person, height int
- childsNumber := 2 está correto
*/

/*
Um estudante de programação tentou fazer declarações de variáveis com seus respectivos tipos em Go,
mas teve várias dúvidas ao fazer isso. A partir disso, ele nos forneceu seu código e pediu a ajuda de um desenvolvedor experiente que pudesse:

Verifique seu código e faça as correções necessárias.
RESPOSTA:

var lastName string = "Smith"
var age int = 35
ok := false
var salary float32 = 45857.90
var firstName string = "Mary"
*/

//------------------------------------------------------ EXERCICIOS DO GIT --------------------------------------------------//

/*
1. O que é um pacote em Golang e qual a sua importância em um programa?
resposta: um pacote é onde fica o código deve ser sempre declarado no começo, então é muito importante
*/

/*
2. Como declarar e utilizar o pacote principal (main) em uma aplicação Go?
package main

		import "fmt"

		func main() {
			fmt.Println("hi")
		}
*/

/*
3. Qual é a diferença entre pacotes internos (standard library) e pacotes personalizados em Go?
resposta: Os pacotes internos ja vêm instalados com a linguagem GO e os pacotes personalizados são pacotes criados por
desenvolvedores para atender necessidades específicas
*/

/*
4. Como importar e utilizar um pacote externo no Go?
resposta:
		import "github.com/user/teste"
	ou
		import (
			"fmt""
			"github.com/user/teste"
		)
*/

/*
5. Qual é o comando utilizado no Go para inicializar um novo módulo e gerenciar dependências?
resposta: para inicializar um módulo utilizar o comando:
	go mod init github.com/idmelo_meli/aulas-go.git ele ira criar um arquivo go.mod
*/

/*
6. O que acontece se você importar um pacote no Go e não utilizá-lo no código?
resposta: não pode, porque no GOLang exibe o erro na hora de compilar, então, não é possível adicionar a importação ou a variável e não utilizá-las
*/

/*
7. Como declarar uma variável global em um pacote e torná-la acessível a outros pacotes?
resposta: a variável deve começar sempre começar com a letra maiúscula para ser pública e assim poder utilizá-la em outros pacotes
*/

/*
8. Qual a diferença entre uma variável declarada com `var` e uma variável inicializada com `:=` em Go?
resposta: var pode ser usado em pacotes e funções, já o := pode ser usado apenas em funções
*/

/*
9. Como você exporta uma variável ou função de um pacote em Golang, e qual a convenção de nomenclatura para isso?
resposta:
para importar: import "github.com/user/teste"
e para usar a função importada deve ser: teste.OiMundo()
*/

/*
10.Explique o conceito de shadowing (sombras) em relação às variáveis no Go e como isso pode afetar o escopo.
resposta: shadowing é quando se é criado uma variável no escopo interno ( por meio de {} ) com o mesmo nome da variável do
escopo externo, dessa forma, mudando o conteúdo da variável externa por um breve momento, não é muito
recomendado pois pode causar confusão
*/

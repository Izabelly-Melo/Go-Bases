package main

import (
	"fmt"
	"os"
)

func main() {
	exerc1()
}

/*
Uma empresa de contabilidade precisa ter acesso aos dados de seus funcionários para poder
fazer várias liquidações. Para isso, eles têm todos os detalhes necessários em um arquivo. .txt.

Você terá que desenvolver a funcionalidade para poder ler o arquivo .txt indicado pelo cliente,
mas ele não passou o arquivo para ser lido pelo nosso programa.
Desenvolva o código necessário para ler os dados do arquivo chamado "customers.txt" (lembre-se do que
vimos sobre o pkg "os"). Como não temos o arquivo necessário, receberemos um erro e, nesse caso, o
programa deve gerar um pânico ao tentar ler um arquivo que não existe, exibindo a seguinte
mensagem “The indicated file was not found or is damaged”.

Apesar disso, a mensagem "execução concluída" sempre será impressa no console.
*/

func exerc1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, erro := os.ReadFile("customers.txt")

	if erro != nil {
		panic("The indicated file was not found or is damaged")
	}

}

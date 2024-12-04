package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {

	file, err := os.Create("alunos.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("=== Menu de Opções ===")
		fmt.Println("1 - Incluir Aluno")
		fmt.Println("2 - Listar Todos os Alunos")
		fmt.Println("3 - Pesquisar Aluno por Matrícula")
		fmt.Println("4 - Alterar Aluno")
		fmt.Println("5 - Excluir Aluno")
		fmt.Println("6 - Sair")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Escolha uma opção:")
		input, _ := reader.ReadString('\n')

		escolha, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			fmt.Println("Erro: Insira número válido!")
			continue
		}

		switch escolha {
		case 1:
			matricula := uuid.New().String()
			nome := cadastra(scanner, "Nome: ")
			telefone := cadastra(scanner, "Telefone: ")
			email := cadastra(scanner, "Email: ")

			_, err = file.WriteString("==Aluno==\n\nMatrícula: " + matricula +
				"\nNome: " + nome +
				"\nTelefona: " + telefone + "\nEmail: " + email)
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo:", err)
				return
			}
			fmt.Println("Aluno cadastrado com sucesso!\n")
		case 2:
			data, err := os.ReadFile("alunos.txt")
			if err != nil {
				fmt.Println("Erro ao ler arquivo!")
				return
			}
			println(string(data))
		case 3:
			/*
				var pesquisa string
				fmt.Print("Digite a Matrícula:")
				fmt.Scanln(&pesquisa)

				file, err = os.Open("alunos.txt")
				scannerLeitura := bufio.NewScanner(file)
				existe = false

				for scannerLeitura.Scan() {
					if strings.Contains(scannerLeitura.Text(), pesquisa) {
						fmt.Println("Encontro: ", scannerLeitura.Text())
						existe = true
					}
				}

				if !existe {
					fmt.Println("Aluno não existe!")
				}
			*/
		case 4:
			//altera
		case 5:
			//excluir
		case 6:
			break
		}

		if escolha == 6 {
			break
		}
	}
	defer file.Close()

}

/*
Com base no que aprendemos até agora, vamos criar um programa em Go para o cadastro de alunos.
Este programa terá os seguintes campos:
Matrícula (UUID)
Nome (string)
Telefone (string)
Email (string)

=== Menu de Opções ===
1 - Incluir Aluno
2 - Listar Todos os Alunos
3 - Pesquisar Aluno por Matrícula
4 - Alterar Aluno
5 - Excluir Aluno
6 - Sair

Matricula: uuid.New().String()

obs:
Cuidados com Entrada de Espaços
fmt.Scanf() não lida bem com entradas que contêm espaços em strings,
como "João Silva". Se houver necessidade de lidar com entradas que contêm espaços,
considere o uso de bufio.Scanner ou use fmt.Scanln em vez de fmt.Scanf().
*/

func cadastra(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)     // Imprime o prompt para o usuário
	scanner.Scan()        // Lê a entrada do usuário
	return scanner.Text() // Retorna o texto digitado
}

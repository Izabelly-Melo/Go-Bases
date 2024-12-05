package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Aluno struct {
	Matricula string
	Nome      string
	Telefone  string
	Email     string
}

func main() {

	file, err := os.Create("alunos.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	var response int
	for {

		fmt.Println("=== Menu de Opções ===")
		fmt.Println("1 - Incluir Aluno")
		fmt.Println("2 - Listar Todos os Alunos")
		fmt.Println("3 - Pesquisar Aluno por Matrícula")
		fmt.Println("4 - Alterar Aluno")
		fmt.Println("5 - Excluir Aluno")
		fmt.Println("6 - Sair")

		fmt.Print("Escolha uma opção:")
		fmt.Scanln(&response)

		var alunos []Aluno
		switch response {
		case 1:
			fmt.Println("Nome: ")
			var nome string
			fmt.Scanln(&nome)

			fmt.Println("Telefone: ")
			var telefone string
			fmt.Scanln(&telefone)

			fmt.Println("Email: ")
			var email string
			fmt.Scanln(&email)

			var aluno = Aluno{
				Matricula: uuid.New().String(),
				Nome:      nome,
				Telefone:  telefone,
				Email:     email,
			}

			alunos = append(alunos, aluno)

			_, err = file.WriteString("Matricula: " + alunos[0].Matricula + ", Nome: " + alunos[0].Nome + ", Telefone:" + alunos[0].Telefone + ", Email: " + alunos[0].Email + "\n")
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo", err)
				return
			}
			fmt.Println("Aluno cadastrado com sucesso!\n")

		case 2:
			println(listarAlunos())
		case 3:
			fmt.Println("Matricula: ")
			var matricula string
			fmt.Scanln(&matricula)

			data, err := os.ReadFile("alunos.txt")
			if err != nil {
				fmt.Println("Erro ao ler arquivo!")
				return
			}

			alunos := strings.Split(string(data), "\n")
			existe := false
			for _, aluno := range alunos {
				if strings.Contains(aluno, matricula) {
					fmt.Println("\n===Aluno===")
					fmt.Println(aluno + "\n")
					existe = true
					break
				}
			}

			if !existe {
				fmt.Println("Aluno não encontrado.")
			}

		case 4:
			fmt.Println("Matricula: ")
			var matricula string
			fmt.Scanln(&matricula)

			lista := listarAlunos()
			linhas := strings.Split(lista, ",")
			var todosAlunos []Aluno

			for _, aluno := range linhas {
				alunos := strings.Split(aluno, ",")

				if alunos[0] == matricula {
					fmt.Println("Nome: ")
					var nome string
					fmt.Scanln(&nome)

					fmt.Println("Telefone: ")
					var telefone string
					fmt.Scanln(&telefone)

					fmt.Println("Email: ")
					var email string
					fmt.Scanln(&email)
					student := Aluno{
						Matricula: alunos[0],
						Nome:      nome,
						Telefone:  telefone,
						Email:     email,
					}

					todosAlunos = append(todosAlunos, student)
				} else {
					if len(alunos) == 4 {
						student := Aluno{
							Matricula: alunos[0],
							Nome:      alunos[1],
							Telefone:  alunos[2],
							Email:     alunos[3],
						}
						todosAlunos = append(todosAlunos, student)
					}

				}

			}
		case 5:

		case 6:
			break
		default:
		}

		if response == 6 {
			break
		}
	}

}

func listarAlunos() string {
	data, err := os.ReadFile("alunos.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo!")
		return ""
	}
	return string(data)
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

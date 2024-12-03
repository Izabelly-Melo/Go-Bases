package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	exerc2("customer.txt")
}

/*
Em seguida, vamos criar um arquivo "customers.txt" com informações sobre os clientes do estúdio.
Agora que o arquivo existe, o pânico não deve ser acionado.

Criamos o arquivo "customers.txt" e adicionamos as informações do cliente.
Estendemos o código do ponto um para que possamos ler esse arquivo e imprimir os
dados que ele contém. Caso não seja possível ler o arquivo, um "panic" deve ser iniciado.
Lembre-se de que sempre que a execução terminar, independentemente do resultado,
devemos ter um "defer" para indicar que a execução foi concluída. Lembre-se também
de fechar os arquivos ao final de seu uso.

exemplo:
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
*/

func exerc2(arquivo string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	dir, _ := os.Getwd()

	filename := filepath.Join(dir, arquivo)
	file, err := os.Open(filename)
	if err != nil {
		panic("Erro ao abrir arquivo\n" + err.Error())
	}

	result, erro := os.ReadFile(arquivo)

	if erro != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println(string(result))

	defer file.Close()
}

// ● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
// da cobrança das despesas restantes. Desenvolva e implemente uma função para
// gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
// Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
// interrompa a execução e aborte

// ● Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
// fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
// código, implemente a função para ler um arquivo chamado “customers.txt” (como no
// exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
// retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
// agora. Esse erro deve:

// 1. Gerar um panic;
// 2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
// está danificado”, e continuar com a execução do programa normalmente.

// Requisitos gerais:
// ● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
// ● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
// um valor de erro (por exemplo, aqueles que tentam ler arquivos).
// Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
// que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
// o caso de erro retornado)

package main

import (
	"fmt"
	"os"

	guuid "github.com/google/uuid"
)

type Cliente struct {
	arquivo   string
	nome      string
	sobrenome string
	rg        string
	telefone  string
	endereco  string
}

func GerarID() string {
	newId := guuid.New().String()
	if newId == "" {
		panic("Erro ao gerar um novo ID.")
	}
	return newId
}

func VerificarSeNomeJaExiste(nome string) bool {
	LerArquivo()
	return false
}

func LerArquivo() {
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado.")
	}
}

func handlePanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

func main() {

	handlePanic()

	if !VerificarSeNomeJaExiste("Heymer") {

		cliente := Cliente{
			arquivo:   GerarID(),
			nome:      "Heymer",
			sobrenome: "Britto",
			rg:        "123456",
			telefone:  "1732555555",
			endereco:  "não te interessa",
		}

		fmt.Printf("Nome do cliente: %s \n ", cliente.nome)
	}
}

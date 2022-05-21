package main

import "fmt"

type loja struct {
	listProdutos []produtos
}

type produtos struct {
	tipo      string
	nome      string
	preco     float64
	adicional float64
}

type Produtos interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar()
}

//Contrato do apartamento que não existe ainda, é da jéssica
//Contrato de código (interface), garante a execução de algo que ainda não existe

//Somente a estrutura da LOJA conhece o total
func (l *loja) Total() float64 {
	var total float64

	for _, produto := range l.listProdutos {
		total += produto.preco + produto.adicional
	}

	return total
}

func (l *loja) Adicionar(p produtos) {
	l.listProdutos = append(l.listProdutos, p)
}

func NovoProduto(tipo, nome string, preco float64) *produtos {
	return &produtos{
		tipo:  tipo,
		nome:  nome,
		preco: preco,
	}
}

func NovaLoja(listProdutos ...produtos) *loja {
	return &loja{listProdutos: listProdutos}
}

func main() {

	//* -> 100
	listaProdutos := []produtos{}

	//* -> 101
	listaProdutos = append(listaProdutos, *NovoProduto("Pequeno", "Celular", 1000.00))
	//* -> 102
	listaProdutos = append(listaProdutos, *NovoProduto("Grande", "Sofa", 5000.00))

	//* -> 200
	//* -> -> 100
	//* -> -> -> 101
	//* -> -> -> 102
	loja2 := *NovaLoja(listaProdutos...)

	fmt.Printf("Valor total produtos na Loja: R$%.2f\n", loja2.Total())

	loja2.Adicionar(*NovoProduto("Medio", "TV", 2000.00))

	fmt.Printf("Valor total produtos na Loja: R$%.2f\n", loja2.Total())
}

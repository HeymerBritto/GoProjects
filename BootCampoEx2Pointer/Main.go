// Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
// produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
// mesmo endereço de memória no main do programa e nas funções.

//estrutura = struct
//array sempre append (incluir) delete (deletar)

// - Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).

// - Produto: Nome, preço, quantidade.
// Algumas funções necessárias:
// - Novo produto: recebe nome e preço, e retorna um produto.
// - Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
// - Deletar produtos: recebe um usuário, apaga os produtos do usuário.

package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

func NovoProduto(nome string, preco float64) *produto {
	return &produto{
		nome:  nome,
		preco: preco,
	}
}

func Adicionar(u *usuario, p *produto, q int) {
	p.quantidade += q
	u.produtos = append(u.produtos, *p)
}

//recebe (nos parenteses)
func Deletar(u usuario) {
	u.produtos = nil
}

func main() {
	user := usuario{
		nome:      "Manu",
		sobrenome: "Moreira",
		email:     "manu@gmail.com",
	}
	Adicionar(&user, NovoProduto("Cama", 1000.00), 1)

	fmt.Printf("Produto: %s", user.produtos[0].nome)
}

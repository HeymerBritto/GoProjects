/*
Exercício 2 - Olá {nome}

1. Crie dentro da pasta go-web um arquivo chamado main.go
2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
“mensagem” e diga Olá seguido do seu nome.
3. Acesse o end-point para verificar se a resposta está correta.

Exercício 3 - Listar Entidade

Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista
do tema escolhido.
1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
correspondentes.
2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
3. Crie uma handler para o endpoint chamada "GetAll".
4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id          int32   `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int32   `json:"idade"`
	Altura      float32 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"dataCriacao"`
}

func GetAll() []User {

	jsonFile, _ := os.Open("users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []User
	json.Unmarshal([]byte(byteValue), &users)

	defer jsonFile.Close()

	// if user.nome == "heymer"
	// return user

	return users
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem": "Olá Heymer",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, GetAll())
	})

	r.Run() // listen and serve on 0.0.0.0:8080

}

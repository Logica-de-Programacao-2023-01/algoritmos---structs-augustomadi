// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra
//	struct com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa como parâmetro e
//	imprima seu endereço completo.
package main

import (
	"fmt"
	"strconv"
)

type dados struct {
	rua    string
	número int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereço dados
}

func adress(p Pessoa) string {
	return p.endereço.rua + " rua " + strconv.Itoa(p.endereço.número) + " " + p.endereço.cidade + " " + p.endereço.estado
}

func main() {
	pessoa1 := Pessoa{
		nome:  "guilherme",
		idade: 18,
		endereço: dados{
			rua:    "cruzeiro",
			número: 44,
			cidade: "brasilia",
			estado: "DF",
		},
	}
	fmt.Printf("Endereço de %s é: %s", pessoa1.nome, adress(pessoa1))
}

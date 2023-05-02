// Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e calcule
//	a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.

package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func area_circ(c1 Circulo) float64 {
	return math.Pi * c1.raio * c1.raio
}

func main() {
	c1 := Circulo{raio: 5.0}
	area1 := area_circ(c1)
	fmt.Print("A area do circulo é: ", area1)
}

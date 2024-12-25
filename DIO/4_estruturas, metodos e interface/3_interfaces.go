package main

import (
	"fmt"
	"math"
)

// interface usada para formas geometricas
type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func main() {
	q := quadrado{lado: 10}
	c := circulo{raio: 5}

	medir(q)
	medir(c)
}

func medir(g geometria) {
	fmt.Println("g:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perimetro:", g.perimetro())
}

//

func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

//as func main e medir tbm poderiam estar aqui embaixo

package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}
type pez struct{}
type pajaro struct{}

func (perro) mover() string {
	return "Soy un perro y puedo caminar"
}

func (pajaro) mover() string {
	return "Soy un pajaro y puedo volar"
}

func (pez) mover() string {
	return "Soy un pez y puedo nadar"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {

	p := perro{}
	moverAnimal(p)
	pa := pajaro{}
	moverAnimal(pa)
	pe := pez{}
	moverAnimal(pe)

}

package main

import "fmt"

type animal interface {
	whatType() string
	howDoesScream() string
}

type dog struct {
	name string
}

type lion struct {
	jungle string
}

func (d dog) whatType() string {
	return "Bark"
}

func (l lion) whatType() string {
	return "Roar"
}

func (d dog) howDoesScream() string {
	return "domestic"
}

func (l lion) howDoesScream() string {
	return "wild"
}

func describe(a animal) {
	fmt.Println("=================")
	fmt.Println(a)
	fmt.Println(a.howDoesScream())
	fmt.Println(a.whatType())
}

func main() {
	d := dog{name: "Rex"}
	l := lion{jungle: "Amazonia"}
	describe(d)
	describe(l)
	fmt.Println("=================")
}

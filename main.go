package main

import "fmt"

func main() {
	OlaMundo()
	Variaveis()
	Tipos()
	Concatenar()
	fmt.Println(Conversão(1.45))
}

func OlaMundo() {
	fmt.Println("Ola Mundo Cruel!")
}

func Variaveis() {
	nome, idade, altura := "Carlos", 45, 1.78
	fmt.Printf("Meu nome e %v, minha idade é %v anos, tenho %3.2f de altura.\n", nome, idade, altura)
}

func Tipos() {
	nome, idade, altura := "Carlos", 45, 1.78
	fmt.Printf("Variável Nome\tTipo: %T\tValor: %v\n", nome, nome)
	fmt.Printf("Variável Idade\tTipo: %T\tValor: %v\n", idade, idade)
	fmt.Printf("Variável Altura\tTipo: %T\tValor: %v\n", altura, altura)
}

func Concatenar() {
	nome, sobrenome := "Carlos Alberto de", "Lima Campos"

	fmt.Println(nome + " " + sobrenome)
}

func Conversão(num float64) int {
	return int(num)
}

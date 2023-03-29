package main

import (
	"fmt"
)

// Definindo a struct "Pessoa"
type Pessoa struct {
	nome  string
	idade int
}

// Definindo a struct "Estudante", que usa a composição para "herdar" as propriedades de "Pessoa"
type Estudante struct {
	Pessoa
	notaFinal float64
}

// Adicionando uma função à struct "Pessoa" que retorna a idade da pessoa
func (p Pessoa) obterIdade() int {
	return p.idade
}

// Adicionando uma função à struct "Pessoa" que retorna o nome da pessoa
func (p Pessoa) obterNome() string {
	return p.nome
}

func main() {
	// Criando uma instância de "Estudante"
	estudante := Estudante{Pessoa{"João", 18}, 8.5}

	// Acessando as funções da struct "Pessoa" a partir da instância de "Estudante"
	fmt.Println("Nome:", estudante.obterNome())
	fmt.Println("Idade:", estudante.obterIdade()) //é aqui que rola a herança, reutilização do método.

	// Acessando as propriedades de "Estudante"
	fmt.Println("Nota final:", estudante.notaFinal)
}

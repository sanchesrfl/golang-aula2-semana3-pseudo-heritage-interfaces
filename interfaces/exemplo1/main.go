package main

import "fmt"

// Define a interface "Animal" que especifica um método "MakeSound"
type Animal interface {
	MakeSound() string
}

// Define a struct "Cat" que implementa a interface "Animal"
type Cat struct {
	Name string
}

// Implementa o método "MakeSound" da interface "Animal" para a struct "Cat"
func (c Cat) MakeSound() string {
	return "meowh!"
}

// Define a struct "Dog" que implementa a interface "Animal"
type Dog struct {
	Name string
}

// Implementa o método "MakeSound" da interface "Animal" para a struct "Dog"
func (d Dog) MakeSound() string {
	return "Hof hof!"
}

func main() {
	// Cria uma variável "a" do tipo "Animal" que pode ser tanto um gato quanto um cachorro
	var a Animal

	// Cria uma instância da struct "Cat" e atribui a variável "a"
	a = Cat{Name: "Tikito"}
	fmt.Println(a.MakeSound())

	// Cria uma instância da struct "Dog" e atribui a variável "a"
	a = Dog{Name: "Mamão"}
	fmt.Println(a.MakeSound())
}

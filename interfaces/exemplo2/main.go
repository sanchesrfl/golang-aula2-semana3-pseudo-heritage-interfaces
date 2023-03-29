package main

import "fmt"

// Define a interface que especifica um método "GetPrice"
type Item interface {
	GetPrice() float64
}

// Define uma struct que representa um hambúrguer
type Hamburguer struct {
	Name  string
	Price float64
}

// Implementa o método "GetPrice" da interface "Item" para a struct "Hamburguer"
func (b Hamburguer) GetPrice() float64 {
	return b.Price
}

// Define uma struct que representa batatas fritas
type Batata struct {
	Size  string
	Price float64
}

// Implementa o método "GetPrice" da interface "Item" para a struct "Batata"
func (f Batata) GetPrice() float64 {
	return f.Price
}

// Define uma struct que representa um refrigerante
type Drink struct {
	Name  string
	Price float64
}

// Implementa o método "GetPrice" da interface "Item" para a struct "Drink"
func (d Drink) GetPrice() float64 {
	return d.Price
}

// Define uma função que recebe um slice de itens do menu e retorna o preço total
func CalculateTotal(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += item.GetPrice()
	}
	return total
}

func main() {
	// Cria uma lista de itens do menu
	items := []Item{
		Hamburguer{Name: "X-Salada", Price: 25},
		Hamburguer{Name: "X-Bacon", Price: 38},
		Batata{Size: "Grande", Price: 18},
		Drink{Name: "Dolly", Price: 3},
	}

	// Calcula o preço total dos itens do menu
	fmt.Println("Preço Total do Pedido: ", CalculateTotal(items))
}

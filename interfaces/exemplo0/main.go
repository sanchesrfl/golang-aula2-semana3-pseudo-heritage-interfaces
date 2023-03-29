//de modo geral podemos dizer que interfaces são conjuntos
//de métodos, ou de assinaturas de métodos.

//exemplo básico
package main

import "fmt"

//declaração interface
type Contador interface {
	Adicionar(i int)
	Valor() int
}

//implementação basica das interfaces

//declaramos uma struct
type Status struct {
	valor int
}

//criamos métodos que vinculam-se a struct e implementam
// as assinaturas (métodos) da interface.
func (s *Status) Adicionar(i int) {
	s.valor += i
}

//implementação segundo método da interface Contador
func (s *Status) Valor() int {
	return s.valor
}

//função que recebe dado do tipo da interface fica "lockada" na interface
//isto, é só consegue chamar método da interface em seu ambiente de execução interno
func StartAnalytics(c Contador, i int) {
	c.Adicionar(i)
	print(c.Valor())
}

func main() {

	//iniciamos struct Status
	stats := Status{1}

	//checagem valor status
	fmt.Println(stats)
	//execução função da interface
	StartAnalytics(&stats, 2)

}

//tambem podemos utilizar interfaces vazias
//é uma maneira de passarmos um argumento sem saber seu tipo
//no entanto, não é recomendável escrever código Go assim!
//exemplo de função utilizando interface vazia.

func AlgumaCoisa(i interface{}) string {
	if i == 0 {
		return "zero"
	}
	if i == "" {
		return "empty string"
	}
	return "something else"
}

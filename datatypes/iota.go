//Conceitos de Iota e Expressoes Constantes

package datatypes

import (
	"fmt"
)

//const pi = 3.1415 //Esta constante pode ser acessada em todo o pacote

const (
	first  = 1
	second = "second"
	//Third  = Iota
	//Fourth = Iota
	//Fifth  = Iota
)

func main_Iota() {

	fmt.Println(first, second)
	//p1 = 1.2 //Irá dar erro devido a ser uma constante

	//const c int = 3 Apenas para inteiros (Valores com float nao retornava)
	const c = 3
	fmt.Println(c + 3)

	// a bunch of code

	//fmt.Println(float32(c) + 1.2) - Faz a conversao de inteiro para float
	fmt.Println(c + 1.2)

}

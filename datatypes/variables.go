//Declaracao de variaveis com as suas primitivas
package datatypes

import (
	"fmt"
)

func variaveis() {
	var i int
	i = 42
	fmt.Println(i)

	/* FLOAT: só existe 2 tipos
	FLOAT32 - Ponto flutuante 32bits
	FLOAT64 - Ponto flutuante 64bits
	*/
	var f float32 = 3.14
	fmt.Println(f)

	//Uma forma diferente de Declarar as variaveis
	firstName := "Arthur"
	fmt.Println(firstName)

	//Booleanos
	b := true
	fmt.Println(b)

	//Funcao Complex
	c := complex(3, 4)
	fmt.Println(c)

	//Multiplas declarações numa so linha
	r, im := real(c), imag(c)
	fmt.Println(r, im)

}

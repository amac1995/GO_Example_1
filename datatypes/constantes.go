/*
Constantes - Funcionam da mesma forma que as variaveis,
a diferenca é que as constantes nao podem alterar
o seu valor e as variaveis podem.
*/
package datatypes

import (
	"fmt"
)

func main_Constantes() {
	const pi = 3.1415
	fmt.Println(pi)
	//p1 = 1.2 //Irá dar erro devido a ser uma constante

	//const c int = 3 Apenas para inteiros (Valores com float nao retornava)
	const c = 3
	fmt.Println(c + 3)

	// a bunch of code

	//fmt.Println(float32(c) + 1.2) - Faz a conversao de inteiro para float
	fmt.Println(c + 1.2)

}

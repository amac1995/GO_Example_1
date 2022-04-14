//Apontadores
package datatypes

import (
	"fmt"
)

func main_Apontadores() {
	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	//firstName := "Tricia"
	fmt.Println(ptr, *ptr)
}

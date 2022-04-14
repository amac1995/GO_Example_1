//Maps

package Collections

import (
	"fmt"
)

func main_maps() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	//Editar valores da lista
	m["foo"] = 27
	fmt.Println(m)
	//Eliminar valores da lista
	delete(m, "foo")
	fmt.Println(m)

}

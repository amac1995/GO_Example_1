//Slices

package collections

import (
	"fmt"
)

func main_slices() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	//Os 2 pontos significa do inicio ao fim da matriz
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

}

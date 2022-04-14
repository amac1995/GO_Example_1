//Arrays

package collections

import (
	"fmt"
)

func Arrays() {
	// Duas maneiras diferentes de declarar um array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	//fmt.Println(arr[4]) //Erro devido ao array ter 3 elementos

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}

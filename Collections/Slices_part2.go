//Slices Part 2

package Collections

import (
	"fmt"
)

func main_slices_part2() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	//Append -> Acrescenta elementos na matriz
	slice = append(slice, 4, 42, 27)

	fmt.Println(slice)

	s2 := slice[1:]  //2-3-4-42-27
	s3 := slice[:2]  //1-2
	s4 := slice[1:2] //2

	fmt.Println(s2, s3, s4)

}

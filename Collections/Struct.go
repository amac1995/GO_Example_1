//Struct

package Collections

import (
	"fmt"
)

func main_struct() {
	type user2 struct {
		ID        int
		FirstName string
		LastName  string
	}
	var us user2

	us.ID = 1
	us.FirstName = "Arthur"
	us.LastName = "Dent"

	fmt.Println(us)

	u2 := user2{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}

	fmt.Println(u2)

}

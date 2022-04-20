package programscrud

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func MainConditions() {

	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}

	if u1 == u2 {
		println("Users Iguais")
	} else if u1.FirstName == u2.FirstName {
		println("Users Similares")
	} else {
		println("Users diferentes")
	}

}

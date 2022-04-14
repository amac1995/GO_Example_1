package collections

type user struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*user
	nextID = 1
)

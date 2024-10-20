package data

type User struct {
	ID    int
	Name  string
	Books []Book
}

type Book struct {
	ID     int
	Title  string
	Author string
}

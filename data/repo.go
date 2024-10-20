package data

import "fmt"

type Repository interface {
	AddBook(userID int, book Book) error
	DeleteBook(userID, bookID int) error
	GetUser(userID int) (*User, error)
}

type InMemoryRepository struct {
	Users      map[int]*User
	nextBookID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Users:      make(map[int]*User),
		nextBookID: 1,
	}
}

func (r *InMemoryRepository) AddBook(userID int, book Book) error {
	user, ok := r.Users[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	book.ID = r.nextBookID
	r.nextBookID++
	user.Books = append(user.Books, book)
	return nil
}

func (r *InMemoryRepository) DeleteBook(userID, bookID int) error {
	user, ok := r.Users[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	for i, book := range user.Books {
		if book.ID == bookID {
			user.Books = append(user.Books[:i], user.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func (r *InMemoryRepository) GetUser(userID int) (*User, error) {
	user, ok := r.Users[userID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

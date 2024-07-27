package repositories

import "github.com/cnc-csku/go-clean-arch-example/entities"

type BookRepository interface {
	GetAll() []entities.Book
	Create(book entities.Book) (*entities.Book, error)
}
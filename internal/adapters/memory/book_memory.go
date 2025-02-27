package memory

import (
	"github.com/cnc-csku/go-clean-arch-example/entities"
	"github.com/cnc-csku/go-clean-arch-example/exceptions"
	"github.com/cnc-csku/go-clean-arch-example/repositories"
)

type BookMemory struct {
	books []entities.Book
}

func NewBookMemory() repositories.BookRepository {
	return &BookMemory{
		books: []entities.Book{},
	}
}

func (b *BookMemory) GetAll() []entities.Book {
	return b.books
}

func (b *BookMemory) Create(book entities.Book) (*entities.Book, error) {
	for _, b := range b.books {
		if b.ID == book.ID {
			return nil, exceptions.ErrBookIdDup
		}
	}

	b.books = append(b.books, book)

	return &book, nil
}
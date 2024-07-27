package usecases

import (
	"github.com/cnc-csku/go-clean-arch-example/entities"
	"github.com/cnc-csku/go-clean-arch-example/repositories"
)

type BookUseCases interface {
	GetAllBooks() []entities.Book
	CreateBook(book entities.Book) (*entities.Book, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookUseCases {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (b *bookService) GetAllBooks() []entities.Book {
	return b.bookRepo.GetAll()
}

func (b *bookService) CreateBook(book entities.Book) (*entities.Book, error) {
	return b.bookRepo.Create(book)
}
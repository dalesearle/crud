package operator

import "crud/book"

type CrudOperator interface {
	Create(book *book.Book) *book.Book
	ReadAll() []*book.Book
	Read(id int) *book.Book
	Update(book *book.Book) *book.Book
	Delete(id int) (*book.Book, error)
}

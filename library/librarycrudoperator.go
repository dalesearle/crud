package library

import (
	"crud/book"
	"errors"
	"sort"
	"strconv"
	"sync"
)

type LibraryCrudOperator struct {
	library map[int]*book.Book
	lock    sync.Mutex
}

func New() *LibraryCrudOperator {
	sc := LibraryCrudOperator{
		library: make(map[int]*book.Book),
	}
	b := book.New("George R.R. Martin", "Game Of Thrones")
	sc.library[b.Id] = b
	b = book.New("George R.R. Martin", "A Clash Of Kings")
	sc.library[b.Id] = b
	b = book.New("George R.R. Martin", "A Storm Of Swords")
	sc.library[b.Id] = b
	b = book.New("George R.R. Martin", "A Feast Of Crows")
	sc.library[b.Id] = b
	b = book.New("George R.R. Martin", "A Dance With Dragons")
	sc.library[b.Id] = b
	b = book.New("George Rat Bastard Martin", "No More Books For You, Sucker Fish")
	sc.library[b.Id] = b
	b = book.New("George Orwell", "1984")
	sc.library[b.Id] = b
	b = book.New("George Orwell", "Animal Farm")
	sc.library[b.Id] = b
	b = book.New("Anne Rand", "Atlas Shrugged")
	sc.library[b.Id] = b
	b = book.New("Harper Lee", "To Kill A Mocking Bird")
	sc.library[b.Id] = b
	return &sc
}

func (sc *LibraryCrudOperator) Create(b *book.Book) *book.Book {
	b = book.New(b.Author, b.Title)
	sc.lock.Lock()
	sc.library[b.Id] = b
	sc.lock.Unlock()
	return b
}

func (sc *LibraryCrudOperator) Read(id int) *book.Book {
	return sc.library[id]
}

func (sc *LibraryCrudOperator) ReadAll() []*book.Book {
	rval := make(ResultSet, 0, len(sc.library))
	for _, v := range sc.library {
		rval = append(rval, v)
	}
	sort.Sort(rval)
	return rval
}

func (sc *LibraryCrudOperator) Update(book *book.Book) *book.Book {
	sc.lock.Lock()
	arch := sc.library[book.Id]
	if arch == nil {
		sc.lock.Unlock()
		return sc.Create(book)
	}
	if book.Author != "" {
		arch.Author = book.Author
	}
	if book.Title != "" {
		arch.Title = book.Title
	}
	sc.lock.Unlock()
	return arch
}

func (sc *LibraryCrudOperator) Delete(id int) (*book.Book, error) {
	sc.lock.Lock()
	defer sc.lock.Unlock()
	book := sc.library[id]
	if book != nil {
		delete(sc.library, id)
		return book, nil
	}
	return nil, errors.New("Unable to locate book with ID " + strconv.Itoa(id))
}

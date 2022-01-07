package simplecrudder

import (
	"crud/structs"
	"sort"
)

type SimpleCrudder struct {
	library map[int]*structs.Book
}

func NewSimpleCrudder() *SimpleCrudder {
	sc := SimpleCrudder{
		library: make(map[int]*structs.Book),
	}
	book := structs.NewBook("George R.R. Martin", "Game Of Thrones")
	sc.library[book.Id] = book
	book = structs.NewBook("George R.R. Martin", "A Clash Of Kings")
	sc.library[book.Id] = book
	book = structs.NewBook("George R.R. Martin", "A Storm Of Swords")
	sc.library[book.Id] = book
	book = structs.NewBook("George R.R. Martin", "A Feast Of Crows")
	sc.library[book.Id] = book
	book = structs.NewBook("George R.R. Martin", "A Dance With Dragons")
	sc.library[book.Id] = book
	book = structs.NewBook("George Rat Bastard Martin", "No More Books For You, Sucker Fish")
	sc.library[book.Id] = book
	book = structs.NewBook("George Orwell", "1984")
	sc.library[book.Id] = book
	book = structs.NewBook("George Orwell", "Animal Farm")
	sc.library[book.Id] = book
	book = structs.NewBook("Anne Rand", "Atlas Shrugged")
	sc.library[book.Id] = book
	book = structs.NewBook("Harper Lee", "To Kill A Mocking Bird")
	sc.library[book.Id] = book
	return &sc
}

func (sc *SimpleCrudder) Create(book *structs.Book) {
	book = structs.NewBook(book.Author, book.Title)
	sc.library[book.Id] = book
}

func (sc *SimpleCrudder) Read() []*structs.Book {
	rval := make(structs.Library, 0, len(sc.library))
	for _, v := range sc.library {
		rval = append(rval, v)
	}
	sort.Sort(rval)
	return rval
}

func (sc *SimpleCrudder) Update(original, updated string) {

}

func (sc *SimpleCrudder) Delete(id int) {
	delete(sc.library, id)
}

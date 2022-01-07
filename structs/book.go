package structs

import "sync"

var id = 1
var idMutex = new(sync.Mutex)

type Book struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

func NewBook(author, title string) *Book {
	book := Book{
		Id:     genId(),
		Author: author,
		Title:  title,
	}
	return &book
}

func genId() int {
	idMutex.Lock()
	var rval = id
	id++
	idMutex.Unlock()
	return rval
}

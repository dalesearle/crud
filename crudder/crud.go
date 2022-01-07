package crudder

import "crud/structs"

type Crud interface {
	Create(book *structs.Book)
	Read() []*structs.Book
	Update(original, updated string)
	Delete(id int)
}

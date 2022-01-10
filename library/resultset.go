package library

import (
	"crud/book"
)

type ResultSet []*book.Book

func (l ResultSet) Len() int           { return len(l) }
func (l ResultSet) Less(i, j int) bool { return l[i].Id < l[j].Id }
func (l ResultSet) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

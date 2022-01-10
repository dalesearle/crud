package handler

import (
	"crud/book"
	"crud/library"
	"crud/operator"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type LibraryHandler struct {
	crudOperator operator.CrudOperator
}

func New() *LibraryHandler {
	h := LibraryHandler{
		crudOperator: library.New(),
	}
	return &h
}

func (sch *LibraryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		handleCreateRequest(w, req, sch.crudOperator)
	case http.MethodGet:
		handleReadRequest(w, req, sch.crudOperator)
	case http.MethodPut:
		handleUpdateRequest(w, req, sch.crudOperator)
	case http.MethodDelete:
		handleDeleteRequest(w, req, sch.crudOperator)
	default:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported method " + req.Method))
	}
}

func handleReadRequest(w http.ResponseWriter, req *http.Request, crudOperator operator.CrudOperator) {
	var rval []byte
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id, err := getBookIdFromURL(req)
	if err != nil {
		rval, _ = json.Marshal(crudOperator.ReadAll())
	} else {
		rval, _ = json.Marshal(crudOperator.Read(id))
	}
	w.Write(rval)
}

// Create
func handleCreateRequest(w http.ResponseWriter, req *http.Request, crudOperator operator.CrudOperator) {
	w.Header().Set("Content-Type", "text/plain")
	book, err := unMarshallBook(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse a book from the given document"))
	} else {
		book = crudOperator.Create(book)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Added resource, " + strconv.Itoa(book.Id) + " - " + book.Title))
	}
}

func handleDeleteRequest(w http.ResponseWriter, req *http.Request, crudOperator operator.CrudOperator) {
	w.Header().Set("Content-Type", "text/plain")
	id, err := getBookIdFromURL(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unable to delete a resource without an identifier"))
		return
	}
	book, err := crudOperator.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("deleted " + book.Title))
	}
}

func handleUpdateRequest(w http.ResponseWriter, req *http.Request, crudOperator operator.CrudOperator) {
	w.Header().Set("Content-Type", "text/plain")
	book, err := unMarshallBook(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse a book update from the given document. " + err.Error()))
	} else {
		book = crudOperator.Update(book)
		w.WriteHeader(http.StatusOK)
		json, _ := json.Marshal(book)
		w.Write(json)
	}
}

func unMarshallBook(req *http.Request) (*book.Book, error) {
	b := book.Book{}
	err := json.NewDecoder(req.Body).Decode(&b)
	return &b, err
}

func getBookIdFromURL(req *http.Request) (int, error) {
	parts := strings.Split(req.URL.Path, "/")
	id := parts[len(parts)-1]
	return strconv.Atoi(id)
}

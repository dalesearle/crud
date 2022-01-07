package handler

import (
	"crud/simplecrudder"
	"crud/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

//TODO: crudder interface
type SimpleCrudHandler struct {
	crudder *simplecrudder.SimpleCrudder
}

func NewSimpleCrudHandler() *SimpleCrudHandler {
	h := SimpleCrudHandler{
		crudder: simplecrudder.NewSimpleCrudder(),
	}
	return &h
}

func (sch *SimpleCrudHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		serviceGetRequest(w, req, sch.crudder)
	case http.MethodPost:
		servicePostRequest(w, req, sch.crudder)
	case http.MethodDelete:
		serviceDeleteRequest(w, req, sch.crudder)
	case http.MethodPatch:
		servicePatchRequest(w, req, sch.crudder)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func serviceGetRequest(w http.ResponseWriter, req *http.Request, crudder *simplecrudder.SimpleCrudder) {
	json, _ := json.Marshal(crudder.Read())
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// Create
func servicePostRequest(w http.ResponseWriter, req *http.Request, crudder *simplecrudder.SimpleCrudder) {
	b := unMarshallBook(req)
	crudder.Create(b)
}

func serviceDeleteRequest(w http.ResponseWriter, req *http.Request, crudder *simplecrudder.SimpleCrudder) {
	parts := strings.Split(req.URL.Path, "/")
	id := parts[len(parts)-1]
	// TODO: error handling
	idActual, _ := strconv.Atoi(id)
	crudder.Delete(idActual)
	io.WriteString(w, parts[len(parts)-1])
}

// Create
func servicePatchRequest(w http.ResponseWriter, req *http.Request, crudder *simplecrudder.SimpleCrudder) {
	io.WriteString(w, "Handle PATCH\n")
}

func unMarshallBook(req *http.Request) *structs.Book {
	b := structs.Book{}
	// TODO: handle the error
	json.NewDecoder(req.Body).Decode(&b)
	return &b
}

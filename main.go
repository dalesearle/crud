package main

import (
	"crud/handler"
	"log"
	"net/http"
)

func main() {
	//deleteHandler := func(w http.ResponseWriter, req *http.Request) {
	//	req.URL.
	//		io.WriteString(w, "Delete")
	//}
	h := handler.NewSimpleCrudHandler()

	http.Handle("/library", h)
	http.Handle("/library/", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
